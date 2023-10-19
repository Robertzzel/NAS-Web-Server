package models

type FileDetails struct {
	Name         string
	Size         int64
	IsDir        bool
	Type         string
	ImageData    []byte
	CreatingTime int64
}
