package main

import (
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type vector struct {
	x int
	y int
}

type tile struct {
	pos vector
	char string
}

type game struct {
	score  int
	maxPos vector
	frame int
	rectangles []rectangle
}


type rectangle struct {
	id string
	char string
	fill string
	pos vector
	size vector
}

// Collection of methods the snake-life repo used for...
// drawing on the terminal. Me having to call them like that...
// probably means that I'm doing something wrong, but I'm...
// too lazy to actually think of this now, so I made this wrapper.
func p(t tile){
	moveCursor(t.pos)
	draw(t.char)
	//render()
}

// Used to generate a rectangle object with... 
// random pos and random size
func newRectangle(char string, fill string) *rectangle{
	var pos = randomPosition()
	// TODO: Check for wall collisions
	var v = vector {
		x: rand.Intn(20) + 1,
		y: rand.Intn(12) + 2,

	}
	// Look trough rectangles to avoid collision
	// game.rectangles 

	return &rectangle{
		id: strconv.Itoa(rand.Intn(20) * rand.Intn(10)),
		pos: pos,
		size: v,
		char: char,
		fill: fill,
	}
}


func main() {
	game := newGame()
	game.beforeGame()
	
	// "Game loop"
	for {
		maxX, maxY := getSize()
		game.maxPos.x = maxX
		game.maxPos.y = maxY

		// This clears the terminal window
		// clear()
		
		var r = newRectangle("#", strconv.Itoa(game.frame))
		placeRectangle(r)
		game.rectangles = append(game.rectangles, *r)
		// I use frames only as filler texture...
		// for rectangles
		game.frame++
	}
}
// From snake-life repo
func newGame() *game {
	rand.Seed(time.Now().UnixNano())
	game := &game{
		frame: 0,
	}

	return game
}
// From snake-life repo
func randomPosition() vector {
	width, height := getSize()
	x := rand.Intn(width) + 4
	y := rand.Intn(height) + 6

	return vector{x: x, y: y}
}


func makeTile (char string, pos vector) tile {
	return tile { 
		char: char,
		pos: pos,
	}
}

// I made this wrapper function, because for some reason...
// I couldn't make the Sleep(arguments) a variable.
func wait(){
	time.Sleep(time.Millisecond * 1)
}

// This wack ass function shouldn't have 4x for loops that draw directly to the screen...
// instead it should create an array of tiles, that a rendering function walks trough...
// at the renering step. 
func makeRectangle(pos vector, width int, height int, char string){
	// Each loop is used to draw one of the outlines in the rectangle.

	// We do some dumb maths to determine the outlines starting...
	// and ending positions, from rectangle-objects width, height and position vectors.
	var buildingXwallStart = pos.x - (width / 2)
	var buildingYwallStart = pos.y - (height / 2)
	a := 0
	for a < width {
		var v  = vector {
			x: buildingXwallStart + a, 
			y: buildingYwallStart,
		}
		p(makeTile(char, v))
		wait()
		a++
	}
	b := 0
	for b < height {	
		var v  = vector {	
			x: buildingXwallStart + width, 
			y: buildingYwallStart + b,
		}
		p(makeTile(char, v))
		wait()
		b++
	}
	render()
	c := 0
	initX := buildingXwallStart + width
	for c < width {	
		var v  = vector {
			x: initX - c, 
			y: buildingYwallStart + height,
		}
		p(makeTile(char, v))
		wait()
		c++
	}
	d := 0
	initY := buildingYwallStart + height
	for d < height {
		var v  = vector {
			x: buildingXwallStart, 
			y: initY - d,
		}
		p(makeTile(char, v))
		wait()
		d++
	}
}

// Method that takes a rectangle-object, and does some stupid maths.
// As the result the rectangles interior will be filled with the...
// fillRectangle value of the rectangle-object.
func fillRectangle(r *rectangle) {
	y := r.size.y - 2
	x := r.size.x - 2
	var run = true
	for run {
		makeRectangle(r.pos, x, y, r.fill)
		if x < 2 && y < 2 {
			run = false
		}
		if(x >= 2){
			x--
		}
		if(y >= 2){
			y--
		}
		
	}
}
func placeRectangle(r *rectangle) {
	makeRectangle(r.pos, r.size.x, r.size.y, r.char)
	fillRectangle(r)
}

// From snake-life repo
func (g *game) beforeGame() {
	hideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			g.over()
		}
	}()
}


// From snake-life repo
func (g *game) over() {
	clear()
	showCursor()

	moveCursor(vector{x:1, y:1})
	draw("rect len " + strconv.Itoa(len(g.rectangles)) + " a.")

	render()

	os.Exit(0)
}

