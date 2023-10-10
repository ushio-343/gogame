package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Marisa struct {
	mari *canvas.Image
	X, Y float32
}

func NewMarisa(x float32, y float32, img *canvas.Image) *Marisa {
	return &Marisa{
		mari: img,
		X:    x,
		Y:    y,
	}

}

func (m *Marisa) Move(e *fyne.KeyEvent, container *fyne.Container) {
	switch e.Name {
	case fyne.KeyUp:
		m.Y -= 10
	case fyne.KeyDown:
		m.Y += 10
	case fyne.KeyLeft:
		m.X -= 10
	case fyne.KeyRight:
		m.X += 10
	}
	fmt.Println(m.Y, m.X)
	m.mari.Move(fyne.NewPos(m.X, m.Y))
	container.Refresh()
}
