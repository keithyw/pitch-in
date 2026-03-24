package tags

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type TagRepository interface {
	CountTags(filter repository.Filter) (int64, error)
	CreateTag(tag Tag) (*Tag, error)
	DeleteTag(id int64) error
	FindTagsBy(filter repository.Filter) ([]Tag, error)
	GetTag(id int64) (*Tag, error)
	UpdateTag(tag Tag) (*Tag, error)
}

type tagRepositoryImpl struct {
	store database.DBStore
}

func NewTagRepository(store database.DBStore) TagRepository {
	return &tagRepositoryImpl{
		store: store,
	}
}

func (t *tagRepositoryImpl) CountTags(filter repository.Filter) (int64, error) {
	return t.store.Count(&Tag{}, filter)
}

func (t *tagRepositoryImpl) CreateTag(tag Tag) (*Tag, error) {
	var newTag Tag
	err := t.store.Create(&tag, tag.ToMap(), &newTag)
	if err != nil {
		return nil, err
	}
	return &newTag, nil
}

func (t *tagRepositoryImpl) DeleteTag(id int64) error {
	return t.store.Delete(&Tag{BaseModel: model.BaseModel{ ID: id }})
}

func (t *tagRepositoryImpl) FindTagsBy(filter repository.Filter) ([]Tag, error) {
	var tags []Tag
	err := t.store.FindBy(&Tag{}, filter, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *tagRepositoryImpl) GetTag(id int64) (*Tag, error) {
	var tag Tag
	err := t.store.Get(&Tag{ BaseModel: model.BaseModel{ ID: id }}, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *tagRepositoryImpl) UpdateTag(tag Tag) (*Tag, error) {
	var updatedTag Tag
	err := t.store.Update(&tag, tag.ToMap(), &updatedTag)
	if err != nil {
		return nil, err
	}
	return &updatedTag, nil
}