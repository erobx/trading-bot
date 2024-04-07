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

func (ls *LogService) GetSkin(name, wear string, price float32) (Skin, error) {
	ls.logger.Printf("Searching for %s...\n", name)
	skin, err := ls.svc.GetSkin(name, wear, price)
	if err != nil {
		ls.logger.Printf("Could not find: %s, %s, $%.2f\n", name, wear, price)
	} else {
		ls.logger.Printf("Found skin: %s, %s, $%.2f\n", name, wear, price)
	}
	return skin, err
}

func (ls *LogService) AddSkin(name, wear string, price float32) error {
	ls.logger.Printf("Attempting to add skin: %s, %s, $%.2f...\n", name, wear, price)
	skin := ls.svc.AddSkin(name, wear, price)
	if skin != nil {
		ls.logger.Printf("Could not list skin: %s, %s, $%.2f\n", name, wear, price)
		return skin
	}
	ls.logger.Printf("New skin on market: %s, %s, $%.2f\n", name, wear, price)
	return skin
}

func (ls *LogService) RemoveSkin(name, wear string, price float32) error {
	ls.logger.Printf("Removing skin: %s, %s...\n", name, wear)
	err := ls.svc.RemoveSkin(name, wear, price)
	if err != nil {
		ls.logger.Printf("Could not remove skin: %s, %s\n", name, wear)
		return err
	}
	ls.logger.Printf("Skin removed: %s, %s", name, wear)
	return err
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
