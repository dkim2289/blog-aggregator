package main

import (
	"context"
	"fmt"

	"github.com/dkim2289/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("unfollowed feed: %s\n", feed.Name)
	return nil
}
