package term

import (
	"log"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func InputQueue(screen tcell.Screen) bool {
	event := <-screen.EventQ()
	switch event := event.(type){
		case *tcell.EventResize: {
			ScrSync(screen)
		} 

		case *tcell.EventKey: {
			if event.Key() == tcell.KeyEscape {
				return true 
			}
		}
	}
	return false
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

