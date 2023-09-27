package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Chen struct {
	postX, postY float32
	status       bool
	che          *canvas.Image
}

func NewChen(postx float32, posy float32, img *canvas.Image) *Chen {
	return &Chen{
		postX:  postx,
		postY:  posy,
		status: true,
		che:    img,
	}

}

func (c *Chen) Run() {
	var incX float32 = 50
	var incY float32 = 50
	c.status = true
	for c.status {
		if c.postX < 50 || c.postX > 1250 {
			incX *= -1
		}
		if c.postY < 50 || c.postY > 670 {
			incY *= -1
		}
		c.postX += incX
		c.postY += incY
		fmt.Println(c.postX, c.postY)
		c.che.Move(fyne.NewPos(c.postX, c.postY))
		time.Sleep((500 * time.Millisecond))
	}
}

func (c *Chen) Setstatus(status bool) {
	c.status = status
}
