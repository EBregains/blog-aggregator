package main

import (
	"fmt"
	"os"

	"github.com/EBregains/notice-it/internal/config"
)

func main() {
	// Arguments
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Missing command or argument.")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	appState := state{
		cfg: &cfg,
	}
	cmds := commands{
		list: make(map[string]func(*state, command) error),
	}
	cmds.register("login", Login)
	err = cmds.run(&appState, command{
		name:      args[1],
		arguments: args[2:],
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(appState.cfg)
}
