package main

import (
	"image/color"
)

// this defines the drawing api that the platform needs to implement
type Drawer interface {
	clear_screen(c color.RGBA)
	fill_rect(x, y, w, h int, c color.RGBA)
}

const (
	title  = "Golang Move Square"
	width  = 500
	height = 500
	size   = 50
	speed  = 500
	// Keys
	// originally used iota but it gave the wrong values in wasm for some reason
	none  = 0
	esc   = 1
	left  = 2
	up    = 3
	right = 4
	down  = 5
)

var (
	background = color.RGBA{0, 0, 0, 255}       // black background
	col        = color.RGBA{255, 255, 255, 255} // white color for square
	running    = true
	x, y       = 0, 0
	x_dir      = 0.0
	y_dir      = 0.0
	m_left     = false
	m_up       = false
	m_right    = false
	m_down     = false
	// The platform layers, minifb and web_wasm provide the implementations for drawer
	// we use this to draw
	drawer Drawer
)

//export input
func input(key int, pressed bool) {
	if key == esc {
		running = false
		return
	}

	move := false

	if pressed {
		move = true
	}
	if key == left {
		m_left = move
	}
	if key == up {
		m_up = move
	}
	if key == right {
		m_right = move
	}
	if key == down {
		m_down = move
	}

}

//export update
func update(dt float64) {
	x_dir = 0.0
	y_dir = 0.0

	// update direction
	if m_left {
		x_dir += -1.0
	}
	if m_up {
		y_dir += -1.0
	}
	if m_right {
		x_dir += 1.0
	}
	if m_down {
		y_dir += 1.0
	}

	x += int(x_dir * speed * dt)
	y += int(y_dir * speed * dt)

	// collision with walls
	if x < 0 {
		x = 0
	}

	if x+size > width {
		x = width - size
	}

	if y < 0 {
		y = 0
	}

	if y+size > height {
		y = height - size
	}
}

//export draw
func draw() {
	drawer.clear_screen(background)
	drawer.fill_rect(x, y, size, size, col)
}

// I'm exporting getters for width and height to use in JS
// I could access the memory directly via the wasm object but this is simpler
//export get_width
func get_width() int {
	return width
}

//export get_height
func get_height() int {
	return height
}
