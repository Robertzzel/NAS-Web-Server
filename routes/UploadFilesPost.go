package routes

import (
	"net/http"
)

func UploadFilesPost(w http.ResponseWriter, r *http.Request) {
	//session, err := GetSession(c)
	//if err != nil {
	//	return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	//}
	//
	//request := c.Request()
	//if err := UploadFile(session, c.Param("name"), request.Body, request.ContentLength); err != nil {
	//	return c.JSON(http.StatusBadRequest, err.Error())
	//}
	//
	//return c.JSON(http.StatusOK, "")
}
