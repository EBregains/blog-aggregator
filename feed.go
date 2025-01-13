package main

import (
	"context"
	"fmt"
)

func ScrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Items in %s:\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Printf(" * %s\n", item.Title)
	}
	return nil
}
