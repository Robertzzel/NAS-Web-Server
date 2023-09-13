package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
	"strings"
)

func CreateDirectoryPost(c echo.Context) error {
	session, err := sessionService.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	pathDict := make(map[string]string)
	err = c.Bind(&pathDict)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	dirPath, pathExists := pathDict["path"]
	if !pathExists {
		return c.JSON(http.StatusBadRequest, "'message': 'no path provided'")
	}
	dirPath = path.Join(session.BasePath + dirPath)
	if !strings.HasPrefix(dirPath, "/") && filesService.IsPathSafe(dirPath) {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	if err = os.Mkdir(dirPath, 0770); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	return c.JSON(http.StatusOK, "")
}
