package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed handler expects two arguments (name, url)")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	currentUser := s.cfg.CurrentUserName

	user, err := s.dbQueries.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	params := database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	}

	feed, err := s.dbQueries.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	fmt.Println(feed)

	return nil
}
