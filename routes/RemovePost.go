package routes

import (
	"net/http"
)

func RemovePost(w http.ResponseWriter, r *http.Request) {
	//session, err := sessionService.GetSession(c)
	//if err != nil {
	//	return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	//}
	//
	//pathDict := make(map[string]string)
	//if err = c.Bind(&pathDict); err != nil {
	//	return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	//}
	//
	//currentPath, pathExists := pathDict["path"]
	//if !pathExists || !strings.HasPrefix(currentPath, "/") {
	//	return c.JSON(http.StatusUnauthorized, "'message': 'You have no access'")
	//}
	//
	//currentPath = path.Join(session.BasePath, currentPath)
	//if err = filesService.RemoveFile(currentPath); err != nil {
	//	return c.JSON(http.StatusBadRequest, "'message': 'cannot delete file'")
	//}
	//
	//return c.JSON(http.StatusOK, "")
}
