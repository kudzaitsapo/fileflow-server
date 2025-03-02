package handlers

import (
	"net/http"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
	"github.com/kudzaitsapo/fileflow-server/internal/utils"
)


func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10MB is the default used by FormFile
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		WriteJsonError(w, http.StatusBadRequest, "Unable to get file")
		return
	}
	defer file.Close()

	// get project key from the headers
	projectKey := r.Header.Get("project-key")
	if projectKey == "" {
		WriteJsonError(w, http.StatusBadRequest, "Project key is required")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	// validate project key by getting the project
	project, err := appStore.Projects.GetByKey(r.Context(), projectKey)
	if err != nil {
		WriteJsonError(w, http.StatusNotFound, "Project not found")
		return
	}

	// 1. Get file information => file name, size, etc.
	storedFile := &store.StoredFile{
		FileName: handler.Filename,
		FileSize: handler.Size,
		MimeType: handler.Header.Get("Content-Type"),
		Folder:   "uploads",
		SavedAs:  handler.Filename,
		OriginalExtension: utils.GetFileExtension(handler.Filename),
		ProjectID: project.ID,
	}
	// 2. Store the file in the database
	err = appStore.StoredFiles.Create(r.Context(), storedFile)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Unable to store file")
		return
	}
	// 3. Compress the file
	// 4. Return the file id to the client

	// Compress the file
	compressedFileName, err := utils.CompressFile(file, handler)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Unable to compress file")
		return
	}

	WriteJson(w, http.StatusOK, map[string]string{"filename": compressedFileName})
}