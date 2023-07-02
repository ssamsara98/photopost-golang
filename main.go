package main

import (
	"photopost/bootstrap"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	_ = bootstrap.RootApp.Execute()
}
