package routes

import (
	"NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func RenameFilePost(c echo.Context) error {
	cookie, err := c.Cookie("ftp")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[cookie.Value]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	request := make(map[string]string)
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	oldName, oldNameExists := request["oldName"]
	newName, newNameExists := request["newName"]
	if !oldNameExists || !newNameExists {
		return c.JSON(http.StatusBadRequest, "'message': 'Bad parameters'")
	}

	oldName = userDetails.BasePath + oldName
	newName = userDetails.BasePath + newName

	if operations.IsPathSafe(oldName) && operations.IsPathSafe(newName) {
		return c.JSON(http.StatusBadRequest, "'message': 'Bad parameters'")
	}

	if err = os.Rename(oldName, newName); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	return c.JSON(http.StatusOK, "")
}
