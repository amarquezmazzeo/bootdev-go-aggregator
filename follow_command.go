package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("follow handler expects a single argument (url)")
	}
	url := cmd.args[0]

	feedID, err := s.dbQueries.GetFeedID(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed id for url provided: %w", err)
	}

	params := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feedID,
	}
	feedFollow, err := s.dbQueries.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Feed: %s, User: %s\n",
		feedFollow.FeedName,
		feedFollow.UserName)

	return nil
}
