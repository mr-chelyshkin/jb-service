package main

import (
	"context"
	"flag"

	"github.com/mr-chelyshkin/jb-service/app"
	_ "github.com/mr-chelyshkin/jb-service/app/docs"
	serviceApp "github.com/mr-chelyshkin/jb-service/app/internal/app"
)

func init() {
	flag.StringVar(&app.LogFile, "log-file", "", "log filepath.")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	a, err := serviceApp.New()
	if err != nil {
		panic(err)
	}
	a.Run(ctx)
}
