package main

import (
	"flag"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("api_recorder")
var plain_format = logging.MustStringFormatter(`%{time:2006-02-01 15:04:05.000} %{level:.4s} %{message}`)

func main() {
	log.Info("Starting API Recoder")
	config_flag := flag.String("config", "", "Path to config file")
	flag.Parse()
	var config Config

	if *config_flag == "" {
		config = DefaultConfig()
		log.Info("Using default config")
	} else {
		var err error

		config, err = LoadConfig(*config_flag)
		if err != nil {
			log.Fatalf("Error loading config file: %v", err)
		}
	}

	rm := NewRequestsManager(config)
	rm.Run()
}
