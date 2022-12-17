package main

import (
	"heroes/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8080")
}
