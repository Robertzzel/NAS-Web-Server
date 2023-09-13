package routes

import (
	"NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func DownloadFileInlineGet(c echo.Context) error {
	session, err := sessionService.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	file := c.Param("file")

	file = session.BasePath + file

	fileInfo, err := os.Stat(file)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}

	if fileInfo.IsDir() {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	return c.Inline(file, file)
}
