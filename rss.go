package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/amarquezmazzeo/gator/internal/database"
	"github.com/lib/pq"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error preparing rss fetch feed request: %w", err)
	}
	request.Header.Set("User-agent", "gator")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error executing rss fetch feed request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading rss fetch feed response body: %w", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling rss fetch feed response body: %w", err)
	}

	err = unescapeFeed(&feed)
	if err != nil {
		return nil, fmt.Errorf("error unscaping rss feed: %w", err)
	}

	return &feed, nil
}

func unescapeFeed(feed *RSSFeed) error {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		item := &feed.Channel.Item[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("agg handler expects 1 arguments (time_between_reqs)")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func scrapeFeeds(s *state) error {
	const uniqueViolation = pq.ErrorCode("23505")

	feed, err := s.dbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}
	err = s.dbQueries.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	for _, item := range fetchedFeed.Channel.Item {
		if len(item.Title) == 0 {
			continue
		}
		fmt.Println(item.Title)
		// Parsing PubDate to sql.NullTime
		pubDate := sql.NullTime{}
		if t, err := parseTime(item.PubDate); err == nil {
			pubDate = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		} else {
			log.Printf("warning: could not parse pubDate (%s): %v", item.PubDate, err)
		}
		postParams := database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}
		err = s.dbQueries.CreatePost(context.Background(), postParams)
		// TODO: handle duplicate post creations gracefully
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
				if pqErr.Code == uniqueViolation {
					// ignore fails due to duplicate post
				} else {
					log.Printf("warning: error creating post in db: %v", err)
				}
			} else {
				return fmt.Errorf("error creating post in db: %w", err)
			}
		}
	}
	return nil
}

func parseTime(s string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		// add more as needed
	}
	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse time: %s", s)
}
