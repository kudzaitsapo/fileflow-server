package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
	"github.com/kudzaitsapo/fileflow-server/internal/utils"
)

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data

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
		WriteJsonError(w, http.StatusNotFound, fmt.Sprintf("Project not found for key: %s", projectKey))
		return
	}

	// validate upload size based on project settings
	err := r.ParseMultipartForm(project.MaxUploadSize << 20)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to upload file: %s", err), http.StatusBadRequest)
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

	// File type validation based on project settings
	isAllowed, validationErr := appStore.ProjectAllowedFileTypes.FileTypeIsAllowed(r.Context(), project.ID, handler.Header.Get("Content-Type"))

	if !isAllowed || validationErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "File type not allowed")
		return
	}

	// Assign Icons based on file type
	fileType, fileTypeRetrievalErr := appStore.FileTypes.GetByMimeType(r.Context(), handler.Header.Get("Content-Type"))
	var fileIcon string
	if fileTypeRetrievalErr != nil {
		fileIcon = ""
	} else {
		if fileType != nil {
			fileIcon = fileType.Icon
		} else {
			fileIcon = ""
		}
	}

	// 1. Get file information => file name, size, etc.
	storedFile := &store.StoredFile{
		FileName:          handler.Filename,
		FileSize:          handler.Size,
		MimeType:          handler.Header.Get("Content-Type"),
		Folder:            folder,
		SavedAs:           storedFileName,
		OriginalExtension: utils.GetFileExtension(handler.Filename),
		ProjectID:         project.ID,
		Icon:              fileIcon,
	}
	// 2. Store the file in the database
	storErr := appStore.StoredFiles.Create(r.Context(), storedFile)
	if storErr != nil {
		log.Printf("Error storing file: %v", storErr)
		WriteJsonError(w, http.StatusInternalServerError, "Unable to store file")
		return
	}

	// 3. Compress the file and save it
	saveErr := utils.CompressAndSaveFile(file, storedFileName, folder)
	if saveErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to save file: %s", saveErr))
		return
	}

	// 4. Return the stored file to the client
	WriteJson(w, http.StatusOK, storedFile)
}

func HandleFileDownload(w http.ResponseWriter, r *http.Request) {
	// Get the file ID from the URL
	fileID := r.PathValue("id")
	if fileID == "" {
		WriteJsonError(w, http.StatusBadRequest, "File ID is required")
		return
	}

	// get project key from the headers
	projectKey := r.Header.Get("ff-project-key")
	if projectKey == "" {
		WriteJsonError(w, http.StatusBadRequest, "Project key is required")
		return
	}

	uuidFileId, convErr := uuid.Parse(fileID)
	if convErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	// Get the file from the database
	storedFile, storErr := appStore.StoredFiles.GetByIdAndProjectKey(r.Context(), uuidFileId, projectKey)
	if storErr != nil {
		WriteJsonError(w, http.StatusNotFound, fmt.Sprintf("Unable to find file with id: %s", fileID))
		return
	}

	// Decompress the file
	// TODO: implement a way to choose between file based and stream based file serving
	// depending on file size
	filePath, decompressErr := utils.DecompressFile(storedFile.SavedAs, storedFile.Folder)
	if decompressErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to decompress file: %s", decompressErr))
		return
	}

	// Return the file to the client
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", storedFile.FileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", storedFile.FileSize))
	w.Header().Set("Content-Type", storedFile.MimeType)

	uploadedAt, err := time.Parse(time.RFC3339, storedFile.UploadedAt)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Invalid upload time format")
		return
	}
	http.ServeContent(w, r, storedFile.FileName, uploadedAt, filePath)
}

func HandleFilesList(w http.ResponseWriter, r *http.Request) {
	// get project key from the headers
	projectId := r.PathValue("id")
	if projectId == "" {
		WriteJsonError(w, http.StatusBadRequest, "Project id is required")
		return
	}

	intProjectId, convErr := strconv.ParseInt(projectId, 10, 64)
	if convErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	// get query params for limit and offset
	limit, offset := GetPaginationParams(r)

	// Get the files from the database
	storedFiles, storErr := appStore.StoredFiles.GetAllByProjectId(r.Context(), intProjectId, limit, offset)
	if storErr != nil {
		WriteJsonError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get files: %v", storErr))
		return
	}

	WriteJson(w, http.StatusOK, storedFiles)
}

func HandleFileInfo(w http.ResponseWriter, r *http.Request) {
	// get project key from the headers
	fileId := r.PathValue("id")
	if fileId == "" {
		WriteJsonError(w, http.StatusBadRequest, "File id is required")
		return
	}

	uuidFileId, convErr := uuid.Parse(fileId)
	if convErr != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid file ID")
		return
	}

	currentApp := app.GetCurrentApplication()
	appStore := currentApp.Store

	// Get the file from the database
	storedFile, storErr := appStore.StoredFiles.GetById(r.Context(), uuidFileId)
	if storErr != nil {
		WriteJsonError(w, http.StatusNotFound, fmt.Sprintf("Unable to find file with id: %s", fileId))
		return
	}

	// Return the file to the client
	WriteJson(w, http.StatusOK, storedFile)
}
