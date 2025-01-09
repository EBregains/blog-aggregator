package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (cmds *commands) register(name string, f func(*state, command) error) error {
	_, exists := cmds.list[name]
	if exists {
		return fmt.Errorf("command \"%s\" already exits", name)
	}
	cmds.list[name] = f
	return nil
}

func (cmds *commands) run(s *state, cmd command) error {
	fun, exists := cmds.list[cmd.name]
	if !exists {
		return fmt.Errorf("error: command \"%s\" is not registered", cmd.name)
	}
	err := fun(s, cmd)
	return err
}
