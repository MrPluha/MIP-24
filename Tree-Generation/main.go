package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxStart = 50001
	minStart = 40000
)

type number struct {
	value       int
	points      int
	bank        int
	totalPoints int
	// 0 no winner yet, 1 win start player, 2 win second player
	winPlayer int
}

type gameState struct {
	// prevState  *gameState
	final      bool
	nextStates []*gameState
	current    []*number
}

func main() {

	startState := prepareGameStart(getStartValues(os.Args[1:]))

	fmt.Println("GAME START VALUES")
	for _, v := range startState.current {
		fmt.Printf("%d\n", v.value)
	}

	startState.findNextState()

	startState.printEndpoints()
}

func getStartValues(args []string) []int {
	if len(args) > 0 {
		var inputStartValues []int
		for _, v := range args {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("could not change input argument %s to integer, err: %s", v, err.Error()))
			}

			inputStartValues = append(inputStartValues, i)
		}
		return inputStartValues
	}

	return generateStartNr(time.Now().UnixNano())
}

func (gs *gameState) printEndpoints() {
	if gs.final {
		gs.calculateGameEnd()
		for _, nr := range gs.current {
			fmt.Println()
			fmt.Println()
			fmt.Printf("value %d\n", nr.value)
			fmt.Printf("point %d\n", nr.points)
			fmt.Printf("bank %d\n", nr.bank)
			fmt.Printf("total points %d\n", nr.totalPoints)
			fmt.Printf("Player %d would win, 1 = player who took the first turn, 2= 2nd player\n", nr.winPlayer)
			fmt.Println()
			fmt.Println()
		}
	}

	for _, next := range gs.nextStates {
		next.printEndpoints()
	}
}

func (gs *gameState) calculateGameEnd() {
	for _, nr := range gs.current {
		if nr.points%2 == 0 {
			nr.totalPoints = nr.points - nr.bank

			if nr.totalPoints%2 == 0 {
				nr.winPlayer = 1
				continue
			}

			nr.winPlayer = 2

			continue
		}

		nr.totalPoints = nr.points + nr.bank
		if nr.totalPoints%2 == 0 {
			nr.winPlayer = 1
			continue
		}

		nr.winPlayer = 2
	}
}

func prepareGameStart(startNumbers []int) *gameState {
	var nrs []*number

	for _, value := range startNumbers {
		nrs = append(nrs, newNumber(value, 0, 0))
	}

	return newGameState(nrs, nil)
}

func (gs *gameState) findNextState() {
	for _, currentNumber := range gs.current {
		if currentNumber == nil {
			continue
		}

		var nextNumbers []*number

		nextValues := currentNumber.calculateNextNumbers()
		for _, v := range nextValues {
			new := newNumber(v, currentNumber.points, currentNumber.bank)
			new.calculatePointsAndBank()
			nextNumbers = append(nextNumbers, new)
		}

		if len(nextNumbers) == 0 {
			continue
		}

		nextGs := newGameState(nextNumbers, gs)

		gs.nextStates = append(gs.nextStates, nextGs)

		nextGs.findNextState()
	}

	if gs.nextStates == nil {
		gs.final = true
	}
}

func newGameState(nrs []*number, parent *gameState) *gameState {
	return &gameState{
		current: nrs,
		// prevState: parent,
	}
}

func newNumber(v int, points, bank int) *number {
	if v < 1 {
		return nil
	}

	return &number{
		value:  v,
		points: points,
		bank:   bank,
	}
}

func generateStartNr(seed int64) []int {
	var intInRange []int
	out := make([]int, 5)

	for i := minStart; i < maxStart; i++ {
		if i%3 == 0 && i%4 == 0 && i%5 == 0 {
			intInRange = append(intInRange, i)
		}
	}

	r := rand.New(rand.NewSource(seed))

	for j := 0; j < 5; j++ {
		out[j] = intInRange[r.Intn(len(intInRange))]
	}

	return out
}

func (nr *number) calculatePointsAndBank() {
	x := nr.value

	if x%2 == 0 {
		nr.points++
	} else {
		nr.points--
	}

	if x%10 == 0 {
		nr.bank++
	} else if x%5 == 0 {
		nr.bank++
	}
}

func (nr *number) calculateNextNumbers() []int {
	var out []int

	x := nr.value

	if x%3 == 0 {
		out = append(out, x/3)
	}

	if x%4 == 0 {
		out = append(out, x/4)
	}

	if x%5 == 0 {
		out = append(out, x/5)
	}

	return out
}
