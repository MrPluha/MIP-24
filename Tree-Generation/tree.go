package tree

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

type Number struct {
	Value       int
	Points      int
	Bank        int
	TotalPoints int
	// 0 no winner yet, 1 win start player (min), -1 win second player(max)
	WinPlayer int
}

type GameState struct {
	// prevState  *gameState
	Final      bool
	Current    []*Number
	NextStates []*GameState
}

func Test() {

	startState := PrepareGameStart(GetStartValues(os.Args[1:]))

	fmt.Println("GAME START VALUES")
	for _, v := range startState.Current {
		fmt.Printf("%d\n", v.Value)
	}

	startState.FindNextState()

	startState.PrintEndpoints()
}

func GetStartValues(args []string) []int {
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

func (gs *GameState) PrintEndpoints() {
	if gs.Final {
		gs.CalculateGameEnd()
		for _, nr := range gs.Current {
			fmt.Println()
			fmt.Println()
			fmt.Printf("Value %d\n", nr.Value)
			fmt.Printf("point %d\n", nr.Points)
			fmt.Printf("Bank %d\n", nr.Bank)
			fmt.Printf("total Points %d\n", nr.TotalPoints)
			fmt.Printf("Player %d would win, 1 = player who took the first turn, -1 = 2nd player\n", nr.WinPlayer)
			fmt.Println()
			fmt.Println()
		}
	}

	for _, next := range gs.NextStates {
		next.PrintEndpoints()
	}
}

func (gs *GameState) CalculateGameEnd() {
	for _, nr := range gs.Current {
		if nr.Points%2 == 0 {
			nr.TotalPoints = nr.Points - nr.Bank

			if nr.TotalPoints%2 == 0 {
				nr.WinPlayer = 1
				continue
			}

			nr.WinPlayer = -1

			continue
		}

		nr.TotalPoints = nr.Points + nr.Bank
		if nr.TotalPoints%2 == 0 {
			nr.WinPlayer = 1
			continue
		}

		nr.WinPlayer = -1
	}
}

func PrepareGameStart(startNumbers []int) *GameState {
	var nrs []*Number

	for _, Value := range startNumbers {
		nrs = append(nrs, newNumber(Value, 0, 0))
	}

	return newGameState(nrs, nil)
}

func (gs *GameState) FindNextState() {
	for _, currentNumber := range gs.Current {
		if currentNumber == nil {
			continue
		}

		var nextNumbers []*Number

		nextValues := currentNumber.calculateNextNumbers()
		for _, v := range nextValues {
			new := newNumber(v, currentNumber.Points, currentNumber.Bank)
			new.calculatePointsAndBank()
			nextNumbers = append(nextNumbers, new)
		}

		if len(nextNumbers) == 0 {
			continue
		}

		nextGs := newGameState(nextNumbers, gs)

		gs.NextStates = append(gs.NextStates, nextGs)

		nextGs.FindNextState()
	}

	if gs.NextStates == nil {
		gs.Final = true
	}
}

func newGameState(nrs []*Number, parent *GameState) *GameState {
	return &GameState{
		Current: nrs,
		// prevState: parent,
	}
}

func newNumber(v int, Points, Bank int) *Number {
	if v < 1 {
		return nil
	}

	return &Number{
		Value:  v,
		Points: Points,
		Bank:   Bank,
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

func (nr *Number) calculatePointsAndBank() {
	x := nr.Value

	if x%2 == 0 {
		nr.Points++
	} else {
		nr.Points--
	}

	if x%10 == 0 {
		nr.Bank++
	} else if x%5 == 0 {
		nr.Bank++
	}
}

func (nr *Number) calculateNextNumbers() []int {
	var out []int

	x := nr.Value

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
