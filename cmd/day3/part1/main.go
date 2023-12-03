package main

import (
	"adventOfCode2023/internal"
	"fmt"
)

func main() {
	walkableFileReader := internal.NewWalkableFileReader("../simple")
	fmt.Printf("%+v\n", walkableFileReader)
	for i := 0; i < 15; i++ {
		char, _ := walkableFileReader.MoveEast()
		fmt.Println(string(char[0]))
	}
	res, err := walkableFileReader.ReadNorth()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadWest()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadEast()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadSouth()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadNorthWest()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadNorthEast()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadSouthWest()
	fmt.Println(string(res[0]), err)
	res, err = walkableFileReader.ReadSouthEast()
	fmt.Println(string(res[0]), err)

	fmt.Println(string(walkableFileReader.MoveNorth()[0]))
	fmt.Println(string(walkableFileReader.MoveWest()[0]))
	fmt.Println(string(walkableFileReader.MoveSouth()[0]))
	fmt.Println(string(walkableFileReader.MoveSouth()[0]))
	char, _ := walkableFileReader.MoveEast()
	fmt.Println(string(char[0]))
}
