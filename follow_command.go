package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("follow handler expects a single argument (url)")
	}
	url := cmd.args[0]

	currentUser := s.cfg.CurrentUserName

	userID, err := s.dbQueries.GetUserID(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting current user id: %w", err)
	}
	feedID, err := s.dbQueries.GetFeedID(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed id for url provided: %w", err)
	}

	params := database.CreateFeedFollowParams{
		UserID: userID,
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
