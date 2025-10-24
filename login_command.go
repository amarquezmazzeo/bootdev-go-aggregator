package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login handler expects a single argument (username)")
	}
	user := cmd.args[0]

	// check if user exists in db
	_, err := s.dbQueries.GetUser(context.Background(), user)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", user)
	return nil
}
