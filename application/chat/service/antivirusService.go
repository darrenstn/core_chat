package service

type AntivirusService interface {
	ScanImage(filePath string) error
}
