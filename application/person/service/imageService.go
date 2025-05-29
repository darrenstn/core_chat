package service

type ImageService interface {
	ResizeImage(path string, dstDir string, width, height int) (string, error)
	GetProfileImagePathOrDefault(path string, fallback string) string
}
