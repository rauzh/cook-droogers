package main

import (
	"context"
	"cookdroogers/app"
	"cookdroogers/config"
	"cookdroogers/pkg/logger"
	"log/slog"
)

func main() {

	runApplication(&logger.LoggerFactorySlog{})

}

func runApplication(loggerFactory logger.LoggerFactory) {

	ctx := context.Background()
	log := loggerFactory.Logger(ctx)

	appConfig := config.ParseConfig()
	if appConfig == nil {
		log.Error("Failed to parse config")
		return
	}

	cdApp := app.App{Config: appConfig}

	err := cdApp.Init(log)
	if err != nil {
		log.Error("Failed to initialize app: ", slog.Any("error", err))
		return
	}

	// switch cdApp.Config.Mode {
	// case "techUI":
	// 	for {
	// 		err := techUI.RunMenu(&cdApp, log)
	// 		if errors.Is(err, techUI.ErrEXIT) {
	// 			break
	// 		}
	// 		if err != nil {
	// 			log.Error("Error running techUI: ", slog.Any("error", err))
	// 		}
	// 	}
	// default:
	// 	log.Info("Unknown mode")
	// }
}
