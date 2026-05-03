package ingestion

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/keithyw/pitch-in/internal/domains/assets"
	"github.com/keithyw/pitch-in/internal/domains/content/items"
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	"github.com/keithyw/pitch-in/pkg/response"
)

type IngestionRequest struct {
	Name string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	Tags []string `json:"tags" validate:"required"`
}

type IngestionHandler struct {
	assetService assets.AssetService
	itemService items.ItemService
	tagService tags.TagService
	log *slog.Logger
}

func NewIngestionHandler(assetService assets.AssetService, itemService items.ItemService, tagService tags.TagService, log *slog.Logger) *IngestionHandler {
	return &IngestionHandler{
		assetService: assetService,
		itemService: itemService,
		tagService: tagService,
		log: log,
	}
}

func (h *IngestionHandler) Post(w http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value("UserID").(int64)
	if !ok {
		response.ErrorJSON(w, http.StatusUnauthorized, "Invalid user ID")
		return
	}

	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	item := IngestionRequest{
		Name: req.FormValue("name"),
		Description: req.FormValue("description"),
		Tags: req.Form["tags"],
	}

	file, header, err := req.FormFile("file")
	if err != nil {
		response.ErrorJSON(w, http.StatusBadRequest, "File is required")
		return
	}
	defer file.Close()

	asset := assets.Asset{
		AssetFields: assets.AssetFields{
			ObjectKey: &header.Filename,
		},
	}

	// use this to associate to the item content
	newAsset, err := h.assetService.CreateAsset(asset, file)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	newTags := []tags.Tag{}
	for _, tagName := range item.Tags {
		newTags = append(newTags, tags.Tag{
			TagFields: tags.TagFields{
				Tag: &tagName,
			},
		})
	}

	itemReq := items.Item{
		ItemFields: items.ItemFields{
			Name: &item.Name,
			Description: &item.Description,
		},
		UserID: &userID,
		Tags: newTags,
	}


	newItem, err := h.itemService.CreateItem(itemReq)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: turn entity types into constant
	err = h.assetService.AttachEntity(newItem.ID, newAsset.ID, "items")
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	tagMap, err := h.tagService.GetIDsByTags(itemReq.Tags)
	if err != nil {
		response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
		return	
	}
	
	if len(tagMap) > 0 {
		for _, id := range tagMap {
			err = h.tagService.AttachEntity(newItem.ID, id, "items")
			if err != nil {
				response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	var tagsToCreate []string
	for _, tagName := range itemReq.Tags {
		lowerTagName := strings.ToLower(*tagName.Tag)
		if _,exists := tagMap[lowerTagName]; !exists {
			tagsToCreate = append(tagsToCreate, lowerTagName)
		}
	}

	for _, newTag := range tagsToCreate {
		tag := tags.Tag{
			TagFields: tags.TagFields{
				Tag: &newTag,
			},
		}
		newTag, err := h.tagService.CreateTag(tag)
		if err != nil {
			response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = h.tagService.AttachEntity(newItem.ID, newTag.ID, "items")
		if err != nil {
			response.ErrorJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	response.JSON(w, http.StatusCreated, nil)
}