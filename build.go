package main

import (
	"os"
)

func main() {
	target := "desktop"

	if len(os.Args) > 1 {
		target = os.Args[1]
	}

}

func desktop() {

}

func web() {
}

func server() {
}

func minifb() {
}
