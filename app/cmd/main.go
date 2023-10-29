package main

import (
	"context"

	_ "github.com/mr-chelyshkin/jb-service/app/docs"
	"github.com/mr-chelyshkin/jb-service/app/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.New()
	if err != nil {
		panic(err)
	}
	a.Run(ctx)
}
