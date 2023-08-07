package models

type UserMemoryDetails struct {
	Username string `json:"username"`
	Used     int64  `json:"used"`
	Max      int64  `json:"max"`
}
