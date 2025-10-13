package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	fmt.Printf("Config contents: %+v\n", cfg)

	currentState := state{cfg: &cfg}
	currentCommands := commands{make(map[string]func(*state, command) error)}

	currentCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	commandStruct := command{name: commandName, args: commandArgs}
	err = currentCommands.run(&currentState, commandStruct)
	if err != nil {
		log.Fatalf("could not run command: %v", err)
	}
	// err = cfg.SetUser("marqar")
	// if err != nil {
	// 	log.Fatalf("could not set new user: %v", err)
	// }

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	fmt.Printf("Config contents: %+v\n", cfg)
}
