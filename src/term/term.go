package term

import (
	"log"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func PutStringStyled(screen tcell.Screen, x int, y int, str string, style tcell.Style){
	screen.PutStrStyled(x, y, str, style)
}

func PutString(screen tcell.Screen, x int, y int, str string){
	screen.PutStr(x, y, str)
}

func PutChar(screen tcell.Screen, x int, y int, character string, style tcell.Style){
	screen.Put(x, y, character, style)
}

func BaseStyle() tcell.Style {
	return tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
}

func ScrSync(screen tcell.Screen) {
	screen.Sync()
}

func ScrShow(screen tcell.Screen) {
	screen.Show()
}

func ScrSetStyle(screen tcell.Screen, style tcell.Style) {
	screen.SetStyle(style)
}

func CreateScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return screen
}

func Initialize(screen tcell.Screen)  {
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func ClearScreen(screen tcell.Screen){
	screen.Clear()
}

