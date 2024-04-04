package main

import (
	"log"
	"os"
)

type LogService struct {
	logger *log.Logger
	svc    Service
}

func NewLogService(svc Service) Service {
	return &LogService{
		logger: openLogFile("./logs/userInfo.log"),
		svc:    svc,
	}
}

func (ls *LogService) FindSkin(name string) *Skin {
	ls.logger.Printf("Searching for %s...\n", name)
	skin := ls.svc.FindSkin(name)
	if skin == nil {
		ls.logger.Printf("Could not find %s\n", name)
		return skin
	}
	ls.logger.Printf("Found skin %s\n", name)
	return skin
}

func (ls *LogService) ListSkin(name string, price float32) error {
	ls.logger.Printf("Attempting to add skin %s...\n", name)
	skin := ls.svc.ListSkin(name, price)
	if skin != nil {
		ls.logger.Printf("Could not list skin %s\n", name)
		return skin
	}
	ls.logger.Printf("New skin on market: %s, $%.2f\n", name, price)
	return skin
}

func openLogFile(path string) *log.Logger {
	// Recreate log file each run for testing
	os.Remove("./logs/userInfo.log")
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil
	}
	infoLog := log.New(logFile, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.SetOutput(logFile)

	return infoLog
}

func clearLog() {

}
