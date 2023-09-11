package routes

import (
	. "NAS-Server-Web/models"
	"NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ListPost(c echo.Context) error {
	session := operations.GetSession(c)
	if session == "" {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[session]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	pathDict := make(map[string]string)
	err := c.Bind(&pathDict)
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
		files, err := os.ReadDir(currentPath)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
		}

		var contents []FileDetails
		for _, file := range files {
			fileType, _ := GetFileType(filepath.Join(currentPath, file.Name()))
			fileDetails := FileDetails{Size: 0, Name: file.Name(), IsDir: file.IsDir(), Type: fileType}
			info, err := file.Info()
			if err == nil {
				fileDetails.Size = info.Size()
				fileDetails.CreatingTime = info.ModTime().Unix()
			}
			if strings.Contains(fileType, "image") {
				fileDetails.ImageData, err = Resize(filepath.Join(currentPath, file.Name()), 64, 64)
				if err != nil {
					fileDetails.ImageData = nil
				}
			}

			contents = append(contents, fileDetails)
		}

		var sendData []byte
		if contents != nil {
			sendData, err = json.Marshal(contents)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "'message': 'Server error'")
			}
		} else {
			sendData = []byte("")
		}

		return c.JSONBlob(http.StatusOK, sendData)
	}

	return c.JSON(http.StatusUnauthorized, "'message': 'Does not exist'")
}
