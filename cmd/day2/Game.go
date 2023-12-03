package day2

import (
	"fmt"
	"strings"
)

type Game struct {
	Id     int
	Rounds []*Round
}

func (g *Game) String() string {
	if len(g.Rounds) == 0 {
		return fmt.Sprintf("Game ID:%d -- Skipped for non conformance", g.Id)
	}
	var str strings.Builder
	for _, round := range g.Rounds {
		str.WriteString(round.String())
	}
	return fmt.Sprintf("Game ID:%d\nGames:\n%s", g.Id, str.String())
}

type Round struct {
	Cubes []*Cube
}

func (r *Round) String() string {
	var str strings.Builder
	for _, cube := range r.Cubes {
		str.WriteString(cube.String())
	}
	return fmt.Sprintf("Round:\n%s", str.String())
}

type Cube struct {
	CubeColor string
	Count     int
}

func (c *Cube) String() string {
	return fmt.Sprintf("\tCount: %d\n\tColor: %s\n", c.Count, c.CubeColor)
}

var BLUE = "blue"
var RED = "red"
var GREEN = "green"

var ColorMap = map[string]int{BLUE: 14, RED: 12, GREEN: 13}

type GamePart2 struct {
	MaxColorsMap map[string]int
}

func (g *GamePart2) ProcessPull(pullLine string) {
	var count int
	var color string
	var max int
	if pullLine[1] == ' ' {
		count = int(pullLine[0] - '0')
		color = pullLine[2:]
		if max, _ = g.MaxColorsMap[color]; count > max {
			g.MaxColorsMap[color] = count
		}
	} else {
		count = (int(pullLine[0]-'0') * 10) + int(pullLine[1]-'0')
		color = pullLine[3:]
		if max, _ = g.MaxColorsMap[color]; count > max {
			g.MaxColorsMap[color] = count
		}
	}
}
