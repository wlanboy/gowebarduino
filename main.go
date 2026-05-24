package main

import (
	app "github.com/wlanboy/gowebarduino/application"
)

func main() {
	a := app.GoService{}
	a.Initialize()

	a.Run()

	a.WaitForShutdown()
}
