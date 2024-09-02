package main

import (
	_ "dddapib/docs" // swag generated files
	"dddapib/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
