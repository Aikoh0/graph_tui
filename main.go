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
	sc_w, sc_h := screen.Size()
	var g Graph = create_graph(string(file), sep)
	var max_h int = max_height(g)
	var bar_h int = bar_height(max_h, sc_h)
	var bar_w int = bar_width(g, sc_w)

	g = define_pos(g, bar_w)

	var coords []Coord = generate_coords(g, bar_h, bar_w)

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Clear()
			emitStr(screen, defStyle, coords)
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlR {
				screen.Clear()
				emitStr(screen, defStyle, coords)
				screen.Sync()
			}
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

