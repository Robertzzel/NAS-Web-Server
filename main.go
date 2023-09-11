package main

import (
	. "NAS-Server-Web/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	})
	e.POST("/api/login", LoginPOST)
	e.POST("/api/list", ListPost)
	e.POST("/api/rm", RemovePost)
	e.GET("/api/dwat/:file", DownloadFileAttachmentGet)
	e.GET("/api/dwin/:file", DownloadFileInlineGet)
	e.GET("/api/dwdr/:file", DownloadDirectoryGet)
	e.POST("/api/upload", UploadFilesPost)
	e.POST("/api/directory", CreateDirectoryPost)
	e.POST("/api/rename", RenameFilePost)
	e.GET("/api/details", UserDetailsGet)

	e.Logger.Fatal(e.Start("localhost:33334"))
}
