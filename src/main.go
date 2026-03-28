package main

import (
	"os"
	"main/term"
	"main/internal"
	"main/util"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)


func main(){
	args := os.Args
	filepath := ""
	if len(args) > 1 && len(args) < 3 {
		filepath = args[1]
	}

	style := term.CreateStyle(color.Black, color.Red)
	list := internal.Make_List()
	file := internal.OpenFile(filepath)
	internal.Append_Buffer(util.ParseBytes(internal.ReadFile(file)), filepath, &list)
	internal.CloseFile(file)

	screen := term.CreateScreen()
	term.Initialize(screen)
	screen.SetStyle(style)
	screen.Clear()
	term.Render_Buffer_Generic(list.Buffers[list.Index].Data, 0, 0, screen, style)

	quit := func() {
		maybe_panic := recover()
		screen.Fini()
		if maybe_panic != nil {
			panic(maybe_panic)
		}
	}
	defer quit()

	for {
		screen.Show()
		event := <-screen.EventQ()
		switch event := event.(type){
			case *tcell.EventResize: {
				screen.Sync()
			} 

			case *tcell.EventKey: {
				switch event.Key() {
					case tcell.KeyEscape: {
						return
					}
					case tcell.KeyUp: {
						internal.Cursor_Add_Y(-1, &list.Buffers[list.Index])
						screen.Clear()
						term.Render_Buffer_Generic(list.Buffers[list.Index].Data, 0, list.Buffers[list.Index].Cursor.Y, screen, style)
					}
					case tcell.KeyDown: {
						internal.Cursor_Add_Y(1, &list.Buffers[list.Index])
						screen.Clear()
						term.Render_Buffer_Generic(list.Buffers[list.Index].Data, 0, list.Buffers[list.Index].Cursor.Y, screen, style)
					}
					case tcell.KeyRight: {
					}
					case tcell.KeyLeft: {
					}
				}
			}
		}
		
	}
}
