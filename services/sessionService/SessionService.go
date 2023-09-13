package sessionService

import (
	"NAS-Server-Web/models"
	. "NAS-Server-Web/settings"
	"errors"
	"github.com/labstack/echo/v4"
	"os"
	"path"
)

var (
	sessions = make(map[string]models.UserSession)
)

func GetSession(ctx echo.Context) (models.UserSession, error) {
	request := ctx.Request()
	authHeaders, exists := request.Header["Authentication"]
	if !exists || len(authHeaders) != 1 {
		return models.UserSession{}, errors.New("no authentication header found")
	}

	session, exists := sessions[authHeaders[0]]
	if !exists {
		return models.UserSession{}, errors.New("session not found")
	}

	return session, nil
}

func NewSession(key string, username string) error {
	directory := path.Join(BasePath, username)
	if _, err := os.Stat(directory); err != nil {
		return err
	}
	sessions[key] = models.UserSession{BasePath: directory, Username: username}
	return nil
}
