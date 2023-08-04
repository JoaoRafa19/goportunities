package main

import (
	"github.com/JoaoRafa19/goplaning/config"
	"github.com/JoaoRafa19/goplaning/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize configs 
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrF("Config inicialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}