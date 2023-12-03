package day2

var BLUE = "blue"
var RED = "red"
var GREEN = "green"

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
