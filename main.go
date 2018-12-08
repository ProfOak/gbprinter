package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"os"
	"strings"
)

var (
	// Split 0-255 into 4 parts to come up with the GameBoy colors.
	black     uint8
	darkGray  uint8 = 84
	lightGray uint8 = 169
	white     uint8 = 255

	// I couldn't decide whether or not to use color values or in-between color
	// values for determining what new color to translate the original pixel's
	// color value into. Some images looked better using the hard boundries
	// of the four color values above. Some images looked better using the three
	// in between values below. Hold onto the in-between values in case I revisit
	// this in the future.

	// Find the average between the neighboring colors and use them as
	// color boundries that the image will round down to.
	low    = (black + darkGray) / 2
	medium = (darkGray + lightGray) / 2
	high   = (lightGray + white) / 2

	// This is the original GameBoy's resolution.
	height = 144
	width  = 160
)

func main() {

	var palette *Palette

	choiceMessage := "Color choices: " + strings.Join(paletteNames(), ", ")
	paletteChoice := flag.String("palette", "grayscale", choiceMessage)
	flag.Parse()

	if p, ok := palettes[*paletteChoice]; ok {
		palette = p
	} else {
		fmt.Printf("Invalid color palette: %s\n", *paletteChoice)
		fmt.Println(choiceMessage)
		os.Exit(1)
	}

	for _, filename := range flag.Args() {

		// Open the image.
		sourceImage := openImage(filename)

		// TODO: Figure out how to get better colors.
		// Calculate the average brightness using the relative luminance formula.
		//brightness := averageBrightness(sourceImage)

		// Configure the global variables in main to reflect the average brightness.
		//configureColors(brightness)

		// Resize and recolor the image based on the color palette.
		output := gbPrinter(sourceImage, palette)

		// format: sourceName_paletteName.png
		filename = fmt.Sprintf(
			"%s_%s.png", strings.Split(filename, ".")[0], *paletteChoice)
		saveImage(filename, output)
	}
}

// gbPrinter takes a source image and produces a new image based on the palette.
func gbPrinter(sourceImage image.Image, palette *Palette) image.Image {

	sourceImage = resize.Resize(
		uint(width),
		uint(height),
		sourceImage,
		resize.NearestNeighbor,
	)

	screen := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < sourceImage.Bounds().Max.X; x++ {
		for y := 0; y < sourceImage.Bounds().Max.Y; y++ {
			screen.Set(x, y, transformColor(sourceImage.At(x, y), palette))
		}
	}

	return screen
}

// configureColors will change the global variables based on the brightness
// of the current image.
func configureColors(brightness uint8) {

	var ceiling float32

	// Use this as an upper limit of how high the white color can go.
	// Also try to prevent overflows.
	brightness = min(brightness, 127)
	brightness *= 2

	ceiling = float32(brightness)

	white = 255
	lightGray = uint8(ceiling)
	darkGray = uint8(ceiling * (2.0 / 3.0))
	black = 0

	low = (black + darkGray) / 2
	medium = (darkGray + lightGray) / 2
	high = (lightGray + white) / 2
}

// averageBrightness is used to configure the w/b color values.
func averageBrightness(img image.Image) uint8 {

	var sum int
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			sum += int(relativeLuminance(img.At(x, y)))
		}
	}

	return uint8(sum / (height * width))
}

// relativeLuminance calculates the relative luminance from a pixel of an image
// and effectively returns the brightness of a pixel.
// source: https://en.wikipedia.org/wiki/Relative_luminance
func relativeLuminance(currentColor color.Color) uint8 {

	// Color.RGBA returns uint32 but color.RBGA uses uint8
	tempr, tempg, tempb, _ := currentColor.RGBA()

	// RGBA() on a uint8 color is bit shifted left by 8. Undo this.
	r := uint8(tempr >> 8)
	g := uint8(tempg >> 8)
	b := uint8(tempb >> 8)

	// This number is an approximation.
	return uint8((0.2126 * float32(r)) +
		(0.715 * float32(g)) +
		(0.0722 * float32(b)))
}

// transformColor takes a pixel from an image and calculates the brightness.
// The brightness value is then rounded down to one of the 4 gameboy colors.
// The given palette is used to replace the color with another color scheme
// from the GameBoy color.
func transformColor(oldColor color.Color, palette *Palette) color.RGBA {

	var newColor color.RGBA

	// This number is an approximation.
	relative := relativeLuminance(oldColor)

	// Round the relative value to the nearest GameBoy color pallet value.
	if relative < darkGray {
		newColor = palette.Black
	} else if relative < lightGray {
		newColor = palette.DarkGray
	} else if relative < white {
		newColor = palette.LightGray
	} else {
		newColor = palette.White
	}

	return newColor
}
