package routes

import (
	"NAS-Server-Web/models"
	. "NAS-Server-Web/services/filesService"
	. "NAS-Server-Web/services/sessionService"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserDetailsGet(c echo.Context) error {
	session, err := GetSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	usedMemory, err := GetUserUsedMemory(session.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	user := models.UserMemoryDetails{Username: session.Username, Max: MemoryPerUsed, Used: usedMemory}

	res, err := json.Marshal(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	return c.JSONBlob(http.StatusOK, res)
}
