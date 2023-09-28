package views

import (
	"demo/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

type GameScene struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

var c *models.Chen
var m *models.Marisa

func NewGameScene(window fyne.Window, app fyne.App) *GameScene {

	return &GameScene{window: window, app: app, container: container.NewWithoutLayout()}
}

func (g *GameScene) music() {
	for {
		file, err := os.Open("assets/backmusic.mp3")
		if err != nil {
			log.Panicln(err)
		}
		defer file.Close()

		streamer, format, err := mp3.Decode(file)
		if err != nil {
			log.Panicln(err)
		}
		defer streamer.Close()

		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		<-done
	}
}

func (g *GameScene) startscene() {
	
	image := canvas.NewImageFromURI(storage.NewFileURI("./assets/gensokyo.png"))
	image.Resize(fyne.NewSize(1366, 768))
	image.Move(fyne.NewPos(0, 0))

	marisa := canvas.NewImageFromFile("./assets/Marisa.png")
	marisa.Resize(fyne.NewSize(70, 70))
	marisa.Move(fyne.NewPos(50, 50))

	m = models.NewMarisa(50, 50, marisa)

	chen := canvas.NewImageFromURI(storage.NewFileURI("./assets/chen_asset.png"))
	chen.Resize(fyne.NewSize(70, 70))
	chen.Move(fyne.NewPos(100, 100))

	c = models.NewChen(100, 100, chen)

	go c.Run()

	go g.music()

	g.window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		go m.Move(e, g.container)
		fmt.Println("te")
	})

	stopbotton := widget.NewButton("detener el juego", g.StopGame)
	stopbotton.Resize(fyne.NewSize(150, 30))
	stopbotton.Move(fyne.NewPos(300, 10))

	exit := widget.NewButton("salir", g.ExitGame)
	exit.Resize(fyne.NewSize(150, 30))
	exit.Move(fyne.NewPos(460, 10))

	g.window.SetContent(container.NewWithoutLayout(image, chen, marisa, stopbotton, exit))
}

func (g *GameScene) StopGame() {
	c.Setstatus(false)
}

func (g *GameScene) ExitGame() {
	g.window.Close()
}
