package main

import (
	"context"
	"fmt"
	"time"

	"github.com/EBregains/notice-it/internal/database"
	"github.com/google/uuid"
)

func Login(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects a single argument: login <username>")
	}
	username := cmd.arguments[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Welcome %s! Your user has been set.", username)
	return nil
}

func Register(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects a single argument: register <username>")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("User %s succesfuly created.", s.cfg.CurrentUserName)
	fmt.Println(user)
	return nil
}

func Users(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}

func Agg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects only one argument: agg <time-between-reqs>")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %vm\n", timeBetweenReqs.Minutes())
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		ScrapeFeeds(s)
	}
}

func AddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("expects two argument: addfeed <name> <url>")
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed succesfuly adedd: %v", feed)
	return nil
}

func Feeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("expects no arguments: feeds")
	}
	feedsWithUsernames, err := s.db.GetFeedsWithUsernames(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feed\tUrl\tUsername")
	for _, feed := range feedsWithUsernames {
		fmt.Printf("%s\t%s\t%s\n", feed.Name, feed.Url, feed.UserName.String)
	}
	return nil
}

func Follow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects a single argument: follow <url>")
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	followedFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("User %s successfuly followed %s", followedFeed.UserName, followedFeed.FeedName)
	return nil
}

func Following(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("expects no arguments: following")
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}

func Unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects one argument: unfollow <feed-url>")
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Unfollowed %s", feed.Name)
	return nil
}

// -------------
//
//	DANGER ZONE
//
// -------------
func RESET_USERS(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Users table has been reset succesfuly.")
	return nil
}
