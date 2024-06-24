package main

import (
	"log"
	"smart-contracts-research/configs"
	"smart-contracts-research/internal/app"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}
	app.Run(cfg)
}
