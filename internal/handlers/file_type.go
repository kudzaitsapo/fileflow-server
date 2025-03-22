package handlers

import (
	"fmt"
	"net/http"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

type FileTypeCreateRequest struct {
	Name        string `json:"name"`
	MimeType    string `json:"mimetype"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type FileTypeResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	MimeType    string `json:"mimetype"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CreatedAt   string `json:"created_at"`
}

func HandleCreateMimeType(w http.ResponseWriter, r *http.Request) {
	var payload FileTypeCreateRequest
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, fmt.Sprintf("invalid request payload: %v", err))
		return
	}

	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	fileType := &store.FileType{
		Name:        payload.Name,
		Description: payload.Description,
		MimeType:    payload.MimeType,
		Icon:        payload.Icon,
	}

	err := appStorage.FileTypes.Create(r.Context(), fileType)

	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "failed to create project")
		return
	}

	response := &FileTypeResponse{
		ID:          fileType.ID,
		Name:        fileType.Name,
		Description: fileType.Description,
		MimeType:    fileType.MimeType,
		Icon:        fileType.Icon,
		CreatedAt:   fileType.CreatedAt,
	}

	WriteJson(w, http.StatusCreated, response)

}

func HandleGetAllFileTypes(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	limit, offset := GetPaginationParams(r)

	fileTypes, err := appStorage.FileTypes.GetAll(r.Context(), limit, offset)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("failed to fetch file types: %v", err))
		return
	}

	response := make([]FileTypeResponse, 0, len(fileTypes))
	for _, fileType := range fileTypes {
		response = append(response, FileTypeResponse{
			ID:          fileType.ID,
			Name:        fileType.Name,
			Description: fileType.Description,
			MimeType:    fileType.MimeType,
			Icon:        fileType.Icon,
			CreatedAt:   fileType.CreatedAt,
		})
	}

	WriteJson(w, http.StatusOK, response)
}
