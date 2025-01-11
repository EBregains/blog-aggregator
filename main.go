package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/EBregains/notice-it/internal/config"
	"github.com/EBregains/notice-it/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Arguments
	args := os.Args

	// Generate Config
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	// Open and create Database
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	// Initialize state
	appState := state{
		cfg: &cfg,
		db:  dbQueries,
	}
	// register Commands
	cmds := commands{
		list: make(map[string]func(*state, command) error),
	}

	cmds.register("login", Login)
	cmds.register("register", Register)
	cmds.register("users", Users)
	cmds.register("agg", Agg)
	cmds.register("addfeed", middlewareLoggedIn(AddFeed))
	cmds.register("feeds", Feeds)
	cmds.register("follow", middlewareLoggedIn(Follow))
	cmds.register("following", middlewareLoggedIn(Following))
	cmds.register("unfollow", middlewareLoggedIn(Unfollow))
	cmds.register("reset", RESET_USERS) // DANGER ZONE!

	err = cmds.run(&appState, command{
		name:      args[1],
		arguments: args[2:],
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
