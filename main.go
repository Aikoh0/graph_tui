package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
)

type Coord struct {
	x int
	y int
	color string
}

type Bar struct {
	x int
	y int
	base int
	coords []Coord
}

type Graph struct {
	bars []Bar
	nbr_col int
}

// BLOCKING TODO: Handle bar color here w\ bar.color attr instead of coords.color
// TODO:Color should be handled differently for sorted graphs
// TODO Split in multiple file and split handle in 2 functions for clarity
func create_graph(file string, sep string) (Graph) {
	var sorted bool = true // TODO: Add sorted parameter and change color accordingly
	var rows []string = strings.Split(file, "\n")
	rows = rows[:len(rows)-1]
	var g = Graph{}
	
	if sorted {
		var row_nbr int = 0
		var col_nbr int = 1
		for row_nbr < len(rows) {
			var row []string = strings.Split(rows[row_nbr], sep)
			val, _ := strconv.Atoi(row[col_nbr])
			var bar = Bar{x: row_nbr, y: val}
			g.bars = append(g.bars, bar)
			g.nbr_col = len(row)
			row_nbr++
			if row_nbr == len(rows) && col_nbr != len(row) - 1 {
				row_nbr = 0
				col_nbr++
			}
		}
	} else {
		for i, _ := range(rows) {
			var row []string = strings.Split(rows[i], sep)
			row = row[1:]
			for _, val := range(row) {
				var y, err = strconv.Atoi(val)
				if err != nil {
					fmt.Println("converison error:", err)
					os.Exit(1)
				}
				var b = Bar{x: i, y: y}
				g.bars = append(g.bars, b)
				g.nbr_col = len(row)
			}	
		}
	}
	return g
}

func define_pos(graph Graph, bar_w int) (Graph) {
	for i, _ := range(graph.bars) {
		var base int = 1+i+(bar_w*i)
		graph.bars[i].base = base 
	}
	return graph
}

func max_height(graph Graph) (int) {
	var max_h = graph.bars[0].y
	for _, bar := range(graph.bars) {
		if bar.y > max_h {
			max_h = bar.y
		}
	}
	return max_h
}

func bar_width(graph Graph, screen_w int) (int) {
	return (screen_w - len(graph.bars)) / len(graph.bars)
}

// TODO: Weird size given by s.Size() need to look into it
func bar_height(max_h int, screen_h int) (int) {
	return (screen_h / max_h)+ (screen_h / 10)
}

func generate_coords(graph Graph, bar_h int, bar_w int) ([]Coord) {
	var coords []Coord
	color_names := []string{"maroon","green","olive","navy","purple","teal","silver","gray","red","lime","yellow","blue"}
	for i, bar := range(graph.bars) {
		var j int = 0
		for j < bar_w {
			var h int = 0
			for h < bar.y *bar_h {
				var c = Coord{x: bar.base + j, y: 1+h, color:color_names[i%graph.nbr_col]}
				coords = append(coords, c)	
				h++
			}
			j++
		}
	}
	return coords
}

func emitStr(s tcell.Screen, style tcell.Style, coords []Coord) {
	var _, h int = s.Size()
	for _, coord := range coords {
	style = tcell.StyleDefault.Foreground(tcell.GetColor(coord.color)).Background(tcell.GetColor(coord.color))
		var comb []rune	
		s.SetContent(coord.x, h - coord.y, rune(0), comb, style)
	}
}

func main() {
	if len(os.Args) > 2 {
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
	fmt.Print("max_h: ", max_h)
	fmt.Print("g: ", g)
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

