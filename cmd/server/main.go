package main

import (
	"flag"
	"log"
	"os"
	"translateBot/internal/bot"
	"translateBot/internal/config"
)

var flagConfig = flag.String("config", "./config/dev.yml", "path to the config file")

func main() {

	flag.Parse()

	logger := log.Default()
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		logger.Fatal(err)
		os.Exit(-1)
	}
	bot, err := bot.New(*cfg, 40, true)
	if err != nil {
		logger.Fatal(err)
		os.Exit(-1)
	}
	bot.Listen(logger)
}
