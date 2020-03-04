package termloader

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	loader := New(CharsetConfigs["default"])
	want := "*termloader.Loader"
	got := reflect.TypeOf(loader).String()
	if got != want {
		t.Errorf("New returned incorrect type, got: %s, want: %s", got, want)
	}
}

func TestLoader_Start(t *testing.T) {
	loader := New(CharsetConfigs["default"])
	loader.Color = Green
	loader.Text = "Testing"
	loader.Start()
	if !loader.active {
		t.Error("Start did not start the Loader")
	}
	time.Sleep(2 * time.Second)
	loader.Stop()
}

func TestLoader_Stop(t *testing.T) {
	loader := New(CharsetConfigs["default"])
	loader.Color = Green
	loader.Text = "Testing"
	loader.Start()
	time.Sleep(2 * time.Second)
	loader.Stop()
	if loader.active {
		t.Error("Stop did not stop the Loader")
	}
}

func TestColorString(t *testing.T) {
	input := "Testing"
	want := fmt.Sprintf("\033[32m%s\033[0m", input)
	if got := ColorString(input, Green); got != want {
		t.Errorf("ColorString returned incorrect output, got: %s, want: %s", got, want)
	}
}
