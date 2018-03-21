package main

import (
	"time"

	"github.com/bharath-srinivas/termloader"
)

func main() {
	loader := termloader.New(termloader.Charsets[0], 100*time.Millisecond) // construct a new loader with config
	loader.Color = termloader.Green                                        // provide a color for the loader
	loader.Text = "Loading"                                                // provide a text to show above loader

	loader.Start()              // start the loader
	time.Sleep(5 * time.Second) // sleep for sometime to simulate a task
	loader.Stop()               // stop the loader

	time.Sleep(2 * time.Second) // sleep for sometime before performing next set of operations

	loadingText := termloader.ColorString("Now Loading", termloader.Green) // color the string
	loader.Text = loadingText                                              // provide the colored string as loading text

	loader.Start()              // start the loader
	time.Sleep(5 * time.Second) // sleep for sometime to simulate a task
	loader.Stop()               // stop the loader
}
