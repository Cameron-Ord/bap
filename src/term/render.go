package term

import (
	"github.com/gdamore/tcell/v3"
)

func Fill(rows int, cols int, screen tcell.Screen, style tcell.Style){
	for row := 0; row <= rows; row++ {
		for col := 0; col <= cols; col++{
			PutChar(screen, col, row, " ", style)
		}
	}
}
	
func Render_Buffer_Generic(buffer [][]rune, _x int, _y int, screen tcell.Screen, style tcell.Style){
	row := 0
	section := buffer[_y:]

	for _, line := range section {
		col := 0 
		str := string(line)
		width := 0
		for str != "" {
			str, width = PutChar(screen, col, row, str, style)
			col += width
			if width == 0 {
				break
			}
		}
		row++
	}
}
