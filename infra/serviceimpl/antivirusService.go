package serviceimpl

import (
	"core_chat/application/person/service"
	"errors"
	"os"

	"github.com/dutchcoders/go-clamd"
)

type AntivirusServiceImpl struct {
	clamd *clamd.Clamd
}

func NewAntivirusService() service.AntivirusService {
	// Connect to clamd via TCP (Docker exposes 3310)
	return &AntivirusServiceImpl{
		clamd: clamd.NewClamd("tcp://127.0.0.1:3310"),
	}
}

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
