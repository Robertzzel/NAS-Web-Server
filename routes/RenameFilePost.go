package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
	"path"
)

func RenameFilePost(c echo.Context) error {
	session, err := sessionService.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	request := make(map[string]string)

	if err = c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	oldName, oldNameExists := request["oldName"]
	newName, newNameExists := request["newName"]
	if !oldNameExists || !newNameExists {
		return c.JSON(http.StatusBadRequest, "'message': 'Bad parameters'")
	}

	oldName = path.Join(session.BasePath, oldName)
	newName = path.Join(session.BasePath, newName)

	if err = filesService.RenameFile(oldName, newName); err != nil {
		return c.JSON(http.StatusBadRequest, "'message': 'cannot rename file'")
	}

	return c.JSON(http.StatusOK, "")
}
