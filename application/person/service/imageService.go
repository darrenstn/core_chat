package service

type ImageService interface {
	ResizeImage(path string, dstDir string, width, height int) error
}
