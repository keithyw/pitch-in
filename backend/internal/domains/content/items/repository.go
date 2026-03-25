package items

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type ItemRepository interface {
	CountItems(filter repository.Filter) (int64, error)
	CreateItem(item Item) (*Item, error)
	DeleteItem(id int64) error
	FindItemsBy(filter repository.Filter) ([]Item, error)
	GetItem(id int64) (*Item, error)
	UpdateItem(item Item) (*Item, error)
}

type ItemRepositoryImpl struct {
	store database.DBStore
}

func NewItemRepository(store database.DBStore) ItemRepository {
	return &ItemRepositoryImpl{
		store: store,
	}
}

func (i *ItemRepositoryImpl) CountItems(filter repository.Filter) (int64, error) {
	return i.store.Count(&Item{}, filter)
}

func (i *ItemRepositoryImpl) CreateItem(item Item) (*Item, error) {
	var newItem Item
	err := i.store.Create(&item, item.ToMap(), &newItem)
	if err != nil {
		return nil, err
	}
	return &newItem, nil
}

func (i *ItemRepositoryImpl) DeleteItem(id int64) error {
	return i.store.Delete(&Item{ BaseModel: model.BaseModel{ ID: id }})
}

func (i *ItemRepositoryImpl) FindItemsBy(filter repository.Filter) ([]Item, error) {
	var items []Item
	err := i.store.FindBy(&Item{}, filter, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *ItemRepositoryImpl) GetItem(id int64) (*Item, error) {
	var item Item
	err := i.store.Get(&Item{ BaseModel: model.BaseModel{ ID: id }}, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *ItemRepositoryImpl) UpdateItem(item Item) (*Item, error) {
	var updatedItem Item
	err := i.store.Update(&item, item.ToMap(), &updatedItem)
	if err != nil {
		return nil, err
	}
	return &updatedItem, nil
}

