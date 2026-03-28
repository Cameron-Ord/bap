package term

import (
	"github.com/gdamore/tcell/v3"
)

func Fill(rows int, cols int, screen tcell.Screen, style tcell.Style){
	for row := range rows {
		for col := range cols{
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}
	
func Render_Buffer_Generic(buffer [][]rune, offset_x int, offset_y int, screen tcell.Screen, style tcell.Style){
	row := 0
	for _, line := range buffer[offset_y:]{
		col := 0
		for _, rune := range line {
			screen.SetContent(col, row, rune, nil, style)
			col++
		}
		row++
	}
}
