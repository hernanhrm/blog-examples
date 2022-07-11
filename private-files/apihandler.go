package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	useCase FileManager
}

func newHandler(useCase FileManager) *handler {
	return &handler{useCase: useCase}
}

func (h handler) Upload(c echo.Context) error {
	m := RequestUploaderImage{}

	err := c.Bind(&m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	m.File, err = c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fileName, err := h.useCase.Upload(m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message":   "File uploaded!",
		"file_name": fileName,
	})
}

func (h handler) GetFile(c echo.Context) error {
	fileDetail, err := h.useCase.GetFile(c.QueryParam("filepath"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Blob(http.StatusOK, fileDetail.ContentType, fileDetail.FileBytes)
}

func (h handler) GetEmployees(c echo.Context) error {
	response, err := h.useCase.GetEmployees()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h handler) Presign(c echo.Context) error {
	signedURL, err := h.useCase.Presign(c.QueryParam("filepath"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"signed_url": signedURL,
	})
}
