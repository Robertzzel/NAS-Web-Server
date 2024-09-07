package utils

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Session struct {
	Expires  time.Time
	Username string
}

var (
	sessions = make(map[string]Session)
)

func VerifySession(request *http.Request) Maybe[Session] {
	sessionCookie, err := request.Cookie("ftp")
	if err != nil {
		return None[Session]()
	}

	session, sessionExists := sessions[sessionCookie.Value]
	if !sessionExists {
		return None[Session]()
	}

	sessionExpired := time.Now().After(session.Expires)
	if sessionExpired {
		return None[Session]()
	}

	return Some(session)
}

func CreateSession(username string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "ftp"
	cookie.Value = uuid.NewString()

	sessions[cookie.Value] = Session{Username: username, Expires: time.Now().Add(24 * time.Hour)}
	return cookie
}
