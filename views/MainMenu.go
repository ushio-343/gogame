package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainMenu struct {
	window fyne.Window
	app    fyne.App
}

func NewMainMenu(window fyne.Window, app fyne.App) *MainMenu {
	return &MainMenu{window: window, app: app}
}

func (m *MainMenu) PrincipalMenu() {

	titleLabel := widget.NewLabel("daimmaku")
	titleLabel.Move(fyne.NewPos(680, 50))

	start := widget.NewButton("comensar", m.startGame)
	start.Resize(fyne.NewSize(200, 50))
	start.Move(fyne.NewPos(650, 100))

	exit := widget.NewButton("salir", m.exitGame)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(650, 180))

	m.window.SetContent(container.NewWithoutLayout(start, exit, titleLabel))
}

func (m *MainMenu) startGame() {
	gamescene := NewGameScene(m.window, m.app)
	gamescene.startscene()

}

func (m *MainMenu) exitGame() {
	m.window.Close()
}
