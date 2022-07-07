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
