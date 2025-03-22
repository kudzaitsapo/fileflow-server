package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

type ProjectCreateRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	MaxUploadSize    int64    `json:"max_upload_size"`
	AllowedFileTypes []string `json:"allowed_file_types"`
}

type ProjectResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	CreatedAt        string   `json:"created_at"`
	CreatedById      int64    `json:"created_by_id"`
	ProjectKey       string   `json:"project_key"`
	MaxUploadSize    int64    `json:"max_upload_size"`
	AllowedFileTypes []string `json:"allowed_file_types"`
}

type ApiKeyRegenerationRequest struct {
	ID int64 `json:"id"`
}

func GenerateRandomKey() (string, error) {
	bytes := make([]byte, 12)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random key: %v", err)
	}
	return "proj_" + hex.EncodeToString(bytes), nil
}

func HandleProjectCreation(w http.ResponseWriter, r *http.Request) {
	var payload ProjectCreateRequest
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, fmt.Sprintf("invalid request payload: %v", err))
		return
	}

	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store
	currentUser, userErr := GetCurrentUser(r)

	if userErr != nil {
		log.Printf("Failed to get current user: %v", userErr)
		WriteJsonError(w, http.StatusUnauthorized, userErr.Error())
		return
	}

	projectKey, keyErr := GenerateRandomKey()
	if keyErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate project key: %v", keyErr))
		return
	}

	project := &store.Project{
		Name:          payload.Name,
		Description:   payload.Description,
		CreatedAt:     time.Now().Format(time.RFC3339),
		CreatedById:   currentUser.ID,
		ProjectKey:    projectKey,
		MaxUploadSize: payload.MaxUploadSize,
	}

	err := appStorage.Projects.Create(r.Context(), project)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "failed to create project")
		return
	}

	if len(payload.AllowedFileTypes) > 0 {
		for _, fileType := range payload.AllowedFileTypes {
			dbSavedType, typeGetErr := appStorage.FileTypes.GetByMimeType(r.Context(), fileType)
			if typeGetErr != nil {
				WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("failed to add file types to project: %v", typeGetErr))
				return
			}
			projectAllowedType := &store.ProjectAllowedFileType{
				ProjectID:  project.ID,
				FileTypeID: dbSavedType.ID,
				CreatedAt:  time.Now().Format(time.RFC3339),
			}
			err := appStorage.ProjectAllowedFileTypes.Create(r.Context(), projectAllowedType)
			if err != nil {
				WriteJsonError(w, http.StatusInternalServerError, "failed to add file type to project")
				return
			}
		}
	}

	response := &ProjectResponse{
		ID:               project.ID,
		Name:             project.Name,
		Description:      project.Description,
		CreatedAt:        project.CreatedAt,
		CreatedById:      project.CreatedById,
		ProjectKey:       project.ProjectKey,
		MaxUploadSize:    project.MaxUploadSize,
		AllowedFileTypes: payload.AllowedFileTypes,
	}

	WriteJson(w, http.StatusCreated, response)
}

func HandleProjectList(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	// get query params for limit and offset
	limit, offset := GetPaginationParams(r)

	projects, err := appStorage.Projects.GetAll(r.Context(), limit, offset)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get projects: %v", err))
		return
	}

	response := make([]*ProjectResponse, 0)
	for _, project := range projects {
		response = append(response, &ProjectResponse{
			ID:            project.ID,
			Name:          project.Name,
			Description:   project.Description,
			CreatedAt:     project.CreatedAt,
			CreatedById:   project.CreatedById,
			ProjectKey:    project.ProjectKey,
			MaxUploadSize: project.MaxUploadSize,
		})
	}

	WriteJson(w, http.StatusOK, response)
}

func HandleProjectUpdate(w http.ResponseWriter, r *http.Request) {
	var payload ProjectResponse
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, fmt.Sprintf("invalid request payload: %v", err))
		return
	}

	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	project, projectErr := appStorage.Projects.GetById(r.Context(), payload.ID)
	if projectErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project: %v", projectErr))
		return
	}

	project.Name = payload.Name
	project.Description = payload.Description
	project.MaxUploadSize = payload.MaxUploadSize

	err := appStorage.Projects.Update(r.Context(), project)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update project: %v", err))
		return
	}

	if len(payload.AllowedFileTypes) > 0 {
		for _, fileType := range payload.AllowedFileTypes {
			dbSavedType, typeGetErr := appStorage.FileTypes.GetByMimeType(r.Context(), fileType)
			if typeGetErr != nil {
				WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("failed to add file types to project: %v", typeGetErr))
				return
			}
			projectAllowedType := &store.ProjectAllowedFileType{
				ProjectID:  project.ID,
				FileTypeID: dbSavedType.ID,
				CreatedAt:  time.Now().Format(time.RFC3339),
			}
			err := appStorage.ProjectAllowedFileTypes.Create(r.Context(), projectAllowedType)
			if err != nil {
				WriteJsonError(w, http.StatusInternalServerError, "failed to add file type to project")
				return
			}
		}
	}

	response := &ProjectResponse{
		ID:            project.ID,
		Name:          project.Name,
		Description:   project.Description,
		CreatedAt:     project.CreatedAt,
		CreatedById:   project.CreatedById,
		ProjectKey:    project.ProjectKey,
		MaxUploadSize: project.MaxUploadSize,
	}

	WriteJson(w, http.StatusOK, response)
}

func HandleApiKeyRegeneration(w http.ResponseWriter, r *http.Request) {
	var payload ApiKeyRegenerationRequest
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, fmt.Sprintf("invalid request payload: %v", err))
		return
	}

	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	project, projectErr := appStorage.Projects.GetById(r.Context(), payload.ID)
	if projectErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project: %v", projectErr))
		return
	}

	projectKey, projKeyErr := GenerateRandomKey()
	if projKeyErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate project key: %v", projKeyErr))
		return
	}

	project.ProjectKey = projectKey

	err := appStorage.Projects.Update(r.Context(), project)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update project: %v", err))
		return
	}

	response := &ProjectResponse{
		ID:            project.ID,
		Name:          project.Name,
		Description:   project.Description,
		CreatedAt:     project.CreatedAt,
		CreatedById:   project.CreatedById,
		ProjectKey:    project.ProjectKey,
		MaxUploadSize: project.MaxUploadSize,
	}

	WriteJson(w, http.StatusOK, response)
}

func HandleProjectDeletion(w http.ResponseWriter, r *http.Request) {
	var payload ApiKeyRegenerationRequest
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, fmt.Sprintf("invalid request payload: %v", err))
		return
	}

	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	project, projectErr := appStorage.Projects.GetById(r.Context(), payload.ID)
	if projectErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project: %v", projectErr))
		return
	}

	err := appStorage.Projects.Delete(r.Context(), project.ID)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete project: %v", err))
		return
	}

	WriteJson(w, http.StatusNoContent, nil)
}
