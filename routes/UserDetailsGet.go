package routes

import (
	"NAS-Server-Web/models"
	. "NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserDetailsGet(c echo.Context) error {
	cookie, err := c.Cookie("ftp")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[cookie.Value]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	usedMemory, err := GetUserUsedMemory(userDetails.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	user := models.UserMemoryDetails{Username: userDetails.Username, Max: MemoryPerUsed, Used: usedMemory}

	res, err := json.Marshal(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	return c.JSONBlob(http.StatusOK, res)
}
