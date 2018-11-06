package aoc17

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePuzzleOne(t *testing.T) {
	shortestDistanceCalcFunc := func(cellNumber int) int {
		x, y := getCartesianCoords(cellNumber)

		return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
	}
	//Data from square 1 is carried 0 steps, since it's at the access port.
	//Data from square 12 is carried 3 steps, such as: down, left, left.
	//Data from square 23 is carried only 2 steps: up twice.
	//Data from square 1024 must be carried 31 steps.
	assert.Equal(t, 0, shortestDistanceCalcFunc(1))
	assert.Equal(t, 1, shortestDistanceCalcFunc(2))
	assert.Equal(t, 2, shortestDistanceCalcFunc(3))
	assert.Equal(t, 1, shortestDistanceCalcFunc(4))
	assert.Equal(t, 2, shortestDistanceCalcFunc(5))
	assert.Equal(t, 1, shortestDistanceCalcFunc(6))
	assert.Equal(t, 2, shortestDistanceCalcFunc(7))
	assert.Equal(t, 1, shortestDistanceCalcFunc(8))
	assert.Equal(t, 2, shortestDistanceCalcFunc(9))
	assert.Equal(t, 3, shortestDistanceCalcFunc(10))
	assert.Equal(t, 3, shortestDistanceCalcFunc(12))
	assert.Equal(t, 2, shortestDistanceCalcFunc(23))
	assert.Equal(t, 31, shortestDistanceCalcFunc(1024))
	assert.Equal(t, 552, shortestDistanceCalcFunc(325489))

}

func TestDayThreePuzzleTwo(t *testing.T) {
	//Square 1 starts with the value 1.
	//Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
	//Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
	//Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
	//Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.

	adjacentSquareSummer := func(cellNumber int) int {
		board := &Board{}
		board.set(0, 0, 1)
		nextCellX := 1
		nextCellY := 0
		x := 0
		y := 0
		direction := 'r'
		for i := 2; i <= cellNumber; i++ {
			x = nextCellX
			y = nextCellY
			currentCellValue := neighbourValues(x, y, board)
			board.set(x, y, currentCellValue)
			lnx, lny := leftNeighbour(x, y, direction)
			if board.cell(lnx, lny) == 0 {
				nextCellX = lnx
				nextCellY = lny

				switch direction {
				case 'r':
					direction = 'u'
				case 'l':
					direction = 'd'
				case 'u':
					direction = 'l'
				case 'd':
					direction = 'r'
				}
			} else {
				nextCellX, nextCellY = forwardNeighbour(x, y, direction)
			}

		}
		return board.cell(x, y)

	}

	assert.Equal(t, 1, adjacentSquareSummer(1))
	assert.Equal(t, 1, adjacentSquareSummer(2))
	assert.Equal(t, 2, adjacentSquareSummer(3))
	assert.Equal(t, 4, adjacentSquareSummer(4))
	assert.Equal(t, 5, adjacentSquareSummer(5))

	findAbove := func(minimum int) int {
		for i := 1; ; i++ {
			if adjacentSquareSummer(i) >= minimum {
				return adjacentSquareSummer(i)
			}
		}
	}
	assert.Equal(t, 1, findAbove(1))
	assert.Equal(t, 2, findAbove(2))
	assert.Equal(t, 4, findAbove(3))
	assert.Equal(t, 4, findAbove(4))
	assert.Equal(t, 5, findAbove(5))
	assert.Equal(t, 330785, findAbove(325489))
}

func TestCartesianCoords(t *testing.T) {
	x, y := getCartesianCoords(1)
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)

	x, y = getCartesianCoords(2)
	assert.Equal(t, 1, x)
	assert.Equal(t, 0, y)

	x, y = getCartesianCoords(3)
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)

	x, y = getCartesianCoords(4)
	assert.Equal(t, 0, x)
	assert.Equal(t, 1, y)

	x, y = getCartesianCoords(5)
	assert.Equal(t, -1, x)
	assert.Equal(t, 1, y)

	x, y = getCartesianCoords(6)
	assert.Equal(t, -1, x)
	assert.Equal(t, 0, y)

	x, y = getCartesianCoords(7)
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	x, y = getCartesianCoords(8)
	assert.Equal(t, 0, x)
	assert.Equal(t, -1, y)

	x, y = getCartesianCoords(9)
	assert.Equal(t, 1, x)
	assert.Equal(t, -1, y)

	x, y = getCartesianCoords(10)
	assert.Equal(t, 2, x)
	assert.Equal(t, -1, y)
}

type Board struct {
	board map[int]map[int]int
}

func (b *Board) cell(x, y int) int {
	if b.board == nil {
		return 0
	}

	if b.board[x] == nil {
		return 0
	}

	return b.board[x][y]
}

func (b *Board) set(x, y, cellValue int) {
	if b.board == nil {
		b.board = make(map[int]map[int]int)
	}

	if b.board[x] == nil {
		b.board[x] = make(map[int]int)
	}

	b.board[x][y] = cellValue

}

func neighbourValues(x, y int, b *Board) int {
	return b.cell(x+1, y) +
		b.cell(x+1, y+1) +
		b.cell(x, y+1) +
		b.cell(x-1, y+1) +
		b.cell(x-1, y) +
		b.cell(x-1, y-1) +
		b.cell(x, y-1) +
		b.cell(x+1, y-1)
}
func getCartesianCoords(cellNumber int) (int, int) {
	board := &Board{}

	if cellNumber == 1 {
		return 0, 0
	}
	x := 1
	y := 0
	board.set(0, 0, 1)
	board.set(1, 0, 2)
	direction := 'r'
	for i := 3; i <= cellNumber; i++ {
		leftNeightbourX, leftNeighbourY := leftNeighbour(x, y, direction)
		if board.cell(leftNeightbourX, leftNeighbourY) == 0 {
			board.set(leftNeightbourX, leftNeighbourY, cellNumber)
			x = leftNeightbourX
			y = leftNeighbourY

			switch direction {
			case 'r':
				direction = 'u'
			case 'l':
				direction = 'd'
			case 'u':
				direction = 'l'
			case 'd':
				direction = 'r'
			}
		} else {
			forwardNeighbourX, forwardNeighbourY := forwardNeighbour(x, y, direction)
			board.set(forwardNeighbourX, forwardNeighbourY, cellNumber)
			x = forwardNeighbourX
			y = forwardNeighbourY
		}
	}
	return x, y
}

func forwardNeighbour(x, y int, direction rune) (int, int) {
	switch direction {
	case 'r':
		return x + 1, y
	case 'l':
		return x - 1, y
	case 'u':
		return x, y + 1
	case 'd':
		return x, y - 1
	}
	return 0, 0
}

func leftNeighbour(x, y int, direction rune) (int, int) {
	switch direction {
	case 'r':
		return x, y + 1
	case 'l':
		return x, y - 1
	case 'u':
		return x - 1, y
	case 'd':
		return x + 1, y
	}
	return 0, 0
}
