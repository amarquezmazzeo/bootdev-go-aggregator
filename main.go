package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/config"
	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg       *config.Config
	dbQueries *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	log.Printf("config contents: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	dbQueries := database.New(db)

	currentState := state{cfg: &cfg, dbQueries: dbQueries}
	currentCommands := commands{make(map[string]func(*state, command) error)}

	currentCommands.register("login", handlerLogin)
	currentCommands.register("register", handlerRegister)
	currentCommands.register("reset", handlerReset)

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

	// cfg, err = config.Read()
	// if err != nil {
	// 	log.Fatalf("could not read config file: %v", err)
	// }
	// log.Printf("config contents: %+v\n", cfg)
}
