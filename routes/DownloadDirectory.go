package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func DownloadDirectoryGet(c echo.Context) error {
	session, err := sessionService.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	file := c.Param("file")

	filepath, err := filesService.GetFile(session, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "'message': '"+err.Error()+"'")
	}
	defer func() {
		_ = os.Remove(filepath)
	}()

	return c.File(filepath)
}
