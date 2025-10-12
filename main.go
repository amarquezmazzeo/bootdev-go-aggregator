package main

import (
	"fmt"
	"log"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	fmt.Printf("Config contents: %+v\n", cfg)

	err = cfg.SetUser("marqar")
	if err != nil {
		log.Fatalf("could not set new user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	fmt.Printf("Config contents: %+v\n", cfg)
}
