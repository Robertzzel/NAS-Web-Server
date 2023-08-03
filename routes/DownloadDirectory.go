package routes

import (
	"NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
)

func DownloadDirectoryGet(c echo.Context) error {
	cookie, err := c.Cookie("ftp")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[cookie.Value]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	file := c.Param("file")
	fullFilename := path.Join(userDetails.BasePath, file)

	fileInfo, err := os.Stat(fullFilename)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}
	if !fileInfo.IsDir() {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}

	outputPath := path.Join(userDetails.BasePath, uuid.New().String())
	if err = operations.Zip(fullFilename, outputPath); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Server error on zip'")
	}
	defer func() {
		_ = os.Remove(outputPath)
	}()

	return c.File(outputPath)
}
