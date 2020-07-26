package main

import (
	"producer/test/config"
	"producer/test/pkg"
)

func main() {
	srv := config.New()
	for{
		pkg.Produce()
	}
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
