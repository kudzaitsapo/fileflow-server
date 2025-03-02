package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

type ProjectCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedById int64  `json:"created_by_id"`
	ProjectKey  string `json:"project_key"`
}

func GenerateRandomKey() string {
	return uuid.New().String()
}


func HandleProjectCreation(w http.ResponseWriter, r *http.Request) {
	var payload ProjectCreateRequest
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "invalid request payload")
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

	project := &store.Project{
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   time.Now().Format(time.RFC3339),
		CreatedById: currentUser.ID,
		ProjectKey:  GenerateRandomKey(),
	}

	err := appStorage.Projects.Create(r.Context(), project)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "failed to create project")
		return
	}

	response := &ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		CreatedById: project.CreatedById,
		ProjectKey:  project.ProjectKey,
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
		WriteJsonError(w, http.StatusInternalServerError, "failed to fetch projects")
		return
	}

	response := make([]*ProjectResponse, 0)
	for _, project := range projects {
		response = append(response, &ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			CreatedById: project.CreatedById,
			ProjectKey:  project.ProjectKey,
		})
	}

	WriteJson(w, http.StatusOK, response)
}