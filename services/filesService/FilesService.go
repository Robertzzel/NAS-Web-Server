package filesService

import (
	"NAS-Server-Web/models"
	"NAS-Server-Web/services/configsService"
	"archive/zip"
	"bytes"
	_ "encoding/json"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadFile(username, filename string, reader io.Reader, size int64) error {
	remainingMemory, err := GetUserRemainingMemory(username)
	if err != nil {
		return errors.New("internal error")
	}

	if remainingMemory < size {
		return errors.New("no memory for the upload")
	}

	if !IsPathSafe(filename) {
		return errors.New("bad path")
	}

	create, err := os.Create(filename)
	if err != nil {
		return errors.New("internal error")
	}

	_, err = io.Copy(create, reader)
	if err != nil {
		return errors.New("internal error")
	}

	return nil
}

func SendFile(filename, userDirPath string, w io.Writer) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return zipDirectory(filename, userDirPath, w)
	}

	fileHandler, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileHandler.Close()

	_, err = io.Copy(w, fileHandler)
	return err
}

func RemoveFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return errors.New("'message': 'File does not exist'")
	}

	if err := os.RemoveAll(filepath); err != nil {
		return errors.New("'message': 'Error on files removal'")
	}

	return nil
}

func RenameFile(oldPath, newPath string) error {
	if !IsPathSafe(oldPath) || !IsPathSafe(newPath) {
		return errors.New("bad parameters")
	}

	if err := os.Rename(oldPath, newPath); err != nil {
		return errors.New("internal error")
	}

	return nil
}

func GetUserUsedMemory(username string) (int64, error) {
	configs, err := configsService.NewConfigsService()
	if err != nil {
		return 0, err
	}

	entries, err := os.ReadDir(configs.GetBaseFilesPath())
	if err != nil {
		return 0, err
	}

	for _, dir := range entries {
		if dir.Name() != username {
			continue
		}
		info, err := dir.Info()
		if err != nil {
			return 0, err
		}
		dirSize, err := DirSize(configs.GetBaseFilesPath() + "/" + info.Name())
		if err != nil {
			return 0, err
		}
		return dirSize, nil
	}

	return 0, errors.New("username does not exist")
}

func GetUserRemainingMemory(username string) (int64, error) {
	configs, err := configsService.NewConfigsService()
	if err != nil {
		return 0, err
	}

	used, err := GetUserUsedMemory(username)
	if err != nil {
		return 0, err
	}
	return configs.GetMemoryPerUser() - used, nil
}

func DirSize(path string) (int64, error) {
	var dirSize int64 = 0

	readSize := func(path string, file os.FileInfo, err error) error {
		if file != nil && !file.IsDir() {
			dirSize += file.Size()
		}

		return nil
	}

	if err := filepath.Walk(path, readSize); err != nil {
		return 0, err
	}

	return dirSize, nil
}

func IsPathSafe(path string) bool {
	return !strings.Contains(path, "../")
}

func GetFilesFromDirectory(path string) ([]models.FileDetails, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("no directory with this path")
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contents []models.FileDetails
	for _, file := range files {
		fileType, _ := GetFileType(filepath.Join(path, file.Name()))
		fileDetails := models.FileDetails{Size: 0, Name: file.Name(), IsDir: file.IsDir(), Type: fileType}
		info, err := file.Info()
		if err == nil {
			fileDetails.Size = info.Size()
			fileDetails.CreatingTime = info.ModTime().Unix()
		}
		if strings.Contains(fileType, "image") {
			fileDetails.ImageData, err = Resize(filepath.Join(path, file.Name()), 64, 64)
			if err != nil {
				fileDetails.ImageData = nil
			}
		}

		contents = append(contents, fileDetails)
	}

	return contents, nil
}

func Resize(filepath string, width, height uint) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	newImage := resize.Resize(width, height, img, resize.Lanczos3)
	buffer := new(bytes.Buffer)

	if err = jpeg.Encode(buffer, newImage, nil); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func zipDirectory(inputDirectory, userDirPath string, outputWriter io.Writer) error {
	w := zip.NewWriter(outputWriter)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		inZipFile := strings.TrimPrefix(path, userDirPath+"/")
		f, err := w.Create(inZipFile)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	return filepath.Walk(inputDirectory, walker)
}

func GetFileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	mimeType := mime.TypeByExtension(filePath)
	if mimeType == "" {
		buffer := make([]byte, 512)
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return "", err
		}
		mimeType = http.DetectContentType(buffer[:n])
	}

	return mimeType, nil
}

//func Unzip(source, destination string) error {
//	reader, err := zip.OpenReader(source)
//	if err != nil {
//		return err
//	}
//	defer reader.Close()
//
//	destination, err = filepath.Abs(destination)
//	if err != nil {
//		return err
//	}
//
//	for _, f := range reader.File {
//		err := unzipFile(f, destination)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func unzipFile(f *zip.File, destination string) error {
//	filePath := filepath.Join(destination, f.Name)
//	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
//		return fmt.Errorf("invalid file path: %s", filePath)
//	}
//
//	if f.FileInfo().IsDir() {
//		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
//			return err
//		}
//		return nil
//	}
//
//	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
//		return err
//	}
//
//	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
//	if err != nil {
//		return err
//	}
//	defer destinationFile.Close()
//
//	zippedFile, err := f.Open()
//	if err != nil {
//		return err
//	}
//	defer zippedFile.Close()
//
//	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
//		return err
//	}
//	return nil
//}
