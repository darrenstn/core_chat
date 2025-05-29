package serviceimpl

import (
	"core_chat/application/person/service"
	"errors"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type ImageServiceImpl struct{}

func NewImageService() service.ImageService {
	return &ImageServiceImpl{}
}

// ResizeImage resizes an image at the given path to the specified width and height.
func (s *ImageServiceImpl) ResizeImage(path string, dstDir string, width, height int) (string, error) {
	// Open the image file
	src, err := imaging.Open(path)
	if err != nil {
		return "", errors.New("failed to open image: " + err.Error())
	}

	// Resize and overwrite
	dst := imaging.Resize(src, width, height, imaging.Lanczos)

	filename := filepath.Base(path)

	uploadDir := os.Getenv("UPLOAD_DIR")

	dstPath := filepath.Join(uploadDir, dstDir, filename)

	if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
		return "", errors.New("failed to create destination directory: " + err.Error())
	}

	// Save the resized image (overwrite)
	err = imaging.Save(dst, dstPath)
	if err != nil {
		return "", errors.New("failed to save resized image: " + err.Error())
	}

	return dstPath, nil
}

func (s *ImageServiceImpl) GetProfileImagePathOrDefault(path string, fallback string) string {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return path
	}
	return fallback
}
