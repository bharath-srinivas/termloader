package main

import (
	"time"

	"github.com/bharath-srinivas/termloader"
)

func main() {
	loader := termloader.New(termloader.CharsetConfigs["default"]) // construct a new loader with config
	loader.Color = termloader.Green                                // provide a color for the loader
	loader.Text = "Loading"                                        // provide a text to show above loader

	loader.Start()              // start the loader
	time.Sleep(5 * time.Second) // sleep for sometime to simulate a task
	loader.Stop()               // stop the loader

	time.Sleep(2 * time.Second) // sleep for sometime before performing next set of operations

	loadingText := termloader.ColorString("Now Loading", termloader.Green) // color the string
	loader.Text = loadingText                                              // provide the colored string as loading text

	loader.Start()              // start the loader
	time.Sleep(5 * time.Second) // sleep for sometime to simulate a task
	loader.Stop()               // stop the loader

	time.Sleep(2 * time.Second) // sleep again before performing next set of operations

	loader.Text = ""                              // remove the loading text
	loader.Image.SetPath("../assets/loading.png") // provide the path of the loading image
	loader.Image.SetWidth(55)                     // set custom width for the image
	loader.Image.SetHeight(15)                    // set custom height for the image
	loader.Image.Sharpen(6.5)                     // sharpens the image

	loader.Start()              // start the loader
	time.Sleep(5 * time.Second) // sleep for sometime to simulate a task
	loader.Stop()               // stop the loader
}
