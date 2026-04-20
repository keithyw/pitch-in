package tags

import (
	"fmt"
	"log/slog"

	"github.com/gosimple/slug"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type TagService interface {
	AttachEntity(entityID, tagID int64, entity string) error
	CountTags (filter repository.Filter) (int64, error)
	CreateTag (tag Tag) (*Tag, error)
	DeleteTag (id int64) error
	DetachEntity(entityID, tagID int64, entity string) error
	FindTagsBy (filter repository.Filter) ([]Tag, error)
	GetTag (id int64) (*Tag, error)
	GetIDsByTags(tags []Tag) (map[string]int64, error)
	GetTagsByEntity(entityID int64, entity string) ([]Tag, error)
	UpdateTag (tag Tag) (*Tag, error)
}

type tagServiceImpl struct {
	repository TagRepository
	log *slog.Logger
}

func NewTagService(repo TagRepository, log *slog.Logger) TagService {
	return &tagServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *tagServiceImpl) AttachEntity(entityID, tagID int64, entity string) error {
	err := s.repository.AttachEntity(entityID, tagID, entity)
	if err != nil {
		s.log.Error("Failed attaching entity", "entityID", entityID, "tagID", tagID, "entity", entity, "error", err)
		return fmt.Errorf("attach entity failure: %w", err)
	}
	return nil
}

func (s *tagServiceImpl) CountTags(filter repository.Filter) (int64, error) {
	cnt, err := s.repository.CountTags(filter)
	if err != nil {
		s.log.Error("Failed getting tag count", "error", err)
		return 0, fmt.Errorf("tag count failure: %w", err)
	}
	return cnt, nil
}

func (s *tagServiceImpl) CreateTag(tag Tag) (*Tag, error) {
	slug := slug.Make(*tag.Tag)
	tag.Slug = &slug
	newTag, err := s.repository.CreateTag(tag)
	if err != nil {
		s.log.Error("Failed creating new tag", "tag", tag.Tag, "error", err)
		return nil, fmt.Errorf("create tag error: %w", err)
	}
	return newTag, nil
}

func (s *tagServiceImpl) DeleteTag(id int64) error {
	err := s.repository.DeleteTag(id)
	if err != nil {
		s.log.Error("Failed deleting tag", "id", id, "error", err)
		return fmt.Errorf("delete tag error: %w", err)
	}
	return nil
}

func (s *tagServiceImpl) DetachEntity(entityID, tagID int64, entity string) error {
	err := s.repository.DetachEntity(entityID, tagID, entity)
	if err != nil {
		s.log.Error("Failed detaching entity", "entityID", entityID, "tagID", tagID, "entity", entity, "error", err)
		return fmt.Errorf("detach entity failure: %w", err)
	}
	return nil
}

func (s *tagServiceImpl) FindTagsBy(filter repository.Filter) ([]Tag, error) {
	tags, err := s.repository.FindTagsBy(filter)
	if err != nil {
		s.log.Error("Failed finding tags", "error", err)
		return nil, fmt.Errorf("Find tags by error: %w", err)
	}
	return tags, nil
}

func (s *tagServiceImpl) GetTag(id int64) (*Tag, error) {
	tag, err := s.repository.GetTag(id)
	if err != nil {
		s.log.Error("Failed getting tag", "id", id, "error", err)
		return nil, fmt.Errorf("get tag error: %w", err)
	}
	return tag, nil
}

func (s *tagServiceImpl) GetIDsByTags(tags []Tag) (map[string]int64, error) {
	tagMap, err := s.repository.GetIDsByTags(tags)
	if err != nil {
		s.log.Error("Failed getting tag IDs", "error", err)
		return nil, fmt.Errorf("get tag IDs error: %w", err)
	}
	return tagMap, nil
}

func (s *tagServiceImpl) GetTagsByEntity(entityID int64, entity string) ([]Tag, error) {
	tags, err := s.repository.GetTagsByEntity(entityID, entity)
	if err != nil {
		s.log.Error("Failed getting tags by entity", "entityID", entityID, "entity", entity, "error", err)
		return nil, fmt.Errorf("get tags by entity error: %w", err)
	}
	return tags, nil
}

func (s *tagServiceImpl) UpdateTag(tag Tag) (*Tag, error) {
	slug := slug.Make(*tag.Tag)
	tag.Slug = &slug
	updatedTag, err := s.repository.UpdateTag(tag)
	if err != nil {
		s.log.Error("Failed updating tag", "id", tag.ID, "error", err)
		return nil, fmt.Errorf("Update tag error: %w", err)
	}
	return updatedTag, nil
}