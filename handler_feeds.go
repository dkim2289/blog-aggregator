package main

import (
	"context"
	"fmt"

	"github.com/dkim2289/blog-aggregator/internal/database"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		printFeedWithUser(feed)
		fmt.Println()
		fmt.Println("====================================")
	}

	return nil
}

func printFeedWithUser(feed database.GetFeedsWithUserRow) {
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("User ID: %s\n", feed.UserID)
	fmt.Printf("User Name: %s\n", feed.UserName)
}
