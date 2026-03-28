package term

import (
	"log"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func CreateStyle(bg color.Color, fg color.Color) tcell.Style {
	return tcell.StyleDefault.Background(bg).Foreground(fg)
}

func BaseStyle() tcell.Style {
	return tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)
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
