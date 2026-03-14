package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("following handler expects no additional arguments")
	}

	following, err := s.dbQueries.GetFeedFollowsForUser(
		context.Background(),
		s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user feed follow list: %w", err)
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
