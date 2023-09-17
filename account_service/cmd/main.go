package main

import "github.com/GideonLewis/crypto-analyzer/pkg/logger"

func main() {
	log := logger.NewLogger()
	log.GetLogger().Info("Hello world")
}
