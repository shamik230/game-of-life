package screen

import (
	"fmt"
	"time"
)

type Screen struct {
	width int
	height int
	field Field
	latency time.Duration
}

func NewScreen(width int, height int) (*Screen, error) {
	field, err := NewField(width, height)
	
	if err != nil {
		return nil, err
	}

	return &Screen{
		width: width,
		height: height,
		field: field,
		latency: time.Millisecond * 50,
	}, nil
}

func (s *Screen) SetLatencyMs(ms int) {
	if ms < 0 {
		ms = 0
	}
	s.latency = time.Millisecond * time.Duration(ms);
}

func (s *Screen) RenderFrame() {
	fmt.Print("\033[H\033[2J")
	for _, v := range s.field { 
		fmt.Println(string(v))
	}
	if s.latency > 0 {
		time.Sleep(s.latency)
	}
}

func (s *Screen) SetField(f Field) {
	if f == nil {
		return
	}
	for h := 0; h < s.height; h++ {
		if h > len(f) - 1 {
			break
		}
		for w := 0; w < s.width; w++ {
			if w > len(f[h]) - 1 {
				break
			}
			s.field[h][w] = f[h][w]
		}
	}
}

func (s *Screen) GetWidth() int {
	return s.width
}

func (s *Screen) GetHeight() int {
	return s.height
}