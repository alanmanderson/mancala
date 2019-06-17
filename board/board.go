package board

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidMove returned when an invalid move is requested
var ErrInvalidMove = errors.New("Invalid Move")

// BoardSize is the number of pits in the board
const BoardSize = 14

// Board struct for handling a Board
type Board struct {
	pits      [BoardSize]int
	p1Pit     int
	p2Pit     int
	curPlayer int
}

// Init initializes a board object
func (b *Board) Init() {
	b.p2Pit = 0
	b.p1Pit = BoardSize / 2
	for pos := range b.pits {
		if pos == b.p1Pit || pos == b.p2Pit {
			continue
		}
		b.pits[pos] = 4
	}
	b.curPlayer = 1
}

// Move executes a move from that pit
func (b *Board) Move(pit int) error {
	if !b.isValidPit(pit) {
		return ErrInvalidMove
	}
	beads := b.pits[pit]
	b.pits[pit] = 0
	curPit := pit + 1
	for beads > 0 {
		if b.curPlayer == 1 && curPit == b.p2Pit || b.curPlayer == 2 && curPit == b.p1Pit {
			curPit = (curPit + 1) % BoardSize
			continue
		}
		b.pits[curPit]++
		curPit = (curPit + 1) % BoardSize
		beads--
	}
	if b.curPlayer == 1 && curPit-1 == b.p1Pit {
		return nil
	}
	if b.curPlayer == 2 && curPit-1 == b.p2Pit {
		return nil
	}
	b.curPlayer = b.curPlayer%2 + 1
	if b.IsGameOver() {
		b.EndGame()
	}
	return nil
}

func (b Board) isValidPit(pit int) bool {
	if pit >= BoardSize || pit <= 0 || pit > 13 || b.pits[pit] == 0 || pit == b.p1Pit {
		return false
	}
	if b.curPlayer == 1 && pit > 6 {
		return false
	}
	if b.curPlayer == 2 && pit <= 7 {
		return false
	}
	return true
}

// Print prints the board
func (b Board) Print() string {
	/*
	   -------------------------------------------
	   | |0 | (13) (12) (11) (10) (9 ) (8 ) |7 | |
	   | |__| (1 ) (2 ) (3 ) (4 ) (5 ) (6 ) |__| |
	   -------------------------------------------
	*/
	var sb strings.Builder
	sb.WriteString(strings.Repeat("-", 43) + "\n")
	sb.WriteString(fmt.Sprintf("| |%2d| (%2d) (%2d) (%2d) (%2d) (%2d) (%2d) |%2d| |\n",
		b.pits[0], b.pits[13], b.pits[12], b.pits[11], b.pits[10], b.pits[9], b.pits[8],
		b.pits[7]),
	)
	sb.WriteString(fmt.Sprintf("| |__| (%2d) (%2d) (%2d) (%2d) (%2d) (%2d) |__| |\n",
		b.pits[1], b.pits[2], b.pits[3], b.pits[4], b.pits[5], b.pits[6]),
	)
	sb.WriteString(strings.Repeat("-", 43) + "\n")
	return sb.String()
}

func (b Board) isEqual(cmp Board) bool {
	return b.pits == cmp.pits && b.curPlayer == cmp.curPlayer
}

//EndGame tells if the game is over
func (b *Board) EndGame() {
	for i := 0; i < BoardSize; i++ {
		if i == b.p1Pit || i == b.p2Pit {
			continue
		}
		if i < BoardSize/2 {
			b.pits[b.p1Pit] += b.pits[i]
		} else if i > BoardSize/2 {
			b.pits[b.p2Pit] += b.pits[i]
		}
		b.pits[i] = 0
	}
}

// IsGameOver checks if the game is over
func (b Board) IsGameOver() bool {
	isOver := true
	for i := 1; i < 7; i++ {
		if b.pits[i] > 0 {
			isOver = false
			break
		}
	}
	if isOver {
		return true
	}
	for i := 8; i < 14; i++ {
		if b.pits[i] > 0 {
			return false
		}
	}
	return true
}
