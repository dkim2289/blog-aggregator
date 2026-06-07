package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dkim2289/blog-aggregator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = parsedLimit
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\n", post.Title, post.Url)
	}
	return nil
}
