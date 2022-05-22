package file

import (
	"mime/multipart"
	"net/http"
	"os"
)

func GetMimeTypeFromPath(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		os.Exit(1)
		return "", err
	}

	return GetMimeType(file)
}

func GetMimeTypeFromFileHeader(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	return GetMimeType(file)
}

func GetMimeType(file multipart.File) (string, error) {
	// Create a buffer to store the header of the file in
	buff := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(buff); err != nil {
		return "", err
	}
	defer file.Close()

	fType := http.DetectContentType(buff)
	return fType, nil
}
