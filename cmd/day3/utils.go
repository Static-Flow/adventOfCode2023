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
