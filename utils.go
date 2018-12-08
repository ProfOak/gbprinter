package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

// saveImage reduces the boilerplate code needed to save an image.
func saveImage(filename string, img image.Image) {

	f, err := os.Create(filename)
	check(err)
	if err != nil {
		log.Fatalf("Error creating file: %s\n", filename)
	}
	defer f.Close()

	err = png.Encode(f, img)
	check(err)
}

// openImage reduces the boilerplate code needed to open an image.
func openImage(filename string) image.Image {

	inputFile, err := os.Open(filename)
	check(err)

	img, _, err := image.Decode(inputFile)
	check(err)
	return img
}

// check reduces the boilerplate code for checking error values.
func check(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

// min is a helper function to calculate the min of two uint8s. The min function
// in the standard library only operates on float64.
func min(a, b uint8) uint8 {

	if a < b {
		return a
	}
	return b
}
