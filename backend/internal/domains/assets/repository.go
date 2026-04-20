package assets

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type AssetRepository interface {
	AttachEntity(entityID, assetID int64, entity string) error	
	CountAssets(filter repository.Filter) (int64, error)
	CreateAsset(asset Asset) (*Asset, error)
	DeleteAsset(id int64) error
	DetachEntity(entityID, assetID int64, entity string) error
	FindAssetBy(filter repository.Filter) ([]Asset, error)
	GetAsset(id int64) (*Asset, error)
	GetAssetsByEntity(entityID int64, entity string) ([]Asset, error)
	UpdateAsset(asset Asset) (*Asset, error)
}

type assetRepositoryImpl struct {
	store database.DBStore
}

func NewAssetRepository(store database.DBStore) AssetRepository {
	return &assetRepositoryImpl{
		store: store,
	}
}

func (r *assetRepositoryImpl) AttachEntity(entityID, assetID int64, entity string) error {
	data := map[string]interface{}{
		"entity_id": entityID,
		"asset_id": assetID,
		"entity": entity,
	}

	builder := sq.Insert("asset_entities").SetMap(data).PlaceholderFormat(sq.Question)
	_, err := r.store.GetClient().Exec(r.store.GetContext(), builder)
	return err
}


func (r *assetRepositoryImpl) CountAssets(filter repository.Filter) (int64, error) {
	return r.store.Count(&Asset{}, filter)
}

func (r *assetRepositoryImpl) CreateAsset(asset Asset) (*Asset, error) {
	var newAsset Asset
	err := r.store.Create(&asset, asset.ToMap(), &newAsset)
	if err != nil {
		return nil, err
	}
	return &newAsset, nil
}

func (r *assetRepositoryImpl) DeleteAsset(id int64) error {
	return r.store.Delete(&Asset{BaseModel: model.BaseModel{ ID: id }})
}

func (r *assetRepositoryImpl) DetachEntity(entityID, assetID int64, entity string) error {
	q := sq.Delete("asset_entities").
		Where(sq.Eq{"entity_id": entityID, "asset_id": assetID, "entity": entity}).
		PlaceholderFormat(sq.Question)
	_, err := r.store.GetClient().Exec(r.store.GetContext(), q)
	return err
}

func (r *assetRepositoryImpl) FindAssetBy(filter repository.Filter) ([]Asset, error) {
	var assets []Asset
	err := r.store.FindBy(&Asset{}, filter, &assets)
	if err != nil {
		return nil, err
	}
	return assets, err
}

func (r *assetRepositoryImpl) GetAsset(id int64) (*Asset, error) {
	var asset Asset
	err := r.store.Get(&Asset{ BaseModel: model.BaseModel{ ID: id }}, &asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *assetRepositoryImpl) GetAssetsByEntity(entityID int64, entity string) ([]Asset, error) {
	var assets []Asset
	builder := sq.Select("a.id as id, a.object_key as object_key, a.mime_type as mime_type, a.size_bytes as size_bytes, a.width as width, a.height as height").
		From("assets a").
		Join("asset_entities ea ON a.id = ea.asset_id").
		Where(sq.Eq{"ea.entity_id": entityID, "ea.entity": entity}).
		PlaceholderFormat(sq.Question)
	rows, err := r.store.GetClient().QueryMany(builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.ID, &asset.ObjectKey, &asset.MimeType, &asset.SizeBytes, &asset.Width, &asset.Height); err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (r *assetRepositoryImpl) UpdateAsset(asset Asset) (*Asset, error) {
	var updatedAsset Asset
	err := r.store.Update(&asset, asset.ToMap(), &updatedAsset)
	if err != nil {
		return nil, err
	}
	return &updatedAsset, nil
}