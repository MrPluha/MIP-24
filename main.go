package main

import (
	"fmt"
	tree "game/Tree-Generation"
	"os"
)

func main() {

	runGame()
}

func runGame() {
	startNumbers := tree.GetStartValues(os.Args[1:])
	fmt.Println("StrNR", startNumbers)
	startNode := tree.PrepareGameStart(startNumbers)
	startNode.FindNextState()

	value := minMax(startNode, true)
	fmt.Println("value", value)
}

func minMaxTest(startNumbers []int) int {
	startNode := tree.PrepareGameStart(startNumbers)
	startNode.FindNextState()
	return minMax(startNode, true)

}

// if max == false, it means it's min player
func minMax(state *tree.GameState, max bool) int {
	if state.Final {
		if len(state.Current) != 1 {
			panic("current numbers must be 1")
		}

		state.CalculateGameEnd()

		return state.Current[0].WinPlayer
	}

	if max {
		//Max state

		v := calculatePoints(state.NextStates, false)

		return v

	}

	//Min state
	v := calculatePoints(state.NextStates, true)

	return v
}

func calculatePoints(nextStates []*tree.GameState, max bool) int {
	var values []int

	for _, child := range nextStates {
		values = append(values, minMax(child, max))
	}

	points := getValue(values, max)
	if points == 0 {
		panic("point should not be 0")
	}

	return points
}

func getValue(in []int, max bool) int {
	var hasMax bool
	var hasMin bool

	if len(in) < 1 {
		return 0
	}

	for _, n := range in {
		if n == 1 {
			hasMax = true
		}

		if n == -1 {
			hasMin = true
		}
	}

	if max && hasMax {
		return 1
	}

	if max && !hasMax {
		return -1
	}

	if !max && hasMin {
		return -1
	}

	if !max && !hasMin {
		return 1
	}

	// should never happen
	return 0
}
