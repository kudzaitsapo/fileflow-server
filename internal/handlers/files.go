package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
	"github.com/kudzaitsapo/fileflow-server/internal/utils"
)


func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data

	// TODO: Set the maximum limit on the project
	err := r.ParseMultipartForm(10 << 20) // 10MB is the default used by FormFile
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, parseErr := r.FormFile("file")
	// get folder from form data
	folder := r.FormValue("folder")
	storedFileName := uuid.New().String() + ".ffs"

	if parseErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Unable to get file")
		return
	}
	defer file.Close()

	// get project key from the headers
	projectKey := r.Header.Get("ff-project-key")
	if projectKey == "" {
		WriteJsonError(w, http.StatusBadRequest, "Project key is required")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	// validate project key by getting the project
	project, projErr := appStore.Projects.GetByKey(r.Context(), projectKey)
	if projErr != nil {
		WriteJsonError(w, http.StatusNotFound, "Project not found for key: "+projectKey)
		return
	}

	// TODO: File type validation 

	// 1. Get file information => file name, size, etc.
	storedFile := &store.StoredFile{
		FileName: handler.Filename,
		FileSize: handler.Size,
		MimeType: handler.Header.Get("Content-Type"),
		Folder:   folder,
		SavedAs:  storedFileName,
		OriginalExtension: utils.GetFileExtension(handler.Filename),
		ProjectID: project.ID,
	}
	// 2. Store the file in the database
	storErr := appStore.StoredFiles.Create(r.Context(), storedFile)
	if storErr != nil {
		log.Printf("Error storing file: %v", storErr)
		WriteJsonError(w, http.StatusInternalServerError, "Unable to store file")
		return
	}

	// 3. Compress the file
	saveErr := utils.CompressAndSaveFile(file, storedFileName, folder)
	if saveErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Unable to save file")
		return
	}

	// 4. Return the stored file to the client
	WriteJson(w, http.StatusOK, storedFile)
}