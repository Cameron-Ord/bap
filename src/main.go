package main

import (
	"main/term"
)

func main(){
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
		if(term.InputQueue(screen)){
			return
		}
	}
}
