package routes

import (
	"net/http"
)

func UserDetailsGet(w http.ResponseWriter, r *http.Request) {
	//configs, err := configsService.NewConfigsService()
	//if err != nil {
	//	return err
	//}
	//
	//session, err := GetSession(c)
	//if err != nil {
	//	return c.JSON(http.StatusUnauthorized, "'message': 'You are not logged in'")
	//}
	//
	//usedMemory, err := GetUserUsedMemory(session.Username)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	//}
	//user := models.UserMemoryDetails{Username: session.Username, Max: configs.GetMemoryPerUser(), Used: usedMemory}
	//
	//res, err := json.Marshal(user)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, "'message': 'Internal error'")
	//}
	//
	//return c.JSONBlob(http.StatusOK, res)
}
