package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

type ProjectUser struct {
	ID        int64 `json:"id"`
	ProjectID int64 `json:"project_id"`
	UserInfo  any   `json:"user_info"`
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

	// Assign current user to the project
	projectUser := &store.UserAssignedProject{
		UserID:    currentUser.ID,
		ProjectID: project.ID,
	}

	// create the user assigned project with transaction
	assignErr := appStorage.UserAssignedProjects.CreateWithoutTx(r.Context(), projectUser)

	if assignErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to assign created project to logged in user: %v", assignErr))
		return
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

	SendJsonWithoutMeta(w, http.StatusCreated, response)
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

	projectsCount, countErr := appStorage.Projects.Count(r.Context())
	if countErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project count: %v", countErr))
		return
	}
	meta := &JsonMeta{
		TotalRecords: projectsCount,
		Limit:        limit,
		Offset:       offset,
	}

	SendJson(w, http.StatusOK, response, *meta)
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

	SendJsonWithoutMeta(w, http.StatusOK, response)
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

	SendJsonWithoutMeta(w, http.StatusOK, response)
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

	SendJsonWithoutMeta(w, http.StatusNoContent, nil)
}

func HandleGetProjectInfo(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	projectId := r.PathValue("id")
	intProjectId, convErr := strconv.ParseInt(projectId, 10, 64)
	if convErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	project, projectErr := appStorage.Projects.GetById(r.Context(), intProjectId)
	if projectErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project: %v", projectErr))
		return
	}

	fileTypes := make([]string, 0)

	allowedFileTypes, fileTypeErr := appStorage.ProjectAllowedFileTypes.GetByProjectId(r.Context(), project.ID)
	if fileTypeErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get allowed file types: %v", fileTypeErr))
		return
	}
	for _, fileType := range allowedFileTypes {
		dbSavedType, typeGetErr := appStorage.FileTypes.GetById(r.Context(), fileType.FileTypeID)
		if typeGetErr != nil {
			WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get file type: %v", typeGetErr))
			return
		}
		fileTypes = append(fileTypes, dbSavedType.MimeType)
	}

	response := &ProjectResponse{
		ID:               project.ID,
		Name:             project.Name,
		Description:      project.Description,
		CreatedAt:        project.CreatedAt,
		CreatedById:      project.CreatedById,
		ProjectKey:       project.ProjectKey,
		MaxUploadSize:    project.MaxUploadSize,
		AllowedFileTypes: fileTypes,
	}

	SendJsonWithoutMeta(w, http.StatusOK, response)
}

func HandleGetProjectUsers(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	appStorage := currentApp.Store

	projectId := r.PathValue("id")
	intProjectId, convErr := strconv.ParseInt(projectId, 10, 64)
	if convErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	limit, offset := GetPaginationParams(r)

	projectUsers, projectUserErr := appStorage.UserAssignedProjects.GetByProjectId(r.Context(), intProjectId, limit, offset)
	if projectUserErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project users: %v", projectUserErr))
		return
	}

	response := make([]*ProjectUser, 0)
	for _, projectUser := range projectUsers {
		projUserInfo := &ProjectUser{
			ID:        projectUser.ID,
			ProjectID: projectUser.ProjectID,
			UserInfo: UserResponse{
				ID:        projectUser.User.ID,
				FirstName: projectUser.User.FirstName,
				LastName:  projectUser.User.LastName,
				Email:     projectUser.User.Email,
				CreatedAt: projectUser.User.CreatedAt,
				IsActive:  projectUser.User.IsActive,
				RoleID:    projectUser.User.RoleID,
			},
		}

		response = append(response, projUserInfo)
	}

	usersCount, countErr := appStorage.UserAssignedProjects.CountUsersByProjectId(r.Context(), intProjectId)
	if countErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get project users count: %v", countErr))
		return
	}

	meta := &JsonMeta{
		TotalRecords: usersCount,
		Limit:        limit,
		Offset:       offset,
	}

	SendJson(w, http.StatusOK, response, meta)
}
