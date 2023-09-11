package routes

import (
	"NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func DownloadFileInlineGet(c echo.Context) error {
	session := operations.GetSession(c)
	if session == "" {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[session]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	file := c.Param("file")

	file = userDetails.BasePath + file

	fileInfo, err := os.Stat(file)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}

	if fileInfo.IsDir() {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	return c.Inline(file, file)
}
