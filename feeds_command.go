package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("feeds handler expects no additional arguments")
	}

	feeds, err := s.dbQueries.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching all feeds: %w", err)
	}

	fmt.Println(feeds)

	return nil
}
