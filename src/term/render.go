package term

import (
	"github.com/gdamore/tcell/v3"
)

func Render_Buffer_Generic(buffer [][]rune, screen tcell.Screen){
	x := 0
	y := 0
	for _, line := range buffer {
		PutString(screen, x, y, string(line))
		y++
	}
}
