package main

import (
	"flag"
	"github.com/3boku/mongo-study/config"
	"github.com/3boku/mongo-study/main/app"
)

func main() {
	flag.Parse()
	cfg := config.NewConfig("./config.toml")
	app.NewApp(cfg)
}
