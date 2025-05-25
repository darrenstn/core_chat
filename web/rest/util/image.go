package util

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func IsValidImage(file multipart.File) (bool, string) {
	header := make([]byte, 512)
	_, err := file.Read(header)
	if err != nil {
		return false, ""
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false, ""
	}

	contentType := http.DetectContentType(header)
	switch contentType {
	case "image/jpeg", "image/png":
		return true, contentType
	default:
		return false, contentType
	}
}

func SaveImage(file multipart.File, name string, dir string, extension string) (string, error) {
	filename := name + extension

	uploadDir := os.Getenv("UPLOAD_DIR")
	savePath := filepath.Join(uploadDir, dir)

	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		return "", err
	}

	fullPath := filepath.Join(savePath, filename)

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}
