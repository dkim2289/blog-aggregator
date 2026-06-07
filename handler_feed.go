package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dkim2289/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID: %s\n", feed.ID)
	fmt.Printf("* CreatedAt: %s\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt: %s\n", feed.UpdatedAt)
	fmt.Printf("* Name: %s\n", feed.Name)
	fmt.Printf("* URL: %s\n", feed.Url)
	fmt.Printf("* UserID: %s\n", feed.UserID)
}
