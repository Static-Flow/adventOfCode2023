package day4

import (
	"github.com/Static-Flow/adventOfCode2023/internal"
	"log"
	"strings"
)

var candidatesMap = [100]bool{}

func checkWinner(storedDigit int) bool {
	if candidatesMap[storedDigit] {
		return true
	}
	return false
}

type ScratchCards struct {
	Cards []*Card
}

func NewScratchCards(fileInput string) *ScratchCards {
	sc := &ScratchCards{}
	index := 1
	if iterator, err := internal.NewLineIterator(fileInput); err == nil {
		for iterator.Next() {
			sc.Cards = append(sc.Cards, NewCard(index, strings.Split(iterator.Line(), ":")[1]))
			index++
		}
		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	return sc
}

type Card struct {
	index   int
	Matches int
	Count   int
	line    string
}

func NewCard(index int, line string) *Card {
	c := &Card{
		index: index,
		line:  line,
		Count: 1,
	}
	c.ProcessCard()
	return c
}

func (c *Card) ProcessCard() {
	var storedDigit = 0
	var checkingWinners bool
	var winningNumbers = [10]int{}
	var index int
	for _, character := range c.line {
		if internal.IsRuneANumber(character) {
			// found a number
			if storedDigit == 0 {
				storedDigit = int(character - '0')
			} else {
				storedDigit = storedDigit*10 + int(character-'0')
			}
		} else {
			if storedDigit != 0 {
				if !checkingWinners {
					winningNumbers[index] = storedDigit
					candidatesMap[storedDigit] = true
					index++
				} else {
					if checkWinner(storedDigit) {
						c.Matches++
					}
				}
				storedDigit = 0
			} else if character == 124 {
				// vertical bar
				checkingWinners = true
			}
		}
	}
	if checkWinner(storedDigit) {
		c.Matches++
	}
	for i := 0; i < index; i++ {
		candidatesMap[winningNumbers[i]] = false
	}
}
