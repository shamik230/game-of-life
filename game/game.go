package game

import (
	"fmt"
	"math"

	"github.com/shamik230/game-of-life/screen"
)

type gameField struct {
	prev screen.Field
	next screen.Field
}

const (
	OFF = '🌕'
	ON  = '🌑'
)

func NewGame(gameScreen *screen.Screen) {
	var (
		width  = gameScreen.GetWidth()
		height = gameScreen.GetHeight()
	)

	gameField, err := newGameField(width, height)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gameField.next.SetRandomWith([]rune{OFF, OFF, OFF, OFF, ON}) // initial field

	for {
		gameScreen.SetField(gameField.next)
		gameScreen.RenderFrame()
		gameField.evalNextFrame()
	}
}

func newGameField(width int, height int) (*gameField, error) {
	prev, err := screen.NewField(width, height)
	if err != nil {
		return nil, err
	}

	next, err := screen.NewField(width, height)
	if err != nil {
		return nil, err
	}

	return &gameField{
		prev: prev,
		next: next,
	}, nil
}

func (f *gameField) evalNextFrame() {
	f.prev, f.next = f.next, f.prev

	for i := range f.prev {
		for j := range f.prev[i] {
			count := countNeighbors(f.prev, i, j)
			if f.prev[i][j] == OFF {
				if count == 3 {
					f.next[i][j] = ON
				} else {
					f.next[i][j] = OFF
				}
			} else {
				if count == 2 || count == 3 {
					f.next[i][j] = ON
				} else {
					f.next[i][j] = OFF
				}
			}
		}
	}
}

func modulo(x, base int) int {
	y := math.Floor(float64(x) / float64(base))
	return x - int(y)*base
}

func countNeighbors(f screen.Field, i, j int) (count int) {
	var (
		I int
		J int
	)
	for h := -1; h <= 1; h++ {
		for w := -1; w <= 1; w++ {
			if h == 0 && w == 0 {
				continue
			}

			I = modulo(i+h, len(f))
			J = modulo(j+w, len(f[0]))

			if f[I][J] == ON {
				count++
			}
		}
	}
	return
}
