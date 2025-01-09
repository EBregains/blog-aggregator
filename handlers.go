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
