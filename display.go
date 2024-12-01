package main

import (
	"github.com/gdamore/tcell"
)


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
	return (screen_h / max_h)- (screen_h / 20)
}


func generate_coords(graph Graph, bar_h int, bar_w int) ([]Coord) {
	var coords []Coord
	for _, bar := range(graph.bars) {
		var j int = 0
		for j < bar_w {
			var h int = 0
			for h < bar.y *bar_h {
				var c = Coord{x: bar.base + j, y: 1+h, color:bar.color}
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

