package internal

import (
	"bufio"
	"os"
)

type inputIterator struct {
	scanner *bufio.Scanner
	err     error
}

// NewLineIterator creates a new inputIterator for a given file.
func NewLineIterator(filePath string) (*inputIterator, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return &inputIterator{
		scanner: bufio.NewScanner(file),
	}, nil
}

// Next advances the scanner to the next line.
// It returns false when the scan stops, either by reaching the end of the file or an error.
func (li *inputIterator) Next() bool {
	if li.scanner.Scan() {
		return true
	}
	li.err = li.scanner.Err()
	return false
}

// Line returns the most recent line generated by a call to Next.
func (li *inputIterator) Line() string {
	return li.scanner.Text()
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (li *inputIterator) Err() error {
	return li.err
}
