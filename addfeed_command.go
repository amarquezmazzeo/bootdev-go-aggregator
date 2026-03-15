package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed handler expects two arguments (name, url)")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	params := database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	}

	feed, err := s.dbQueries.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	followParams := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.dbQueries.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Println(feed)

	return nil
}
