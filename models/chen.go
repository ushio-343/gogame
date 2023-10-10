package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Chen struct {
	PostX, PostY float32
	status       bool
	che          *canvas.Image
}

func NewChen(postx float32, posy float32, img *canvas.Image) *Chen {
	return &Chen{
		PostX:  postx,
		PostY:  posy,
		status: true,
		che:    img,
	}

}

func (c *Chen) Run() {
	var incX float32 = 50
	var incY float32 = 50
	c.status = true
	for c.status {
		if c.PostX < 50 || c.PostX > 1250 {
			incX *= -1
		}
		if c.PostY < 50 || c.PostY > 670 {
			incY *= -1
		}
		c.PostX += incX
		c.PostY += incY
		fmt.Println(c.PostX, c.PostY)
		c.che.Move(fyne.NewPos(c.PostX, c.PostY))
		time.Sleep((500 * time.Millisecond))
	}
}

func (c *Chen) Setstatus(status bool) {
	c.status = status
}
