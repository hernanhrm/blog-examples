package main

import "mime/multipart"

const (
	MIMEImageJPEG      = "image/jpeg"
	MIMEImageJPG       = "image/jpg"
	MIMEImagePNG       = "image/png"
	MIMEApplicationPDF = "application/pdf"
)

// RequestUploaderImage is the struct for file upload
type RequestUploaderImage struct {
	File     *multipart.FileHeader `form:"file"`
	Folder   string                `form:"folder"`
	IsPublic bool                  `form:"is_public"`
}

type GetFileResponse struct {
	FileBytes   []byte
	ContentType string
}
