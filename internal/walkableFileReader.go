package internal

import (
	"fmt"
	"os"
	"runtime"
)

var NEWLINE_OFFSET int64

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

type walkableFileReader struct {
	// where the reader is in the file
	location int64
	// the file to walk
	file *os.File
	// this is needed to enable "diagonal" walks
	width     int64
	character [1]byte
}

func NewWalkableFileReader(filePath string) *walkableFileReader {
	wr := &walkableFileReader{location: -1, width: 0, character: [1]byte{}}
	if runtime.GOOS == "windows" {
		NEWLINE_OFFSET = 2
	} else {
		NEWLINE_OFFSET = 1
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	wr.file = file
	character := make([]byte, 1)
	for {
		_, _ = wr.file.Read(character)
		switch character[0] {
		case '\n':
			fallthrough
		case '\r':
			_, _ = wr.file.Seek(0, 0)
			return wr
		default:
			wr.width++
		}
	}
}

func (wfr *walkableFileReader) MoveEast() byte {
	if newOffset, err := wfr.file.Seek(wfr.location+1, 0); err == nil {
		wfr.location = newOffset
		if _, err = wfr.file.ReadAt(wfr.character[:], wfr.location); err == nil {
			return wfr.character[0]
		} else {
			return 0
		}
	}
	return 0
}

func (wfr *walkableFileReader) MoveNorth() byte {
	fmt.Println(wfr.location, wfr.location-NEWLINE_OFFSET-wfr.width)
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveSouth() byte {
	if newOffset, err := wfr.file.Seek(wfr.location+wfr.width+NEWLINE_OFFSET, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveWest() byte {
	fmt.Println(wfr.location, wfr.location-1)
	if newOffset, err := wfr.file.Seek(wfr.location-1, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveNorthWest() byte {
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width-1, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveNorthEast() byte {
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width+1, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveSouthWest() byte {
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width-1, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) MoveSouthEast() byte {
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width+1, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(wfr.character[:], wfr.location)
		return wfr.character[0]
	}
	return 0
}

func (wfr *walkableFileReader) ReadDirection(direction Direction) (byte, error) {
	var err error
	switch direction {
	case NORTH:
		// North
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width)
	case NORTH_WEST:
		// North West
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width-1)
	case WEST:
		// West
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location-1)
	case SOUTH_WEST:
		// South West
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET-1)
	case SOUTH:
		// South
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET)
	case SOUTH_EAST:
		// South East
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET+1)
	case EAST:
		// East
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location+1)
	case NORTH_EAST:
		// North East
		_, err = wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width+1)
	}
	if err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadNorth read the character "above" the current location in the file by peeking a "width" distance back
func (wfr *walkableFileReader) ReadNorth() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadWest reads the character "behind" the current location by peeking one character back
func (wfr *walkableFileReader) ReadWest() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadEast reads the character "ahead" of the current location, if an error occurs it means the file has no more characters
func (wfr *walkableFileReader) ReadEast() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location+1); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadSouth reads the character "below" the current location by peeking a "width" distance ahead
func (wfr *walkableFileReader) ReadSouth() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadNorthWest reads the character "above" and "behind" the current location by peeking a "width"+1 distance behind the current location
func (wfr *walkableFileReader) ReadNorthWest() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width-1); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadNorthEast reads the character "above" and "ahead" of the current location by peeking a "width"-1 distance behind the current location
func (wfr *walkableFileReader) ReadNorthEast() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location-NEWLINE_OFFSET-wfr.width+1); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadSouthWest reads the character "below" and "behind" the current location by peeking a "width"+1 distance ahead of the current location
func (wfr *walkableFileReader) ReadSouthWest() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET-1); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

// ReadSouthEast reads the character "above" and "ahead" of the current location by peeking a "width"-1 distance ahead of the current location
func (wfr *walkableFileReader) ReadSouthEast() (byte, error) {
	if _, err := wfr.file.ReadAt(wfr.character[:], wfr.location+wfr.width+NEWLINE_OFFSET+1); err != nil {
		return 0, err
	}
	return wfr.character[0], nil
}

func (wfr *walkableFileReader) ResetFile() {
	_, _ = wfr.file.Seek(0, 0)
}
