package assets

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type AssetRepository interface {
	CountAssets(filter repository.Filter) (int64, error)
	CreateAsset(asset Asset) (*Asset, error)
	DeleteAsset(id int64) error
	FindAssetBy(filter repository.Filter) ([]Asset, error)
	GetAsset(id int64) (*Asset, error)
	UpdateAsset(asset Asset) (*Asset, error)
}

type AssetRepositoryImpl struct {
	store database.DBStore
}

func NewAssetRepository(store database.DBStore) AssetRepository {
	return &AssetRepositoryImpl{
		store: store,
	}
}

func (r *AssetRepositoryImpl) CountAssets(filter repository.Filter) (int64, error) {
	return r.store.Count(&Asset{}, filter)
}

func (r *AssetRepositoryImpl) CreateAsset(asset Asset) (*Asset, error) {
	var newAsset Asset
	err := r.store.Create(&asset, asset.ToMap(), &newAsset)
	if err != nil {
		return nil, err
	}
	return &newAsset, nil
}

func (r *AssetRepositoryImpl) DeleteAsset(id int64) error {
	return r.store.Delete(&Asset{BaseModel: model.BaseModel{ ID: id }})
}

func (r *AssetRepositoryImpl) FindAssetBy(filter repository.Filter) ([]Asset, error) {
	var assets []Asset
	err := r.store.FindBy(&Asset{}, filter, &assets)
	if err != nil {
		return nil, err
	}
	return assets, err
}

func (r *AssetRepositoryImpl) GetAsset(id int64) (*Asset, error) {
	var asset Asset
	err := r.store.Get(&Asset{ BaseModel: model.BaseModel{ ID: id }}, &asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *AssetRepositoryImpl) UpdateAsset(asset Asset) (*Asset, error) {
	var updatedAsset Asset
	err := r.store.Update(&asset, asset.ToMap(), &updatedAsset)
	if err != nil {
		return nil, err
	}
	return &updatedAsset, nil
}