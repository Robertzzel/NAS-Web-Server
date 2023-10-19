package sessionService

import (
	"NAS-Server-Web/models"
	"NAS-Server-Web/services/configsService"
	"errors"
	"net/http"
	"os"
	"path"
)

var (
	sessions = make(map[string]models.UserSession)
)

func GetSession(cookie *http.Cookie) (models.UserSession, error) {
	session, exists := sessions[cookie.Value]
	if !exists {
		return models.UserSession{}, errors.New("session not found")
	}

	return session, nil
}

func NewSession(key string, username string) error {
	configs, err := configsService.NewConfigsService()
	if err != nil {
		return err
	}

	directory := path.Join(configs.GetBaseFilesPath(), username)
	if _, err := os.Stat(directory); err != nil {
		return err
	}
	sessions[key] = models.UserSession{BasePath: directory, Username: username}
	return nil
}
