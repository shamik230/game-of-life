package game

import (
	"fmt"

	"github.com/shamik230/game-of-life/screen"
)

func NewGame(gameScreen *screen.Screen) {
	gameScreen.SetLatencyMs(100)

	var (
		width  = gameScreen.GetWidth()
		height = gameScreen.GetHeight()
	)

	initialField, err := screen.NewField(width, height)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gameScreen.SetField(initialField)

	for {
		initialField.SetRandomWith([]rune("🟠🟡🟢🟣🟤⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪⚪"))
		gameScreen.SetField(initialField)
		gameScreen.RenderFrame()
	}
}
