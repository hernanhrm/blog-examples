package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/satori/uuid"
)

var allowedImages = map[string]string{
	MIMEImagePNG:       ".png",
	MIMEImageJPG:       ".jpg",
	MIMEImageJPEG:      ".jpg",
	MIMEApplicationPDF: ".pdf",
}

type FileManager struct {
	service Service
}

func NewFileManager(service Service) FileManager {
	return FileManager{service: service}
}

// Upload uploads files to a service
func (u FileManager) Upload(m RequestUploaderImage) (string, error) {
	fileReader, err := m.File.Open()
	if err != nil {
		return "", fmt.Errorf("filemanager.readFile(): %w", err)
	}
	defer fileReader.Close()

	file, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return "", fmt.Errorf("filemanager.ioutil.ReadAll(): %w", err)
	}

	contentType := http.DetectContentType(file)
	fileExt, ok := allowedImages[contentType]
	if !ok {
		return "", fmt.Errorf("filemanager.allowedImages: content type %s not allowed", contentType)
	}

	fileName, err := getFileName(fileExt)
	if err != nil {
		return "", fmt.Errorf("filemanager.getFileName(): %w", err)
	}

	path := filepath.Join(m.Folder, fileName)
	if err := u.service.Upload(file, contentType, path, m.IsPublic); err != nil {
		return "", fmt.Errorf("filemanager.Upload(): %w", err)
	}

	return path, nil
}

func (u FileManager) GetFile(filepath string) (GetFileResponse, error) {
	response, err := u.service.GetFile(filepath)
	if err != nil {
		return GetFileResponse{}, fmt.Errorf("filemanager.GetFile(): %w", err)
	}

	return response, nil
}

func getFileName(extension string) (string, error) {
	return uuid.NewV4().String() + extension, nil
}
