package main

import (
	"flag"
)

func main() {
	flag.Parse()

	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(app.GetConfig().GetHttp().LocalAddr); err != nil {
		panic(err)
	}

}
