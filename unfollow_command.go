package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("unfollow handler expects 1 argument (url)")
	}
	url := cmd.args[0]

	feedID, err := s.dbQueries.GetFeedID(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed id from url provided: %w", err)
	}

	err = s.dbQueries.RemoveFeedFollow(context.Background(), feedID)
	if err != nil {
		return fmt.Errorf("error removing feed from feed follow list: %w", err)
	}
	return nil
}
