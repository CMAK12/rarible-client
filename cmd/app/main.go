package main

import (
	_ "inforce-task/docs" // Import docs for swagger
	"inforce-task/internal/app"
)

// @title Inforce Task API
// @version 1.0
// @description Rarible Client API for fetching NFT data
// @host localhost:8080
// @BasePath /
func main() {
	if err := app.MustRun(); err != nil {
		panic(err)
	}
}
