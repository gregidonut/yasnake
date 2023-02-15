package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {

	a := app.New()

	w := a.NewWindow("YASnake")

	w.SetContent(setupGame())
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)

	w.ShowAndRun()
}

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
	}

	return container.NewWithoutLayout(segments...)
}
