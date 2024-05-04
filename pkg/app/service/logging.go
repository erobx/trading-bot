package service

import (
	"context"
	"log"
	"os"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/types"
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

func (ls *LogService) GetSkin(context context.Context, name, wear string) (model.Skin, error) {
	ls.logger.Printf("Searching for %s...\n", name)
	skin, err := ls.svc.GetSkin(context, name, wear)
	if err != nil {
		ls.logger.Printf("Could not find: %s, %s\n", name, wear)
	} else {
		ls.logger.Printf("Found skin: %s, %s, $%s\n", name, wear, skin.Price.String())
	}
	return skin, err
}

func (ls *LogService) AddSkin(context context.Context, name, wear string, price types.DbDecimal) error {
	ls.logger.Printf("Attempting to add skin: %s, %s, $%s...\n", name, wear, price.String())
	skin := ls.svc.AddSkin(context, name, wear, price)
	if skin != nil {
		ls.logger.Printf("Could not list skin: %s, %s, $%s\n", name, wear, price.String())
		return skin
	}
	ls.logger.Printf("New skin on market: %s, %s, $%s\n", name, wear, price.String())
	return skin
}

func (ls *LogService) RemoveSkin(context context.Context, name, wear string, price types.DbDecimal) error {
	ls.logger.Printf("Removing skin: %s, %s...\n", name, wear)
	err := ls.svc.RemoveSkin(context, name, wear, price)
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
