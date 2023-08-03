package settings

import "NAS-Server-Web/models"

const (
	DatabasePath = "/home/robert/Workspace/FTP-NAS-SV/database.db"
)

var (
	Sessions = make(map[string]models.UserDetails)
)
