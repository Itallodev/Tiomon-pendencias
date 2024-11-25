package main

import (
	"Tiomon/Pendencias/config"
	"Tiomon/Pendencias/routes"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	routes.Initialize()

}
