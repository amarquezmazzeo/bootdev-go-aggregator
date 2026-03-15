package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/amarquezmazzeo/bootdev-go-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return errors.New("browse handler expects at most 1 argument (limit)")
	}
	limit := 2
	if len(cmd.args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("error parsing 'limit' argument to integer: %w", err)
		}
		if limit == 0 {
			return errors.New("limit argument must be > 0")
		}
	}
	limit32 := int32(limit)

	params := database.GetUserPostsParams{
		ID:    user.ID,
		Limit: limit32,
	}
	posts, err := s.dbQueries.GetUserPosts(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\n"+
			"URL: %s\n"+
			"Published Date: %d-%d-%d\n"+
			"Description: %s\n"+
			"Feed Name: %s\n",
			post.Title,
			post.Url,
			post.PublishedAt.Time.Year(),
			post.PublishedAt.Time.Month(),
			post.PublishedAt.Time.Day(),
			post.Description,
			post.FeedName)
	}
	return nil
}
