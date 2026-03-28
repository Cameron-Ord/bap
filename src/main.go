package main

import (
	"os"
	"main/term"
	"main/internal"
	"main/util"
	"github.com/gdamore/tcell/v3"
)


func main(){
	args := os.Args
	filepath := ""
	if len(args) > 1 && len(args) < 3 {
		filepath = args[1]
	}
	list := internal.Make_List()
	
	file := internal.OpenFile(filepath)
	internal.Append_Buffer(util.ParseBytes(internal.ReadFile(file)), filepath, &list)
	internal.CloseFile(file)

	screen := term.CreateScreen()
	term.Initialize(screen)
	term.ScrSetStyle(screen, term.BaseStyle())
	term.ClearScreen(screen)

	quit := func() {
		maybe_panic := recover()
		screen.Fini()
		if maybe_panic != nil {
			panic(maybe_panic)
		}
	}
	defer quit()

	for {
		term.ScrShow(screen)
	
		event := <-screen.EventQ()
		switch event := event.(type){
		default:
			case *tcell.EventResize: {
				term.ScrSync(screen)
			} 

			case *tcell.EventKey: {
				switch event.Key() {
					case tcell.KeyEscape: {
						return
					}
					case tcell.KeyUp: {
						internal.Cursor_Add_Y(-1, &list.Buffers[list.Index])
					}
					case tcell.KeyDown: {
						internal.Cursor_Add_Y(1, &list.Buffers[list.Index])
					}
					case tcell.KeyRight: {}
					case tcell.KeyLeft: {}
				}
			}
		}
		buf := &list.Buffers[list.Index]
		term.Render_Buffer_Generic(buf.Data, buf.Cursor.X, buf.Cursor.Y, screen, term.BaseStyle())
	}
}
