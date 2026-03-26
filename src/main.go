package main

import (
	"fmt"
	"os"
	"main/term"
	"main/sys"
	"github.com/gdamore/tcell/v3"
)

func ParseBytes(buffer []byte) [][]rune {
	const delimiter = '\n'
	const skip = '\r'

	line_buffer := make([][]rune, 0)
	byte_str := make([]byte, 0)

	for _, byte := range buffer {
		if byte != delimiter && byte != skip {
			byte_str = append(byte_str, byte)
		} else if byte == delimiter {
			line_buffer = append(line_buffer, []rune(string(byte_str)))
			byte_str = byte_str[:0]
		}
	}

	if len(byte_str) > 0 {
		line_buffer = append(line_buffer, []rune(string(byte_str)))
	}

	return line_buffer
}

func main(){
	args := os.Args
	filepath := ""
	if len(args) > 1 && len(args) < 3 {
		filepath = args[1]
	}

	file := sys.OpenFile(filepath)
	data := sys.ReadFile(file)
	line_buffer := ParseBytes(data)
	
	for _, line := range line_buffer {
		fmt.Println(string(line))
	}

	sys.CloseFile(file)

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

	}
}
