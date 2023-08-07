package settings

import "NAS-Server-Web/models"

const (
	DatabasePath  = "/home/robert/Workspace/FTP-NAS-SV/database.db"
	BasePath      = "/home/robert/Downloads/NAS"
	MemoryPerUsed = 4 * 1024 * 1024 * 1024
)

var (
	Sessions = make(map[string]models.UserSession)
)
