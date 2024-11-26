package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Coord struct {
	x int
	y int
}

type Bar struct {
	x int
	y int
	base int
	coords []Coord
}

type Graph struct {
	bars []Bar
}

func create_graph(file string, sep string) (Graph) {
	var rows []string = strings.Split(file, "\n")
	rows = rows[:len(rows)-1]
	var g = Graph{}
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

func bar_height(max_h int, screen_h int) (int) {
	return (screen_h / max_h) - (screen_h / 10)
}

func generate_coords(graph Graph, bar_h int, bar_w int) ([]Coord) {
	var coords []Coord
	for _, bar := range(graph.bars) {
		var j int = 0
		for j < bar_w {
			var h int = 0
			for h <bar_h {
				var c = Coord{x: bar.base + j, y: bar.y + h}
				coords = append(coords, c)	
				h++
			}
			j++
		}
	}
	return coords
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
	var g Graph = create_graph(string(file), sep)
	var max_h int = max_height(g)
	var bar_h int = bar_height(max_h, 50)
	var bar_w int = bar_width(g, 100)
	g = define_pos(g, bar_w)
	// TODO: Everything working until here, need verification from here
	//Seams like it works 
	var coords []Coord = generate_coords(g, bar_h, bar_w)
	fmt.Println("max_h:",max_h)
	fmt.Println("bar_width:",bar_w)
	fmt.Println("bar_height:",bar_h)
	fmt.Println("g:", g)
	fmt.Println("coords:", coords)
}

