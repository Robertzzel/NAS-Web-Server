package filesService

import (
	"NAS-Server-Web/models"
	. "NAS-Server-Web/settings"
	"archive/zip"
	"errors"
	"github.com/google/uuid"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func UploadFile(session models.UserSession, filename string, reader io.Reader, size int64) error {
	remainingMemory, err := GetUserRemainingMemory(session.Username)
	if err != nil {
		return errors.New("internal error")
	}

	if remainingMemory < size {
		return errors.New("no memory for the upload")
	}

	dstPath := filepath.Join(session.BasePath, filename)

	if !IsPathSafe(dstPath) {
		return errors.New("bad path")
	}

	create, err := os.Create(dstPath)
	if err != nil {
		return errors.New("internal error")
	}

	_, err = io.Copy(create, reader)
	if err != nil {
		return errors.New("internal error")
	}

	return nil
}

func GetFile(session models.UserSession, filename string) (string, error) {
	fullFilename := session.BasePath + filename
	if !IsPathSafe(fullFilename) {
		return "", errors.New("bad path")
	}

	fileInfo, err := os.Stat(fullFilename)
	if err != nil {
		return "", errors.New("file does not exist")
	}

	if fileInfo.IsDir() {
		outputPath := path.Join(session.BasePath, uuid.New().String())
		if err = zipDirectory(fullFilename, outputPath); err != nil {
			return "", errors.New("internal error")
		}
		return outputPath, nil
	}

	return fullFilename, nil
}

func RemoveFile(session models.UserSession, filepath string) error {
	fullFilePath := session.BasePath + filepath

	_, err := os.Stat(fullFilePath)
	if err != nil {
		return errors.New("'message': 'File does not exist'")
	}

	if err := os.RemoveAll(fullFilePath); err != nil {
		return errors.New("'message': 'Error on files removal'")
	}

	return nil
}

func RenameFile(session models.UserSession, oldPath, newPath string) error {
	oldPath = session.BasePath + oldPath
	newPath = session.BasePath + newPath

	if IsPathSafe(oldPath) && IsPathSafe(newPath) {
		return errors.New("bad parameters")
	}

	if err := os.Rename(oldPath, newPath); err != nil {
		return errors.New("internal error")
	}

	return nil
}

func GetUserUsedMemory(username string) (int64, error) {
	entries, err := os.ReadDir(BasePath)
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
		dirSize, err := DirSize(BasePath + "/" + info.Name())
		if err != nil {
			return 0, err
		}
		return dirSize, nil
	}

	return 0, errors.New("username does not exist")
}

func GetUserRemainingMemory(username string) (int64, error) {
	used, err := GetUserUsedMemory(username)
	if err != nil {
		return 0, err
	}
	return MemoryPerUsed - used, nil
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

func zipDirectory(inputDirectory string, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zip.NewWriter(file)
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

		f, err := w.Create(path)
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
