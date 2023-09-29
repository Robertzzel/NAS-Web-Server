package routes

import (
	"NAS-Server-Web/models"
	"NAS-Server-Web/services/databaseService"
	"NAS-Server-Web/services/sessionService"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func LoginPOST(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	dbInstance, err := databaseService.NewDatabaseService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	ok, err := dbInstance.Login(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	if !ok {
		return c.JSON(http.StatusUnauthorized, "'message': 'Wrong username or password'")
	}

	cookie := new(http.Cookie)
	cookie.Name = "ftp"
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)

	if err := sessionService.NewSession(cookie.Value, user.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	c.SetCookie(cookie)
	marshal, err := json.Marshal(cookie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	return c.JSONBlob(http.StatusOK, marshal)
}
