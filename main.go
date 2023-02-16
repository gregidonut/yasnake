package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

func main() {

	a := app.New()

	w := a.NewWindow("YASnake")
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	w.Canvas().SetOnTypedKey(keyTyped)

	game = setupGame()
	w.SetContent(game)

	go runGame()

	w.ShowAndRun()
}

type snakePart struct {
	x, y float32
}

type moveType int

const (
	moveUp moveType = iota
	moveDown
	moveLeft
	moveRight
)

var (
	snakeParts []snakePart
	game       *fyne.Container
	move       = moveUp
)

// setupGame() will be the starting position for the snake
// since it defines how everything looks before redrawing the screen
// which would be the consequence of the game loop
func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject

	// draw snake by adding 10 rectangle elements to segments slice
	for i := 0; i < 10; i++ {
		r := canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
		r.Resize(fyne.NewSize(10, 10))

		// use iteration variable to compute where the next rectangle will be
		r.Move(fyne.NewPos(90, float32(50+i*10)))

		segments = append(segments, r)

		seg := snakePart{9, float32(5 + i)}
		snakeParts = append(snakeParts, seg)
	}

	return container.NewWithoutLayout(segments...)
}

// refreshGame() is how the ui redraws the screen
// which is by multiplying the units in snakePart's coordinate value
// by 10, which will be the new position of the objects(snake segment)
func refreshGame() {
	for i, seg := range snakeParts {
		rect := game.Objects[i]
		rect.Move(fyne.NewPos(seg.x*10, seg.y*10))
	}

	game.Refresh()
}

// runGame is the main game loop that redraws the whole ui every
// time.Tick() duration, by calling refreshGame() after changing
// the y coordinate value (deducting it, causing the snake to go up
// since the new position every snake segment is being assigned to
// is decreasing
func runGame() {
	for range time.Tick(time.Millisecond * 500) {
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		snakeParts[0].y--
		refreshGame()
	}
}

// keyTyped essentially defines the key binds for snake movements
func keyTyped(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp, fyne.KeyK:
		move = moveUp
	case fyne.KeyDown, fyne.KeyJ:
		move = moveDown
	case fyne.KeyLeft, fyne.KeyH:
		move = moveLeft
	case fyne.KeyRight, fyne.KeyL:
		move = moveRight
	}
}
