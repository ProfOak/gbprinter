package main

import (
	"image/color"
	"sort"
)

// Palette contains four colors because the GameBoy only had four colors.
// This was extended later on to allow each sprite to have their own set of
// four colors, in the GameBoy Color.
type Palette struct {
	Black     color.RGBA
	DarkGray  color.RGBA
	LightGray color.RGBA
	White     color.RGBA
}

// ColorPalettes are a collection of the classic GB and GBC palettes.
type ColorPalettes = map[string]*Palette

var palettes = ColorPalettes{
	"grayscale": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{63, 63, 63, 255},
		LightGray: color.RGBA{157, 157, 157, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "greenscale": &Palette{
		Black:     color.RGBA{15, 56, 15, 255},
		DarkGray:  color.RGBA{48, 98, 48, 255},
		LightGray: color.RGBA{139, 172, 15, 255},
		White:     color.RGBA{155, 188, 15, 255},
	}, "up": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{131, 49, 0, 255},
		LightGray: color.RGBA{255, 173, 99, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "upa": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{148, 58, 58, 255},
		LightGray: color.RGBA{255, 133, 132, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "upb": &Palette{
		Black:     color.RGBA{91, 49, 9, 255},
		DarkGray:  color.RGBA{132, 107, 41, 255},
		LightGray: color.RGBA{206, 156, 133, 255},
		White:     color.RGBA{255, 231, 197, 255},
	}, "left": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{0, 0, 254, 255},
		LightGray: color.RGBA{101, 164, 155, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "lefta": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{83, 82, 140, 255},
		LightGray: color.RGBA{139, 140, 222, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "leftb": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{82, 82, 82, 255},
		LightGray: color.RGBA{165, 165, 165, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "down": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{147, 148, 254, 255},
		LightGray: color.RGBA{254, 148, 148, 255},
		White:     color.RGBA{255, 255, 165, 255},
	}, "downa": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{254, 0, 0, 255},
		LightGray: color.RGBA{255, 255, 0, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "downb": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{125, 73, 0, 255},
		LightGray: color.RGBA{255, 255, 0, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "right": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{255, 66, 0, 255},
		LightGray: color.RGBA{81, 255, 0, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "righta": &Palette{
		Black:     color.RGBA{0, 0, 0, 255},
		DarkGray:  color.RGBA{1, 99, 198, 255},
		LightGray: color.RGBA{123, 255, 48, 255},
		White:     color.RGBA{255, 255, 255, 255},
	}, "rightb": &Palette{
		Black:     color.RGBA{255, 255, 255, 255},
		DarkGray:  color.RGBA{255, 255, 0, 255},
		LightGray: color.RGBA{0, 132, 134, 255},
		White:     color.RGBA{0, 0, 0, 255},
	},
}

// paletteNames is an easy way to generate the key names of the palettes map.
func paletteNames() []string {

	var names []string
	for name := range palettes {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
