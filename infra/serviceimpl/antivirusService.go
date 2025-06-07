package serviceimpl

import (
	chatservice "core_chat/application/chat/service"
	personservice "core_chat/application/person/service"
	"errors"
	"os"

	"github.com/dutchcoders/go-clamd"
)

// Ensure AntivirusServiceImpl implements both interfaces at compile time
var (
	_ personservice.AntivirusService = (*AntivirusServiceImpl)(nil)
	_ chatservice.AntivirusService   = (*AntivirusServiceImpl)(nil)
)

type AntivirusServiceImpl struct {
	clamd *clamd.Clamd
}

// Constructor for person service usecase
func NewPersonAntivirusService() personservice.AntivirusService {
	return &AntivirusServiceImpl{
		clamd: clamd.NewClamd("tcp://127.0.0.1:3310"),
	}
}

// Constructor for chat service usecase
func NewChatAntivirusService() chatservice.AntivirusService {
	return &AntivirusServiceImpl{
		clamd: clamd.NewClamd("tcp://127.0.0.1:3310"),
	}
}

// Shared method used by both interfaces
func (s *AntivirusServiceImpl) ScanImage(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("failed to open file: " + err.Error())
	}
	defer file.Close()

	response, err := s.clamd.ScanStream(file, make(chan bool))
	if err != nil {
		return errors.New("clamd scan error: " + err.Error())
	}

	if response == nil {
		return errors.New("clamd returned nil response")
	}

	for res := range response {
		if res.Status == clamd.RES_FOUND {
			return errors.New("virus detected: " + res.Description)
		}
	}

	return nil
}
