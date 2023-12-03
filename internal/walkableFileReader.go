package internal

import (
	"os"
	"runtime"
)

var NEWLINE_OFFSET int64

type walkableFileReader struct {
	// where the reader is in the file
	location int64
	// the file to walk
	file *os.File
	// this is needed to enable "diagonal" walks
	width int64
}

func NewWalkableFileReader(filePath string) *walkableFileReader {
	wr := &walkableFileReader{location: 0, width: 0}
	if runtime.GOOS == "windows" {
		NEWLINE_OFFSET = 1
	} else {
		NEWLINE_OFFSET = 0
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	wr.file = file
	character := make([]byte, 1)
	for {
		_, _ = wr.file.Read(character)
		if character[0] != '\n' {
			wr.width++
		} else {
			break
		}
	}
	_, _ = wr.file.Seek(0, 0)
	return wr
}

func (wfr *walkableFileReader) MoveEast() (character [1]byte, err error) {
	wfr.location++
	_, err = wfr.file.ReadAt(character[:], wfr.location)
	return character, err
}

func (wfr *walkableFileReader) MoveNorth() (character [1]byte) {
	if newOffset, err := wfr.file.Seek(wfr.location-NEWLINE_OFFSET-wfr.width, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(character[:], wfr.location)
	}
	return character
}

func (wfr *walkableFileReader) MoveSouth() (character [1]byte) {
	if newOffset, err := wfr.file.Seek(wfr.location+wfr.width+NEWLINE_OFFSET, 0); err == nil {
		wfr.location = newOffset
		_, _ = wfr.file.ReadAt(character[:], wfr.location)
	}
	return character
}

func (wfr *walkableFileReader) MoveWest() (character [1]byte) {
	if newOffset, err := wfr.file.Seek(wfr.location-1, 0); err == nil {
		wfr.location = newOffset
	}
	_, _ = wfr.file.Read(character[:])
	return character
}

// ReadNorth read the character "above" the current location in the file by peeking a "width" distance back
func (wfr *walkableFileReader) ReadNorth() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location-NEWLINE_OFFSET-wfr.width)
	return character, err
}

// ReadEast reads the character "ahead" of the current location, if an error occurs it means the file has no more characters
func (wfr *walkableFileReader) ReadEast() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location+1)
	return character, err
}

// ReadWest reads the character "behind" the current location by peeking one character back
func (wfr *walkableFileReader) ReadWest() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location-NEWLINE_OFFSET-1)
	return character, err
}

// ReadSouth reads the character "below" the current location by peeking a "width" distance ahead
func (wfr *walkableFileReader) ReadSouth() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location+wfr.width+NEWLINE_OFFSET)
	return character, err
}

// ReadNorthWest reads the character "above" and "behind" the current location by peeking a "width"+1 distance behind the current location
func (wfr *walkableFileReader) ReadNorthWest() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location-NEWLINE_OFFSET-wfr.width-1)
	return character, err
}

// ReadNorthEast reads the character "above" and "ahead" of the current location by peeking a "width"-1 distance behind the current location
func (wfr *walkableFileReader) ReadNorthEast() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location-NEWLINE_OFFSET-wfr.width+1)
	return character, err
}

// ReadSouthWest reads the character "below" and "behind" the current location by peeking a "width"+1 distance ahead of the current location
func (wfr *walkableFileReader) ReadSouthWest() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location+wfr.width+NEWLINE_OFFSET-1)
	return character, err
}

// ReadSouthEast reads the character "above" and "ahead" of the current location by peeking a "width"-1 distance ahead of the current location
func (wfr *walkableFileReader) ReadSouthEast() (character [1]byte, err error) {
	_, err = wfr.file.ReadAt(character[:], wfr.location+wfr.width+NEWLINE_OFFSET+1)
	return character, err
}
