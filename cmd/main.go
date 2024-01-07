package main

import (
	"fmt"

	"github.com/koha90/url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: SQLite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
