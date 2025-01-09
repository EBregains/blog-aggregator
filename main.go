package main

import (
	"fmt"

	"github.com/EBregains/notice-it/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	cfg.SetUser("Emiliano")
	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
