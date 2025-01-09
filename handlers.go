package main

import "fmt"

func Login(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("expects a single argument: login <username>")
	}
	username := cmd.arguments[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Welcome %s! Your user has been set.", username)
	return nil
}
