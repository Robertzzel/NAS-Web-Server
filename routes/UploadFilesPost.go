package routes

import (
	. "NAS-Server-Web/services/filesService"
	. "NAS-Server-Web/services/sessionService"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UploadFilesPost(c echo.Context) error {
	session, err := GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	request := c.Request()
	if err := UploadFile(session, c.Param("name"), request.Body, request.ContentLength); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
