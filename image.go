package termloader

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

// Braille config
type Braille [2][4]int

// Rune maps each point in braille to a dot identifier and calculates the corresponding unicode symbol.
func (b Braille) Rune() rune {
	lowEndian := [8]int{b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2], b[0][3], b[1][3]}
	var v int
	for i, x := range lowEndian {
		v += x << uint(i)
	}
	return rune(v) + '\u2800'
}

// String returns a unicode braille character.
func (b Braille) String() string {
	return string(b.Rune())
}

// Image config
type Image struct {
	Path    string    // Image path
	Filters *Filters  // Image Filters
	Writer  io.Writer // Stdout
	width   int       // image width
	height  int       // image height
}

// Image filters
type Filters struct {
	Sharpness float64 // Sharpness
}

// SetPath sets the provided image path to the image config.
func (i *Image) SetPath(path string) {
	i.Path = path
}

// SetWidth sets the given custom image width to the image config.
func (i *Image) SetWidth(width int) {
	i.width = width
}

// SetHeight sets the given custom image height to the image config.
func (i *Image) SetHeight(height int) {
	i.height = height
}

// Sharpen sharpens the image by factor of given sigma value. Sigma must be a positive value.
func (i *Image) Sharpen(sigma float64) {
	i.Filters.Sharpness = sigma
}

// Renders the image at the center of the stdout depending the provided terminal width and height of the terminal. The
//offset can be used to adjust the vertical positioning of the image. It'll panic if any error occurs while rendering
//the image.
func (i *Image) Render(tWidth, tHeight, offset int) {
	reader, err := os.Open(i.Path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = reader.Close()
	}()

	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	if i.width == 0 && i.height == 0 {
		_, _ = reader.Seek(0, 0)
		config, _, err := image.DecodeConfig(reader)
		if err != nil {
			panic(err)
		}
		i.width, i.height = config.Width, config.Height
	}

	resizedImage := imaging.Resize(img, i.width, i.height, imaging.Lanczos)
	bounds := resizedImage.Bounds()

	if i.Filters.Sharpness > 0 {
		resizedImage = imaging.Sharpen(resizedImage, i.Filters.Sharpness)
	}

	dst := image.NewPaletted(bounds, []color.Color{color.Black, color.White, color.Transparent})
	draw.FloydSteinberg.Draw(dst, bounds, resizedImage, bounds.Min)

	_, _ = fmt.Fprint(i.Writer, i.vCenter(tHeight, offset))
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 4 {
		_, _ = fmt.Fprint(i.Writer, i.hCenter(tWidth))
		for x := bounds.Min.X; x < bounds.Max.X; x += 2 {
			var braille Braille
			for py := 0; py < 4; py++ {
				for px := 0; px < 2; px++ {
					if px+x >= bounds.Max.X || py+y >= bounds.Max.Y {
						continue
					}
					if dst.At(px+x, py+y) == color.Black {
						braille[px][py] = 1
					}
				}
			}
			if _, err := i.Writer.Write([]byte(braille.String())); err != nil {
				panic(err)
			}
		}
		if _, err := i.Writer.Write([]byte{'\n'}); err != nil {
			panic(err)
		}
	}
}

// helper function to render the image vertically centered on the terminal screen.
func (i *Image) hCenter(tWidth int) string {
	var spaces string
	imgWidth := i.width / 2 // braille width: 2
	center := (tWidth / 2) - (imgWidth / 2)
	for i := 0; i < center; i++ {
		spaces += " "
	}
	return spaces
}

// helper function to render the image horizontally centered on the terminal screen.
func (i *Image) vCenter(tHeight, offset int) string {
	var lines string
	imgHeight := (i.height / 4) + offset // braille height: 4
	center := (tHeight / 2) - (imgHeight / 2)
	for i := 0; i < center; i++ {
		lines += "\n"
	}
	return lines
}
