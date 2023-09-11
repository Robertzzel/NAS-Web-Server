package routes

import (
	"NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func DownloadFileAttachmentGet(c echo.Context) error {
	session := operations.GetSession(c)
	if session == "" {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[session]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	file := c.Param("file")

	fullFilename := userDetails.BasePath + file

	fileInfo, err := os.Stat(fullFilename)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}

	if fileInfo.IsDir() {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename="+file)

	return c.File(fullFilename)
}
