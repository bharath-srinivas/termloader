package termloader

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// constants for ANSI escape sequences.
const (
	escape = "\x1b"
	reset  = 0
)

// map for storing the clear function.
var clear map[string]func()

// valid colors.
var validColors = map[int]bool{
	Red:     true,
	Green:   true,
	Yellow:  true,
	Blue:    true,
	Magenta: true,
	Cyan:    true,
	Gray:    true,
}

// Loader config.
type Loader struct {
	Image    *Image        // Loading image
	Color    int           // Color of the loader
	Delay    time.Duration // Animation speed of the loader
	Text     string        // Text to be displayed above the loader
	Writer   io.Writer     // Stdout
	active   bool          // current state of the loader
	charset  []string      // character set for the loader
	mutex    sync.Mutex    // mutex
	stop     chan bool     // channel for stopping the loader
	hasImage bool          // loading image provided
}

// this ugly hack needs to be removed. Need to find a better way to clear the stdout.
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["darwin"] = clear["linux"]
}

// New returns a pointer to the Loader interface with provided options. Default loader color will be white.
func New(charsetConfig CharsetConfig) *Loader {
	charset := charsetConfig.Charset
	delay := charsetConfig.Delay
	return &Loader{
		Image: &Image{
			Filters: &Filters{},
			Writer:  os.Stdout,
		},
		Delay:   delay,
		Color:   None,
		Writer:  os.Stdout,
		active:  false,
		charset: charset,
		mutex:   sync.Mutex{},
		stop:    make(chan bool, 1),
	}
}

// clrScr will clear the stdout.
func clrScr() {
	if value, ok := clear[runtime.GOOS]; ok {
		value()
	} else {
		panic("your platform is unsupported! couldn't clear terminal screen :(")
	}
}

// hCenter is a helper function which will help in rendering the loader horizontally centered on the terminal screen.
func hCenter(s string, width int) string {
	var spaces string
	regex := regexp.MustCompile("\x1b\\[([0-9]{1,2}(;[0-9]{1,2})?)?[m|K]")
	strLen := strings.Count(regex.ReplaceAllString(s, ""), "") - 1
	halfStrLen := strLen / 2
	center := (width / 2) - halfStrLen

	for i := 0; i < center; i++ {
		spaces += " "
	}
	return spaces
}

// vCenter is a helper function which will help in rendering the loader vertically centered on the terminal screen.
func vCenter(lineCount int, height int) string {
	var lines string
	center := (height / 2) - (lineCount / 2)

	for i := 0; i < center; i++ {
		lines += "\n"
	}
	return lines
}

// validColor will return true if the provided color is valid else returns false.
func validColor(color int) bool {
	valid := false
	if validColors[color] {
		valid = true
	}
	return valid
}

// ColorString will wrap the provided string with ANSI escape sequences for color and return the colored string.
func ColorString(str string, color int) string {
	if color == None || !validColor(color) {
		return str
	}

	prefix := fmt.Sprintf("%s[%dm", escape, color)
	suffix := fmt.Sprintf("%s[%dm", escape, reset)
	return prefix + str + suffix
}

// Starts the loader.
func (l *Loader) Start() {
	if l.active {
		return
	}

	l.active = true
	fd := int(os.Stdin.Fd())
	if ok := terminal.IsTerminal(fd); !ok {
		return
	}

	rendered := false
	renderedImage := false
	go func() {
		lineCount := 1
		width, height, _ := terminal.GetSize(fd)
		if l.Text != "" {
			lineCount += 1
		}

		if l.Image.Path != "" {
			l.hasImage = true
		}

		if !l.hasImage {
			_, _ = fmt.Fprint(l.Writer, vCenter(lineCount, height))
		}
		for {
			for i := 0; i < len(l.charset); i++ {
				select {
				case <-l.stop:
					return
				default:
					l.mutex.Lock()
					if l.hasImage && !renderedImage {
						_, _ = fmt.Fprintf(l.Writer, "%s[?25l", escape) // disable cursor
						l.Image.Render(width, height, lineCount)
						renderedImage = true
					}
					if l.Text != "" && !rendered {
						textCenter := hCenter(l.Text, width)
						_, _ = fmt.Fprintf(l.Writer, "%s[?25l", escape) // disable cursor
						_, _ = fmt.Fprintln(l.Writer, textCenter+l.Text)
						rendered = true
					}
					loaderCenter := hCenter(l.charset[i], width)
					_, _ = fmt.Fprintf(l.Writer, "%s\r", loaderCenter+ColorString(l.charset[i], l.Color))
					l.mutex.Unlock()
					time.Sleep(l.Delay)
				}
			}
		}
	}()
}

// Stops the loader.
func (l *Loader) Stop() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.active {
		l.active = false
		fmt.Printf("%s[?25h", escape) // enable cursor
		l.stop <- true
		clrScr()
	}
}
