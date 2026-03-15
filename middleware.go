package main

import (
	"context"
	"fmt"

	"github.com/amarquezmazzeo/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error retrieving current user: %w", err)
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}
