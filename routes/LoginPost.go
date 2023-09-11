package routes

import (
	"NAS-Server-Web/database"
	"NAS-Server-Web/models"
	. "NAS-Server-Web/settings"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
	"time"
)

func LoginPOST(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	dbInstance, err := database.GetDatabase(DatabasePath)
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
	usersDisrectory := path.Join(BasePath, user.Username)
	if _, err = os.Stat(usersDisrectory); err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	cookie := new(http.Cookie)
	cookie.Name = "ftp"
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	Sessions[cookie.Value] = models.UserSession{BasePath: usersDisrectory, Username: user.Username}
	c.SetCookie(cookie)

	marshal, err := json.Marshal(cookie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}
	return c.JSONBlob(http.StatusOK, marshal)
}
