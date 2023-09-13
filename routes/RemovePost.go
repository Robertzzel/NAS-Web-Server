package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func RemovePost(c echo.Context) error {
	session, err := sessionService.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	pathDict := make(map[string]string)
	if err = c.Bind(&pathDict); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	currentPath, pathExists := pathDict["path"]
	if !pathExists || !strings.HasPrefix(currentPath, "/") {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	if err = filesService.RemoveFile(session, currentPath); err != nil {
		return c.JSON(http.StatusBadRequest, "'message': 'cannot delete file'")
	}

	return c.JSON(http.StatusOK, "")
}
