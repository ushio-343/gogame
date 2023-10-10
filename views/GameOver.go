package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOver struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

func NewGameOver(window fyne.Window, app fyne.App) *GameOver {
	return &GameOver{window: window, app: app, container: container.NewWithoutLayout()}
}

func (s *GameOver) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/game_over.png"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))

	btnGoMenu := widget.NewButton("Principal menu", s.backMenu)
	btnGoMenu.Resize(fyne.NewSize(150, 30))
	btnGoMenu.Move(fyne.NewPos(325, 500))

	exitbutton := widget.NewButton("salir", s.Exit)
	exitbutton.Resize(fyne.NewSize(150, 30))
	exitbutton.Move(fyne.NewPos(325, 560))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnGoMenu))
}

func (s *GameOver) backMenu() {
	mainMenu := NewMainMenu(s.window, s.app)
	mainMenu.PrincipalMenu()
}

func (s *GameOver) Exit() {
	s.window.Close()
}
