package main

import (
	"dddapib/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
