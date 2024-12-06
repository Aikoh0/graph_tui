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


func sorted_bargraph(rows []string, sep string, g Graph, color_names []string) (Graph) {
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
	return g
}


func bar_graph(rows []string, sep string, g Graph, color_names []string) (Graph) {
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
	return g
}


func scatter_plot(rows []string, sep string, g Graph, color_names []string) (Graph) {
	for i, _ := range(rows) {
		var cols []string = strings.Split(rows[i], sep)
		cols = cols[1:]
		for j, col := range(cols) {
			g.nbr_col = len(cols)
			var color = color_names[j%g.nbr_col]
			var val, err = strconv.Atoi(col)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var bar = Bar{x: i, y: val, color: color}
			g.bars = append(g.bars, bar)
		}
		i++
	}
	return g
}

func create_graph(file string, sep string, graph_type string) (Graph) {
	var rows []string = strings.Split(file, "\n")
	rows = rows[:len(rows)-1]
	var g = Graph{}
	color_names := []string{"maroon","green","olive","navy","purple","teal","silver","gray","red","lime","yellow","blue"}
	switch graph_type {
		case "sort_bar":
		g = sorted_bargraph(rows, sep, g, color_names)
		case "scatter":
			g = scatter_plot(rows, sep, g, color_names)
		default:
		g = bar_graph(rows, sep, g, color_names)
	}
	return g
}

