package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login handler expects a single argument (User Name)")
	}
	userName := cmd.args[0]
	_, err := s.dbQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      userName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User registered successfulled. You are now logged in as %s.\n", userName)
	log.Printf("user registered: %s\n", userName)
	return nil

}
