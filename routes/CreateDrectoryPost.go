package routes

import (
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
	"strings"
)

func CreateDirectoryPost(c echo.Context) error {
	cookie, err := c.Cookie("ftp")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[cookie.Value]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	pathDict := make(map[string]string)
	err = c.Bind(&pathDict)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	dirPath, pathExists := pathDict["path"]
	if !pathExists || !strings.HasPrefix(dirPath, "/") {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}
	dirPath = path.Join(userDetails.BasePath + dirPath)

	if err = os.Mkdir(dirPath, 0770); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	return c.JSON(http.StatusOK, "")
}
