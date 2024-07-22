package main

import (
	"github.com/ssamsara98/photopost-golang/src/bootstrap"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
