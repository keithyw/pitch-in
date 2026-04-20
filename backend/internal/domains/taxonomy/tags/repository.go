package tags

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type TagRepository interface {
	AttachEntity(entityID, tagID int64, entity string) error
	CountTags(filter repository.Filter) (int64, error)
	CreateTag(tag Tag) (*Tag, error)
	DeleteTag(id int64) error
	DetachEntity(entityID, tagID int64, entity string) error
	FindTagsBy(filter repository.Filter) ([]Tag, error)
	GetIDsByTags(tags []Tag) (map[string]int64, error)
	GetTag(id int64) (*Tag, error)
	GetTagsByEntity(entityID int64, entity string) ([]Tag, error)
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

func (t *tagRepositoryImpl) AttachEntity(entityID, tagID int64, entity string) error {
	data := map[string]interface{}{
		"entity_id": entityID,
		"tag_id": tagID,
		"entity": entity,
	}
	
	builder := sq.Insert("entity_tags").SetMap(data).PlaceholderFormat(sq.Question)
	_, err := t.store.GetClient().Exec(t.store.GetContext(), builder)
	return err
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

func (t *tagRepositoryImpl) DetachEntity(entityID, tagID int64, entity string) error {
	q := sq.Delete("entity_tags").
		Where(sq.Eq{"entity_id": entityID, "tag_id": tagID, "entity": entity}).
		PlaceholderFormat(sq.Question)
	_, err := t.store.GetClient().Exec(t.store.GetContext(), q)
	return err
}

func (t *tagRepositoryImpl) FindTagsBy(filter repository.Filter) ([]Tag, error) {
	var tags []Tag
	err := t.store.FindBy(&Tag{}, filter, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *tagRepositoryImpl) GetIDsByTags(tags []Tag) (map[string]int64, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	tagMap := make(map[string]int64)

	tagStr := make([]string, len(tags))
	for i, tag := range tags {
		tagStr[i] = *tag.Tag
	}

	builder := sq.Select("id", "tag").
		From("tags").
		Where(sq.Eq{"tag": tagStr}).
		PlaceholderFormat(sq.Question)

	rows, err := t.store.GetClient().QueryMany(builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var tag string
		if err := rows.Scan(&id, &tag); err != nil {
			return nil, err
		}
		tagMap[tag] = id
	}

	return tagMap, nil
}

func (t *tagRepositoryImpl) GetTag(id int64) (*Tag, error) {
	var tag Tag
	err := t.store.Get(&Tag{ BaseModel: model.BaseModel{ ID: id }}, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *tagRepositoryImpl) GetTagsByEntity(entityID int64, entity string) ([]Tag, error) {
	var tags []Tag
	builder := sq.Select("t.id as id, t.tag as tag, t.slug as slug, t.created_at as created_at, t.updated_at as updated_at, t.deleted_at as deleted_at").
		From("tags t").
		Join("entity_tags et ON t.id = et.tag_id").
		Where(sq.Eq{"et.entity_id": entityID, "et.entity": entity}).
		PlaceholderFormat(sq.Question)
	rows, err := t.store.GetClient().QueryMany(builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Tag, &tag.Slug, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (t *tagRepositoryImpl) UpdateTag(tag Tag) (*Tag, error) {
	var updatedTag Tag
	err := t.store.Update(&tag, tag.ToMap(), &updatedTag)
	if err != nil {
		return nil, err
	}
	return &updatedTag, nil
}