package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// This file is almost 100%from snake-life repo
var screen = bufio.NewWriter(os.Stdout)

func hideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func showCursor() {
	fmt.Fprint(screen, "\033[?25h")
}

func moveCursor(v vector) {
	fmt.Fprintf(screen, "\033[%d;%dH", v.y, v.x)
}

func green() {
	fmt.Fprint(screen, "\033[32m")
}

func resetColor() {
	fmt.Fprint(screen, "\033[0m")
}

func clear() {
	fmt.Fprint(screen, "\033[2J")
}

func draw(str string) {
	fmt.Fprint(screen, str)
}

func render() {
	screen.Flush()
}

func getSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return width, height
}
