package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

func GetCurrentUser(r *http.Request) (*store.User, error) {
	// Get the request header auth token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization required to access")
	}

	// Expected format: "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, errors.New("invalid authorization header")
	}

	tokenStr := parts[1]

	// Validate the token
	currentApp := app.GetCurrentApplication()
	authenticator := currentApp.Authenticator
	claims, err := authenticator.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// get the user id from the jwt claims
	subject, err := claims.Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		return nil, err
	}

	// Get the user from the store
	store := currentApp.Store
	user, err := store.Users.GetById(r.Context(), userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetPaginationParams(r *http.Request) (int64, int64) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := int64(10)
	offset := int64(0)

	if limitStr != "" {
		limit, _ = strconv.ParseInt(limitStr, 10, 64)
	}

	if offsetStr != "" {
		offset, _ = strconv.ParseInt(offsetStr, 10, 64)
	}

	return limit, offset
}
