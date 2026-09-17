package main

import (
	"auth/internal/api"
	"auth/internal/config"
	"log"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Cannot read config file: %v", err)
	}

	server := api.NewServer(conf.GRPCServer)

	err = server.Serve()
	if err != nil {
		log.Fatalf("Server error %v\n", err)
	}
}
