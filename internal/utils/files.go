package utils

import (
	"bytes"
	"compress/flate"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CompressAndSaveFile(file multipart.File, savedFileName string, saveFolder string) error {

	outputFileName := filepath.Base(savedFileName)

	if _, dirErr := os.Stat("uploads"); os.IsNotExist(dirErr) {
        os.Mkdir("uploads", 0755)
    }

	if saveFolder != "" {
		if _, dirErr := os.Stat(filepath.Join("uploads", saveFolder)); os.IsNotExist(dirErr) {
			os.Mkdir(filepath.Join("uploads", saveFolder), 0755)
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
    defer file.Close() // Ensure the file is closed after reading

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

func DecompressFile(inputFileName string) (string, error) {

	// Open the compressed file for reading.
	inFile, err := os.Open(inputFileName)
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	// Create a new flate reader to decompress the data.
	reader := flate.NewReader(inFile)
	defer reader.Close()

	// Create the output file name by prefixing with "decompressed_".
	outputFileName := "decompressed_" + filepath.Base(inputFileName)
	outFile, err := os.Create(outputFileName)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Copy the decompressed data to the output file.
	if _, err := io.Copy(outFile, reader); err != nil {
		return "", err
	}

	return outputFileName, nil
}

func SaveCompressedFile(file multipart.File) (*os.File, error) {

	return nil, nil
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}