package routes

import (
	"net/http"
)

func DownloadFileAttachmentGet(w http.ResponseWriter, r *http.Request) {
	//session, err := sessionService.GetSession(c)
	//if err != nil {
	//	return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	//}
	//
	//file := c.Param("file")
	//
	//filepath, err := filesService.GetFile(session, file)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, "'message': '"+err.Error()+"'")
	//}
	//
	//c.Response().Header().Set("Content-Disposition", "attachment; filename="+file)
	//
	//return c.File(filepath)
}
