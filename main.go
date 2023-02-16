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
	snakeParts   []snakePart
	game         *fyne.Container
	head         *canvas.Rectangle
	loopDuration = time.Millisecond * 500
	move         = moveUp
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

	// add a new rectangle that will represent the snake segments
	// (to be animated)
	head = canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
	head.Resize(fyne.NewSize(10, 10))
	head.Move(fyne.NewPos(snakeParts[0].x*10, snakeParts[0].y*10))
	segments = append(segments, head)

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
// the appropriate coordinate value.
// Calculating an old and new position for the head and tail so that
// a *fyne.Animation can be used to render the animation
func runGame() {
	// nextPart is the calculation of the snake's head and
	// will be used by the head's animation as coordinates of
	// where the head rectangle will end in the animation
	nextPart := snakePart{snakeParts[0].x, snakeParts[0].y - 1}

	for {
		oldPos := fyne.NewPos(snakeParts[0].x*10, snakeParts[0].y*10)
		newPos := fyne.NewPos(nextPart.x*10, nextPart.y*10)
		canvas.NewPositionAnimation(oldPos, newPos, loopDuration, func(p fyne.Position) {
			head.Move(p)
			canvas.Refresh(head)
		}).Start()

		end := len(snakeParts) - 1
		canvas.NewPositionAnimation(
			fyne.NewPos(snakeParts[end].x*10, snakeParts[end].y*10),
			fyne.NewPos(snakeParts[end-1].x*10, snakeParts[end-1].y*10),
			loopDuration,
			func(p fyne.Position) {
				tail := game.Objects[end]
				tail.Move(p)
				canvas.Refresh(tail)
			},
		).Start()

		time.Sleep(loopDuration)
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		snakeParts[0] = nextPart

		switch move {
		case moveUp:
			nextPart = snakePart{
				x: nextPart.x,
				y: nextPart.y - 1,
			}
		case moveDown:
			nextPart = snakePart{
				x: nextPart.x,
				y: nextPart.y + 1,
			}
		case moveLeft:
			nextPart = snakePart{
				x: nextPart.x - 1,
				y: nextPart.y,
			}
		case moveRight:
			nextPart = snakePart{
				x: nextPart.x + 1,
				y: nextPart.y,
			}
		}
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
