package day3

var WalkableReader *walkableByteReader

func IsEnginePart(character byte) bool {
	if character != 46 && (character >= 58 || (character >= 33 && character <= 47)) {
		return true
	}
	return false
}

func IsEngineGear(character byte) bool {
	if character == 42 {
		return true
	}
	return false
}

// SearchNumberNeighbours searches each character in a cardinal direction counter-clockwise from the current character and if any are engine parts returns true
func SearchNumberNeighbours(directions []Direction) bool {
	var searchedCharacter byte
	var DIRECTION Direction
	for _, DIRECTION = range directions {
		searchedCharacter = WalkableReader.ReadOrMoveDirection(false, DIRECTION)
		if IsEnginePart(searchedCharacter) {
			return true
		}
	}
	return false
}
