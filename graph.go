package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)


type Bar struct {
	x int
	y int
	base int
	color string
	coords []Coord
}


type Graph struct {
	bars []Bar
	nbr_col int
}


// TODO 2 functions for clarity for sorted
func create_graph(file string, sep string, sorted bool) (Graph) {
	var rows []string = strings.Split(file, "\n")
	rows = rows[:len(rows)-1]
	var g = Graph{}

	color_names := []string{"maroon","green","olive","navy","purple","teal","silver","gray","red","lime","yellow","blue"}
	if sorted {
		var row_nbr int = 0
		var col_nbr int = 1
		for row_nbr < len(rows) {
			var color = color_names[col_nbr]
			var row []string = strings.Split(rows[row_nbr], sep)
			val, _ := strconv.Atoi(row[col_nbr])
			var bar = Bar{x: row_nbr, y: val, color:color}
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
			for j, val := range(row) {
				g.nbr_col = len(row)
				var color = color_names[j%g.nbr_col]
				var y, err = strconv.Atoi(val)
				if err != nil {
					fmt.Println("converison error:", err)
					os.Exit(1)
				}
				var b = Bar{x: i, y: y, color: color}
				g.bars = append(g.bars, b)
			}	
		}
	}
	return g
}

