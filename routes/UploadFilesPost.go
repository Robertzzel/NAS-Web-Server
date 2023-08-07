package routes

import (
	. "NAS-Server-Web/operations"
	. "NAS-Server-Web/settings"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func UploadFilesPost(c echo.Context) error {
	cookie, err := c.Cookie("ftp")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}
	userDetails, hasPath := Sessions[cookie.Value]
	if !hasPath {
		return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	}

	filename := c.FormValue("filename")
	totalSizeParam := c.FormValue("size")
	totalSize, err := strconv.Atoi(totalSizeParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "'message': 'Bad parameters'")
	}
	remainingMemory, err := GetUserRemainingMemory(userDetails.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	}

	if remainingMemory < int64(totalSize) {
		return c.JSON(http.StatusBadRequest, "'message': 'No memory for the upload'")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if remainingMemory < file.Size {
		return c.JSON(http.StatusBadRequest, "'message': 'No memory for the upload'")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstPath := filepath.Join(userDetails.BasePath, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
}
