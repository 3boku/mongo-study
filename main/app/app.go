package app

import "github.com/3boku/mongo-study/config"

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) *App {
	a := &App{}

	return a
}
