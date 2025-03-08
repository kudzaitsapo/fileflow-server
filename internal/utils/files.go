package utils

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// customReadCloser combines an io.Reader with a custom close function
type customReadCloser struct {
	io.Reader
	closeFunc func() error
}

func CompressAndSaveFile(file multipart.File, savedFileName string, saveFolder string) error {

	outputFileName := filepath.Base(savedFileName)

	if _, dirErr := os.Stat("uploads"); os.IsNotExist(dirErr) {
        os.Mkdir("uploads", 0755)
    }

	if saveFolder != "" {
		if _, dirCreateErr := os.Stat(filepath.Join("uploads", saveFolder)); os.IsNotExist(dirCreateErr) {
			createDirErr := os.MkdirAll(filepath.Join("uploads", saveFolder), 0755)
			if createDirErr != nil {
				log.Printf("Error creating directory: %v", createDirErr)
				return createDirErr
			}
		}
		outputFileName = filepath.Join("uploads", saveFolder, outputFileName)
	} else {
		outputFileName = filepath.Join("uploads", outputFileName)
	}

	// Read file content (already an io.ReadCloser)
    content, err := io.ReadAll(file)
    if err != nil {
        return err
    }
    defer file.Close()

    // Compress using deflate
    var b bytes.Buffer
    w, flateErr := flate.NewWriter(&b, flate.BestCompression)
	if flateErr != nil {
		return flateErr
	}

    if _, err := w.Write(content); err != nil {
        return err
    }
    if err := w.Close(); err != nil {
        return err
    }
    compressed := b.Bytes()

    return os.WriteFile(outputFileName, compressed, 0644)

}

func DecompressFile(compressedFileName string, folder string) (*os.File, error) {
	// Create a temporary file for the decompressed output
	fileNameWithFolder := filepath.Base(compressedFileName)
	if folder != "" {
		fileNameWithFolder = filepath.Join("uploads", folder, fileNameWithFolder)
	} else {
		fileNameWithFolder = filepath.Join("uploads", fileNameWithFolder)
	}


	tempDir := os.TempDir()

	// Ensure the temp directory exists.
	if _, tmpDirCreateErr := os.Stat(tempDir); os.IsNotExist(tmpDirCreateErr) {
		tmpDirCreateErr = os.MkdirAll(tempDir, 0755)
		if tmpDirCreateErr != nil {
			return nil, tmpDirCreateErr
		}
	}

	outputFile, err := os.CreateTemp(tempDir, "decompressed-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}

	// Open the compressed file
	compressedFile, err := os.Open(fileNameWithFolder)
	if err != nil {
		outputFile.Close()
		os.Remove(outputFile.Name())
		return nil, fmt.Errorf("failed to open compressed file: %w", err)
	}
	defer compressedFile.Close()

	// Create a flate reader
	flateReader := flate.NewReader(compressedFile)
	defer flateReader.Close()

	// Copy decompressed data to the output file
	_, err = io.Copy(outputFile, flateReader)
	if err != nil {
		outputFile.Close()
		os.Remove(outputFile.Name())
		return nil, fmt.Errorf("decompression failed: %w", err)
	}

	// Seek to the beginning of the file so it can be read from the start
	_, err = outputFile.Seek(0, 0)
	if err != nil {
		outputFile.Close()
		os.Remove(outputFile.Name())
		return nil, fmt.Errorf("failed to seek to beginning of file: %w", err)
	}

	return outputFile, nil
}

// DecompressFileAndReturnStream decompresses a file and returns a ReadCloser.
// This is an alternative implementation that returns a stream instead of a file.
func DecompressFileAndReturnStream(compressedFileName string, folder string) (io.ReadCloser, error) {
	// Open the compressed file

	fileNameWithFolder := filepath.Base(compressedFileName)
	if folder != "" {
		fileNameWithFolder = filepath.Join("uploads", folder, fileNameWithFolder)
	} else {
		fileNameWithFolder = filepath.Join("uploads", fileNameWithFolder)
	}

	compressedFile, err := os.Open(fileNameWithFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to open compressed file: %w", err)
	}

	// Create a flate reader - this is the stream that will be returned
	flateReader := flate.NewReader(compressedFile)

	// Create a custom ReadCloser that closes both the flateReader and the compressedFile
	return &customReadCloser{
		Reader: flateReader,
		closeFunc: func() error {
			err1 := flateReader.Close()
			err2 := compressedFile.Close()
			if err1 != nil {
				return err1
			}
			return err2
		},
	}, nil
}


func (c *customReadCloser) Close() error {
	return c.closeFunc()
}


func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}