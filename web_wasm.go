package main

import (
	"fmt"
	"image/color"
)

// Drawer implementation
type CTX struct{}

// implemented in JS
func _clear_screen(w, h int, c string, c_len int)
func _fill_rect(x, y, w, h int, c string, c_len int)

// these are our Drawer interface functions that dispatch to the js functions
func (CTX) clear_screen(c color.RGBA) {
	// passing the string will just pass a pointer to JS from wasm so we also need to pass the length so we can get the string from wasm memory on the JS side
	col := fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
	_clear_screen(width, height, col, len(col))
}

func (CTX) fill_rect(x, y, w, h int, c color.RGBA) {
	col := fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
	_fill_rect(x, y, w, h, col, len(col))
}

func main() {
	// create drawer interface
	drawer = CTX{}
}
