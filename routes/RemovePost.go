package routes

import (
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func RemovePost(c echo.Context) error {
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

	currentPath, pathExists := pathDict["path"]
	if !pathExists || !strings.HasPrefix(currentPath, "/") {
		return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	}

	currentPath = userDetails.BasePath + currentPath

	fileInfo, err := os.Stat(currentPath)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
	}

	if fileInfo.IsDir() {
		if err := os.RemoveAll(currentPath); err != nil {
			return c.JSON(http.StatusInternalServerError, "'message': 'Error on files removal'")
		}
	} else {
		if err := os.Remove(currentPath); err != nil {
			return c.JSON(http.StatusInternalServerError, "'message': 'Error on files removal'")
		}
	}

	return c.JSON(http.StatusOK, "")
}
