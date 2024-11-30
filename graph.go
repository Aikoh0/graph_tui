package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
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
// TODO 2 functions for clarity for sorted
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

