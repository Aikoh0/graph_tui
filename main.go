package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify a filename")
		os.Exit(1)
	}
	var sorted bool = false
	for _, arg := range(os.Args) {
		if arg == "-s" {
			sorted = true
		}
	}
	var filepath string = os.Args[1]
	var sep string = ";"
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Clear()
			emitGraph(file, sep, sorted, screen)
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlR {
				screen.Clear()
				emitGraph(file, sep, sorted, screen)
				screen.Sync()
			}
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

