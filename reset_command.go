package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("reset handler expects no arguments")
	}
	err := s.dbQueries.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting user database: %w", err)
	}
	fmt.Println("User database has been reset")
	return nil
}
