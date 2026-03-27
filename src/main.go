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
			case *tcell.EventResize: {
				term.ScrSync(screen)
			} 

			case *tcell.EventKey: {
				if event.Key() == tcell.KeyEscape {
					return
				}
			}
		}

		term.Render_Buffer_Generic(list.Buffers[list.Index].Data, screen)

	}
}
