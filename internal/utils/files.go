package utils

import (
	"compress/flate"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CompressFile(file multipart.File, header *multipart.FileHeader) (string, error) {

	outputFileName := "compressed_" + filepath.Base(header.Filename)
	outFile, err := os.Create(outputFileName)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Create a new deflater writer using the best compression level
	deflater, err := flate.NewWriter(outFile, flate.BestCompression)
	if err != nil {
		return "", err
	}
	defer deflater.Close()

	// Copy the contents of the uploaded file into the deflater writer,
	// which compresses the data on the fly
	if _, err := io.Copy(deflater, file); err != nil {
		return "", err
	}

	// Flush any remaining data from the deflater
	if err := deflater.Flush(); err != nil {
		return "", err
	}

	return outputFileName, nil
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