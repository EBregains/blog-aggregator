package main

import (
	"github.com/EBregains/notice-it/internal/config"
	"github.com/EBregains/notice-it/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
