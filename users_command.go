package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("user handler no additional arguments")
	}

	users, err := s.dbQueries.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}
