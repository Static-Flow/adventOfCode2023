package day3

import (
	"log"
	"os"
	"runtime"
)

var OFFSET = 2
var newOffset int

type Direction int

const (
	NORTH Direction = iota
	NORTH_WEST
	WEST
	SOUTH_WEST
	SOUTH
	SOUTH_EAST
	EAST
	NORTH_EAST
)

var CARDINAL_DIRECTIONS = []Direction{NORTH, NORTH_WEST, WEST, SOUTH_WEST, SOUTH, SOUTH_EAST, EAST, NORTH_EAST}

var QUICK_CARDINAL_DIRECTIONS = []Direction{SOUTH_EAST, EAST, NORTH_EAST}

type walkableByteReader struct {
	// where the reader is in the file
	location int
	// the file to walk
	// this is needed to enable "diagonal" walks
	width int
	size  int
	data  []byte
	lines int
}

func NewWalkableByteStream(filePath string) *walkableByteReader {
	wr := &walkableByteReader{location: -1, width: 0}
	if runtime.GOOS == "windows" {
		OFFSET = 2
	} else {
		OFFSET = 1
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	wr.data = file
	wr.size = len(wr.data)
	for i := 0; i < len(wr.data); i++ {
		switch wr.data[i] {
		case '\n':
			fallthrough
		case '\r':
			return wr
		default:
			wr.width++
		}
	}
	wr.lines = (wr.size / wr.width) - 1
	return wr
}

func (wbr *walkableByteReader) NumOfLines() int {
	return wbr.lines
}

func (wbr *walkableByteReader) LastSearchLocation() int {
	return newOffset
}

func (wbr *walkableByteReader) Reset() {
	wbr.location = -1
}

func (wbr *walkableByteReader) SetLocation(location int) {
	wbr.location = location
}

func (wbr *walkableByteReader) GetLocation() int {
	return wbr.location
}

func (wbr *walkableByteReader) ReadAtLocation(location int) byte {
	if location < 0 || location > wbr.size {
		return 0
	}
	return wbr.data[location]
}

func (wbr *walkableByteReader) ReadOrMoveDirection(move bool, direction Direction) byte {
	switch direction {
	case NORTH:
		newOffset = wbr.location - OFFSET - wbr.width
	case NORTH_WEST:
		newOffset = wbr.location - OFFSET - wbr.width - 1
	case WEST:
		newOffset = wbr.location - 1
	case SOUTH_WEST:
		newOffset = wbr.location + wbr.width + OFFSET - 1
	case SOUTH:
		newOffset = wbr.location + wbr.width + OFFSET
	case SOUTH_EAST:
		newOffset = wbr.location + wbr.width + OFFSET + 1
	case EAST:
		newOffset = wbr.location + 1
	case NORTH_EAST:
		newOffset = wbr.location - OFFSET - wbr.width + 1
	}
	if newOffset < 0 || newOffset >= wbr.size {
		return 0
	} else {
		if move {
			wbr.location = newOffset
			return wbr.data[wbr.location]
		} else {
			return wbr.data[newOffset]
		}
	}
}
