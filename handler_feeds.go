package main

import (
	"context"
	"fmt"

	"github.com/gdbeltran/gator/internal/database"
)

func handlerPrintFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		printFeeds(feed)
		fmt.Println("======================================")
	}

	return nil
}

func printFeeds(feed database.GetFeedsRow) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", feed.Username)
}
