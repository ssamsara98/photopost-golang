package main

import (
	"go-photopost/src"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/app"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// //go:generate go run github.com/google/wire/cmd/wire

func main() {
	logger := log.Default()

	err := godotenv.Load()
	if err != nil {
		logger.Panicln(err.Error())
	}

	// // server := InitServer()
	// // server.Start()

	fx.New(
		src.ServerModuleFx,
		lib.LibModuleFx,
		middlewares.MiddlewaresModuleFx,
		app.AppModuleFx,
		users.UsersModuleFx,
		posts.PostsModulefx,
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
