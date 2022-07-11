package main

import "mime/multipart"

const (
	MIMEImageJPEG      = "image/jpeg"
	MIMEImageJPG       = "image/jpg"
	MIMEImagePNG       = "image/png"
	MIMEApplicationPDF = "application/pdf"
	MIMETextPlain      = "text/plain"
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

type Employee struct {
	Name             string `json:"name"`
	Picture          string `json:"picture"`
	PreSignedPicture string `json:"pre_signed_picture"`
}

type Employees []Employee

// MOCK DATA FOR THE EXAMPLE

var employees = Employees{
	{
		Name:    "Hernan Reyes",
		Picture: "employees/149d6d2d-6adb-464a-ad97-1927cffce309.jpg",
	},
	{
		Name:    "Juan Perez",
		Picture: "employees/88c5d68b-4cac-466e-bb81-c0c4deb49534.jpg",
	},
	{
		Name:    "Pedro Perez",
		Picture: "employees/1399d998-b505-4ce2-b49b-bd8073de9b9e.jpg",
	},
}
