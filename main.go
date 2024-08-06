package main

import (
	"github.com/joho/godotenv"
	"github.com/ssamsara98/photopost-golang/src/bootstrap"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
