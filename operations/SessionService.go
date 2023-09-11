package operations

import (
	"github.com/labstack/echo/v4"
)

func GetSession(ctx echo.Context) string {
	request := ctx.Request()
	authHeaders, exists := request.Header["Authentication"]
	if !exists || len(authHeaders) != 1 {
		return ""
	}

	return authHeaders[0]
}
