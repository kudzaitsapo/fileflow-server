package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kudzaitsapo/fileflow-server/cmd/app"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
	RoleID    int64  `json:"role_id"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func HandleAuthentication(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := ReadJson(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	currentApp := app.GetCurrentApplication()
	store := currentApp.Store

	user, err := store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		WriteJsonError(w, http.StatusUnauthorized, fmt.Sprintf("error getting user: %v", err))
		return
	}

	if !user.IsActive {
		WriteJsonError(w, http.StatusUnauthorized, "user account is not active")
		return
	}

	passErr := user.Password.Compare(payload.Password)

	if passErr != nil {
		WriteJsonError(w, http.StatusUnauthorized, fmt.Sprintf("error comparing passwords: %v", passErr))
		return
	}

	claims := &jwt.MapClaims{
		"sub": strconv.FormatInt(user.ID, 10),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": "fileflow-server",
		"aud": user.RoleID,
	}

	token, err := currentApp.Authenticator.GenerateToken(jwt.MapClaims(*claims))
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error generating token")
		return
	}

	SendJsonWithoutMeta(w, http.StatusOK,
		LoginResponse{
			Token: token,
			User: UserResponse{
				ID:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				CreatedAt: user.CreatedAt,
				IsActive:  user.IsActive,
				RoleID:    user.RoleID,
			},
		},
	)
}

func HandleGetAssignedProjects(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	store := currentApp.Store

	user, err := GetCurrentUser(r)
	if err != nil {
		WriteJsonError(w, http.StatusUnauthorized, fmt.Sprintf("error getting current user: %v", err))
		return
	}

	projects, projErr := store.UserAssignedProjects.GetByUserId(r.Context(), user.ID)
	if projErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("error getting projects: %v", projErr))
		return
	}

	projectsData := make([]ProjectResponse, 0)

	for _, assignment := range projects {

		projectsData = append(projectsData, ProjectResponse{
			ID:            assignment.Project.ID,
			Name:          assignment.Project.Name,
			Description:   assignment.Project.Description,
			CreatedAt:     assignment.Project.CreatedAt,
			MaxUploadSize: assignment.Project.MaxUploadSize,
		})
	}

	SendJsonWithoutMeta(w, http.StatusOK, projectsData)
}
