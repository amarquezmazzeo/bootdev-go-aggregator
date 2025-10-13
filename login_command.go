package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login handler expects a single argument (username)")
	}
	user := cmd.args[0]

	err := s.cfg.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", user)
	return nil
}
