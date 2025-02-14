package main

import (
	"github.com/aws-cakap-intern/book-store/config"
	"github.com/aws-cakap-intern/book-store/internal/builder"
	"github.com/aws-cakap-intern/book-store/pkg/db"
	"github.com/aws-cakap-intern/book-store/pkg/server"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	database, err := db.InitDB(&cfg.Database)
	checkError(err)


	publicRoutes := builder.BuildAppPublicRoutes(database)


	srv := server.NewServer(publicRoutes)
	srv.Run(cfg.Port)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}