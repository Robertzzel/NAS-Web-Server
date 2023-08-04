package main

import (
	. "NAS-Server-Web/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowCredentials: true,
	}))

	e.POST("/api/login", LoginPOST)
	e.POST("/api/list", ListPost)
	e.POST("/api/rm", RemovePost)
	e.GET("/api/dwat/:file", DownloadFileAttachmentGet)
	e.GET("/api/dwin/:file", DownloadFileInlineGet)
	e.GET("/api/dwdr/:file", DownloadDirectoryGet)
	e.POST("/api/upload", UploadFilesPost)
	e.POST("/api/directory", CreateDirectoryPost)
	e.Logger.Fatal(e.Start(":8000"))
}
