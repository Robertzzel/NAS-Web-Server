package utils

import (
	"NAS-Server-Web/configurations"
	"crypto/sha256"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

/*
DATABASE FILE STRUCTURE:

<username>,<password sha256>
<username>,<password sha256>
<username>,<password sha256>
...
*/
func CheckUsernameAndPassword(username, password string) Result[bool] {
	contents, err := os.ReadFile(configurations.Database)
	if err != nil {
		return Error[bool](err)
	}

	passwordHash := hash(password)

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		usernameAndPassword := strings.Split(line, ",")

		if len(usernameAndPassword) == 2 && usernameAndPassword[0] == username && passwordHash == usernameAndPassword[1] {
			return Ok(true)
		}
	}

	return Ok(false)
}

func hash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
