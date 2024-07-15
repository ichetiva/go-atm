package main

import (
	"flag"
	"log"

	"github.com/ichetiva/go-atm/internal/config"
	"github.com/ichetiva/go-atm/internal/ports"
	"github.com/ichetiva/go-atm/pkg/logging"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "configPath", "config/config.yaml", "path to config YAML file")
	flag.Parse()
}

// @title ATM API
// @version 1.0
// @basePath /
// @host localhost:8000
func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("load config error", err)
	}

	logger := logging.NewLogger(cfg.LoggingLevel)
	logger.Info("Logger initialized")

	logger.Info("Config initialization")

	logger.Info("App initialization")
	app := ports.NewApp(logger, cfg)
	app.Run()
}
