package handlers

import (
	"net/http"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

type RoleResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserApiResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
	Role      any    `json:"role"`
}

type UserCreateRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RoleID    int64  `json:"role_id"`
}

func HandleUserCreateRequest(w http.ResponseWriter, r *http.Request) {
	var user UserCreateRequest
	if err := ReadJson(w, r, &user); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	userModel := &store.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    user.RoleID,
		IsActive:  true,
	}

	if err := userModel.Password.Set(user.Password); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error setting password")
		return
	}

	if err := appStore.Users.Create(r.Context(), nil, userModel); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error creating user")
		return
	}

	userResponse := UserApiResponse{
		ID:        userModel.ID,
		Email:     userModel.Email,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		CreatedAt: userModel.CreatedAt,
		IsActive:  userModel.IsActive,
	}
	SendJsonWithoutMeta(w, http.StatusCreated, userResponse)
}

func HandleListUsersRequest(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	store := currentApp.Store

	limit, offset := GetPaginationParams(r)

	users, err := store.Users.GetAll(r.Context(), limit, offset)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error listing users")
		return
	}

	userCount, countErr := store.Users.Count(r.Context())
	if countErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error counting users")
		return
	}

	meta := JsonMeta{
		TotalRecords: userCount,
		Limit:        limit,
		Offset:       offset,
	}

	var userResponses []UserApiResponse
	for _, user := range users {
		userResponses = append(userResponses, UserApiResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt,
			IsActive:  user.IsActive,
			Role: RoleResponse{
				ID:          user.Role.ID,
				Name:        user.Role.Name,
				Description: user.Role.Description,
			},
		})
	}

	SendJson(w, http.StatusOK, userResponses, meta)
}

func HandleGetRolesRequest(w http.ResponseWriter, r *http.Request) {
	currentApp := app.GetCurrentApplication()
	store := currentApp.Store

	limit, offset := GetPaginationParams(r)

	roles, err := store.Roles.GetAll(r.Context(), limit, offset)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error listing roles")
		return
	}

	countRoles, countErr := store.Roles.Count(r.Context())
	if countErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, "error counting roles")
		return
	}

	meta := JsonMeta{
		TotalRecords: countRoles,
		Limit:        limit,
		Offset:       offset,
	}

	var roleResponses []RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	SendJson(w, http.StatusOK, roleResponses, meta)
}
