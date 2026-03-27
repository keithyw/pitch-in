package assets

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/keithyw/pitch-in/pkg/storage"
)

type AssetService interface {
	AttachEntity(entityID, assetID int64, entity string) error
	CountAssets(filter repository.Filter) (int64, error)
	CreateAsset(asset Asset, reader io.Reader) (*Asset, error)
	DeleteAsset(id int64) error
	FindAssetBy(filter repository.Filter) ([]Asset, error)
	GetAsset(id int64) (*Asset, error)
	GetAssetsByEntity(entityID int64, entity string) ([]Asset, error)
	UpdateAsset(asset Asset, reader io.Reader) (*Asset, error)
}

type assetServiceImpl struct {
	repository AssetRepository
	client storage.StorageClient 
	log *slog.Logger
}

func NewAssetService(repo AssetRepository, client storage.StorageClient, log *slog.Logger) AssetService {
	return &assetServiceImpl{
		repository: repo,
		client: client,
		log: log,
	}
}

func (s *assetServiceImpl) AttachEntity(entityID, assetID int64, entity string) error {
	err := s.repository.AttachEntity(entityID, assetID, entity)
	if err != nil {
		s.log.Error("Failed attaching entity", "entityID", entityID, "assetID", assetID, "entity", entity, "error", err)
		return fmt.Errorf("attach entity failure: %w", err)
	}
	return nil
}

func (s *assetServiceImpl) CountAssets(filter repository.Filter) (int64, error) {
	cnt, err := s.repository.CountAssets(filter)
	if err != nil {
		s.log.Error("Failed getting asset count", "error", err)
		return 0, fmt.Errorf("asset count failure: %w", err)
	}
	return cnt, nil
}

func (s *assetServiceImpl) CreateAsset(asset Asset, reader io.Reader) (*Asset, error) {
	buffer := make([]byte, 512)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		s.log.Error("Failed reading file", "error", err)
		return nil, fmt.Errorf("Failed to read file: %w", err)
	}

	detectedType := http.DetectContentType(buffer[:n])
	asset.MimeType = &detectedType

	fullReader := io.MultiReader(bytes.NewReader(buffer[:n]), reader)

	res, err := s.client.Put(fullReader, *asset.ObjectKey, *asset.MimeType)
	if err != nil {
		s.log.Error("Failed uploading asset", "object_key", *asset.ObjectKey, "error", err)
		return nil, fmt.Errorf("asset upload failure: %w", err)
	}
	asset.SizeBytes = &res.Size
	newAsset, err := s.repository.CreateAsset(asset)
	if err != nil {
		_ = s.client.Remove(*asset.ObjectKey)
		s.log.Error("Failed creating asset", "object_key", asset.ObjectKey, "error", err)
		return nil, fmt.Errorf("asset create failure: %w", err)
	}
	newAsset.URL = &res.URL
	return newAsset, nil
}

func (s *assetServiceImpl) DeleteAsset(id int64) error {
	asset, err := s.repository.GetAsset(id)
	if err != nil {
		s.log.Error("Asset does not exist to delete", "id", id, "error", err)
		return fmt.Errorf("asset does not exist: %w", err)
	}

	err = s.repository.DeleteAsset(id)
	if err != nil {		
		s.log.Error("Failed deleting asset", "id", id, "error", err)
		return fmt.Errorf("asset delete failure: %w", err)
	}

	err = s.client.Remove(*asset.ObjectKey)
	if err != nil {
		s.log.Error("Failed removing asset from storage", "key", asset.ObjectKey, "error", err)
		return fmt.Errorf("asset storage removal failure: %w", err)
	}
	return nil
}

func (s *assetServiceImpl) DetachEntity(entityID, assetID int64, entity string) error {
	err := s.repository.DetachEntity(entityID, assetID, entity)
	if err != nil {
		s.log.Error("Failed detaching entity", "error", err)
		return fmt.Errorf("detach entity failure: %w", err)
	}
	return nil
}

func (s *assetServiceImpl) FindAssetBy(filter repository.Filter) ([]Asset, error) {
	assets, err := s.repository.FindAssetBy(filter)
	if err != nil {
		s.log.Error("Failed finding asset", "error", err)
		return nil, fmt.Errorf("find by asset error: %w", err)
	}
	for i := range assets {
		resp, err := s.client.Get(*assets[i].ObjectKey)
		if err != nil {
			s.log.Error("Failed getting presign url", "object_id", assets[i].ObjectKey)
			return nil, fmt.Errorf("failed getting presign url: %w", err)
		}
		assets[i].URL = &resp.URL
	}
	return assets, nil
}

func (s *assetServiceImpl) GetAsset(id int64) (*Asset, error) {
	asset, err := s.repository.GetAsset(id)
	if err != nil {
		s.log.Error("Failed getting asset", "id", id, "error", err)
		return nil, fmt.Errorf("Get asset error: %w", err)
	}

	resp, err := s.client.Get(*asset.ObjectKey)
	if err != nil {
		s.log.Error("Could not find asset", "key", asset.ObjectKey, "error", err)
		return nil, err
	}

	asset.URL = &resp.URL
	return asset, nil
}

func (s *assetServiceImpl) GetAssetsByEntity(entityID int64, entity string) ([]Asset, error) {
	assets, err := s.repository.GetAssetsByEntity(entityID, entity)
	if err != nil {
		s.log.Error("Failed getting assets by entity", "entityID", entityID, "entity", entity, "error", err)
		return nil, fmt.Errorf("get assets by entity error: %w", err)
	}
	for i := range assets {
		resp, err := s.client.Get(*assets[i].ObjectKey)
		if err != nil {
			s.log.Error("Could not find asset", "key", assets[i].ObjectKey, "error", err)
			return nil, err
		}
		assets[i].URL = &resp.URL
	}
	return assets, nil
}

func (s *assetServiceImpl) UpdateAsset(asset Asset, reader io.Reader) (*Asset, error) {
	existingAsset, err := s.repository.GetAsset(asset.ID)
	if err != nil {
		s.log.Error("Failed getting asset", "id", asset.ID, "error", err)
		return nil, err
	}

	var newStorageRes *storage.UploadResponse

	if reader != nil {
		buffer := make([]byte, 512)
		n, _ := reader.Read(buffer)
		detectedType := http.DetectContentType(buffer[:n])
		asset.MimeType = &detectedType

		fullReader := io.MultiReader(bytes.NewReader(buffer[:n]), reader)
		resp, err := s.client.Put(fullReader, *asset.ObjectKey, *asset.MimeType)
		if err != nil {
			s.log.Error("Faile uploading asset", "id", asset.ID, "error", err)
			return nil, fmt.Errorf("storage upload failed: %w", err)
		}

		newStorageRes = resp
		asset.SizeBytes = &resp.Size
	}

	updatedAsset, err := s.repository.UpdateAsset(asset)
	if err != nil {
		if newStorageRes != nil {
			_ = s.client.Remove(*asset.ObjectKey)
		}
		s.log.Error("Failed updating asset", "id", asset.ID, "error", err)
		return nil, fmt.Errorf("Update asset error: %w", err)
	}

	if newStorageRes != nil && *existingAsset.ObjectKey != *asset.ObjectKey {
		_ = s.client.Remove(*existingAsset.ObjectKey)
	}

	if newStorageRes != nil {
		updatedAsset.URL = &newStorageRes.URL
	} else {
		resp, _ := s.client.Get(*updatedAsset.ObjectKey)
		updatedAsset.URL = &resp.URL
	}

	return updatedAsset, nil
}