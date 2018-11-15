package aoc17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaySixPuzzleOne(t *testing.T) {
	//The banks start with 0, 2, 7, and 0 blocks. The third bank has the most blocks, so it is chosen for redistribution.
	//Starting with the next bank (the fourth bank) and then continuing to the first bank, the second bank, and so on, the 7 blocks are spread out over the memory banks. The fourth, first, and second banks get two blocks each, and the third bank gets one back. The final result looks like this: 2 4 1 2.
	//Next, the second bank is chosen because it contains the most blocks (four). Because there are four memory banks, each gets one block. The result is: 3 1 2 3.
	//Now, there is a tie between the first and fourth memory banks, both of which have three blocks. The first bank wins the tie, and its three blocks are distributed evenly over the other three banks, leaving it with none: 0 2 3 4.
	//The fourth bank is chosen, and its four blocks are distributed such that each of the four banks receives one: 1 3 4 1.
	//The third bank is chosen, and the same thing happens: 2 4 1 2.
	//At this point, we've reached a state we've seen before: 2 4 1 2 was already seen. The infinite loop is detected after the fifth block redistribution cycle, and so the answer in this example is 5
	loopDetector := func(input []int) int {
		states := make(map[string]bool)
		initialState := fmt.Sprintf("%+v", input)
		states[initialState] = true
		currentState := rebalance(input)
		lastState := fmt.Sprintf("%+v", currentState)
		loopCount := 1
		for !states[lastState] {
			states[lastState] = true
			currentState = rebalance(currentState)
			lastState = fmt.Sprintf("%+v", currentState)
			loopCount++
		}
		return loopCount
	}

	assert.Equal(t, 5, loopDetector([]int{0, 2, 7, 0}))
	assert.Equal(t, 7864, loopDetector([]int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}))

}

func rebalance(input []int) []int {
	newArray := make([]int, len(input))

	highestIndex := 0
	assignmentLeft := input[highestIndex]
	for index, elem := range input {
		if elem > input[highestIndex] {
			highestIndex = index
			assignmentLeft = elem
		}
		newArray[index] = elem
	}

	newArray[highestIndex] = 0

	nextIndex := highestIndex + 1
	for assignmentLeft > 0 {
		if nextIndex >= len(newArray) {
			nextIndex = 0
		}
		newArray[nextIndex] = newArray[nextIndex] + 1
		assignmentLeft--
		nextIndex++
	}

	return newArray
}

func TestRebalance(t *testing.T) {
	assert.Equal(t, []int{2, 4, 1, 2}, rebalance([]int{0, 2, 7, 0}))
	assert.Equal(t, []int{3, 1, 2, 3}, rebalance([]int{2, 4, 1, 2}))
	assert.Equal(t, []int{0, 2, 3, 4}, rebalance([]int{3, 1, 2, 3}))
	assert.Equal(t, []int{1, 3, 4, 1}, rebalance([]int{0, 2, 3, 4}))
	assert.Equal(t, []int{2, 4, 1, 2}, rebalance([]int{1, 3, 4, 1}))
}

func TestDaySixPuzzleTwo(t *testing.T) {

	loopSizeCalc := func(input []int) int {
		states := make(map[string]int)
		initialState := fmt.Sprintf("%+v", input)
		states[initialState] = 1
		currentState := rebalance(input)
		lastState := fmt.Sprintf("%+v", currentState)
		loopCount := 1
		cycleCount := 0
		for {
			currentState = rebalance(currentState)
			lastState = fmt.Sprintf("%+v", currentState)
			if states[lastState] > 0 {
				cycleCount = loopCount - states[lastState]
				break
			}
			states[lastState] = loopCount
			loopCount++

		}
		return cycleCount
	}

	assert.Equal(t, 4, loopSizeCalc([]int{0, 2, 7, 0}))
	assert.Equal(t, 1695, loopSizeCalc([]int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}))
}
