package items

import (
	"log/slog"

	"github.com/gosimple/slug"
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type ItemService interface {
	CountItems(filter repository.Filter) (int64, error)
	CreateItem(item Item) (*Item, error)
	DeleteItem(id int64) error
	FindItemsBy(filter repository.Filter) ([]Item, error)
	GetItem(id int64) (*Item, error)
	UpdateItem(item Item) (*Item, error)
}

type itemServiceImpl struct {
	repository ItemRepository
	tagRepository tags.TagRepository
	log *slog.Logger
}

func NewItemService(repo ItemRepository, tagRepo tags.TagRepository, log *slog.Logger) ItemService {
	return &itemServiceImpl{
		repository: repo,
		tagRepository: tagRepo,
		log: log,
	}
}

func (s *itemServiceImpl) CountItems(filter repository.Filter) (int64, error) {
	cnt, err := s.repository.CountItems(filter)
	if err != nil {
		s.log.Error("Failed getting item count", "error", err)
		return 0, err
	}
	return cnt, nil
}

func (s *itemServiceImpl) CreateItem(item Item) (*Item, error) {
	slug := slug.Make(*item.Name)
	item.Slug = &slug
	newTag, err := s.repository.CreateItem(item)
	if err != nil {
		s.log.Error("Failed creating new item", "item", item.Name, "error", err)
		return nil, err
	}
	return newTag, nil
}

func (s *itemServiceImpl) DeleteItem(id int64) error {
	err := s.repository.DeleteItem(id)
	if err != nil {
		s.log.Error("Failed deleting item", "id", id, "error", err)
		return err
	}
	return nil
}

func (s *itemServiceImpl) FindItemsBy(filter repository.Filter) ([]Item, error) {
	items, err := s.repository.FindItemsBy(filter)
	if err != nil {
		s.log.Error("Failed finding items", "error", err)
		return nil, err
	}
	return items, nil
}

func (s *itemServiceImpl) GetItem(id int64) (*Item, error) {
	item, err := s.repository.GetItem(id)
	if err != nil {
		s.log.Error("Failed getting item", "id", id, "error", err)
		return nil, err
	}

	tags, err := s.tagRepository.GetTagsByEntity(id, "items")
	if err != nil {
		s.log.Error("Failed getting tags by item", "id", id, "error", err)
		return nil, err
	}

	item.Tags = tags
	return item, nil
}

func (s *itemServiceImpl) UpdateItem(item Item) (*Item, error) {
	slug := slug.Make(*item.Name)
	item.Slug = &slug
	updatedItem, err := s.repository.UpdateItem(item)
	if err != nil {
		s.log.Error("Failed updating item", "id", item.ID, "error", err)
		return nil, err
	}
	return updatedItem, nil
}