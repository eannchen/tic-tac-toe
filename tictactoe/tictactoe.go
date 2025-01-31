package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DefaultSize = 3
	PlayerOne   = player{
		id:     1,
		name:   "Player 1",
		symbol: "O",
	}
	PlayerTwo = player{
		id:     2,
		name:   "Player 2",
		symbol: "X",
	}
	PlayerNone = player{}
)

type TicTacToe struct {
	grid       map[[2]int]int
	oneDimSize int
	curTurn    player
	curRecord  currentRecord
}

type currentRecord struct {
	rowCounts  []int
	colCounts  []int
	diagCounts [2]int
}

type option struct {
	oneDimSize int
}

type player struct {
	id     int
	name   string
	symbol string
}

func NewTicTacToe() TicTacToe {
	return TicTacToe{
		grid:       make(map[[2]int]int),
		oneDimSize: DefaultSize,
		curTurn:    PlayerOne,
		curRecord: currentRecord{
			rowCounts: make([]int, DefaultSize),
			colCounts: make([]int, DefaultSize),
		},
	}
}

func (t *TicTacToe) Start() error {
	opt, err := t.readOptions()
	if err != nil {
		return errors.New("[readOptions] error: " + err.Error())
	}
	t.setOptions(opt)

	for {
		pos, err := t.readUserMove()
		if err != nil {
			return errors.New("[readUserMove] error: " + err.Error())
		}
		if winner, isEnd := t.setUserMove(pos[0], pos[1]); isEnd {
			t.printGrid()
			if winner == PlayerNone {
				fmt.Println("Draw!")
			} else {
				fmt.Printf("The winner is %s!\n", winner.name)
			}
			break
		}
		t.printGrid()
		t.takeTurn()
	}
	return nil
}

func (t TicTacToe) readOptions() (option, error) {
	fmt.Print("Enter the size for TicTacToe game (default: 3): ")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return option{}, err
		}
		size, err := t.validateSize(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return option{
			oneDimSize: size,
		}, nil
	}
}

func (t *TicTacToe) setOptions(opt option) {
	t.oneDimSize = opt.oneDimSize
	t.curRecord = currentRecord{
		rowCounts: make([]int, t.oneDimSize),
		colCounts: make([]int, t.oneDimSize),
	}
}

func (t TicTacToe) readUserMove() ([2]int, error) {
	fmt.Printf("Enter the %s position (0-index) in the format 'row,column': ", t.curTurn.name)

	var pos [2]int
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return [2]int{}, err
		}
		pos, err = t.validatePosition(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}

	return pos, nil
}

func (t *TicTacToe) setUserMove(r, c int) (player, bool) {

	if t.curTurn == PlayerOne {
		t.grid[[2]int{r, c}] = PlayerOne.id

		t.curRecord.rowCounts[r]++
		t.curRecord.colCounts[c]++
		if r+c == t.oneDimSize-1 {
			t.curRecord.diagCounts[0]++
		}
		if r == c {
			t.curRecord.diagCounts[1]++
		}

		// user 1 win
		if t.curRecord.rowCounts[r] == t.oneDimSize ||
			t.curRecord.colCounts[c] == t.oneDimSize ||
			t.curRecord.diagCounts[0] == t.oneDimSize ||
			t.curRecord.diagCounts[1] == t.oneDimSize {
			return t.curTurn, true
		}
	} else {
		t.grid[[2]int{r, c}] = PlayerTwo.id

		t.curRecord.rowCounts[r]--
		t.curRecord.colCounts[c]--
		if r+c == t.oneDimSize-1 {
			t.curRecord.diagCounts[0]--
		}
		if r == c {
			t.curRecord.diagCounts[1]--
		}

		// user 2 win
		if t.curRecord.rowCounts[r] == t.oneDimSize*-1 ||
			t.curRecord.colCounts[c] == t.oneDimSize*-1 ||
			t.curRecord.diagCounts[0] == t.oneDimSize*-1 ||
			t.curRecord.diagCounts[1] == t.oneDimSize*-1 {
			return t.curTurn, true
		}
	}

	// tie
	if len(t.grid) == t.oneDimSize*t.oneDimSize {
		return PlayerNone, true
	}

	return PlayerNone, false
}

func (t *TicTacToe) takeTurn() {
	if t.curTurn == PlayerOne {
		t.curTurn = PlayerTwo
	} else {
		t.curTurn = PlayerOne
	}
}

func (t TicTacToe) printGrid() {
	for r := 0; r < t.oneDimSize; r++ {
		var colStr string
		for c := 0; c < t.oneDimSize; c++ {
			if playerID, ok := t.grid[[2]int{r, c}]; ok {
				if playerID == PlayerOne.id {
					colStr += PlayerOne.symbol + " "
				} else {
					colStr += PlayerTwo.symbol + " "
				}
			} else {
				colStr += "- "
			}
		}
		fmt.Printf("%s\n", colStr)
	}
}

func (t TicTacToe) validateSize(input string) (int, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return DefaultSize, nil
	}
	size, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Invalid input. Please enter number for size.")
	}
	if size < 3 {
		return -1, errors.New("Invalid input. The size needs to be >= 3.")
	}
	return size, nil
}

func (t TicTacToe) validatePosition(input string) ([2]int, error) {
	parts := strings.Split(strings.TrimSpace(input), ",")
	if len(parts) != 2 {
		return [2]int{}, errors.New("Invalid input. Please enter in the format 'row,column'.")
	}
	r, err := strconv.Atoi(parts[0])
	if err != nil {
		return [2]int{}, errors.New("Invalid input. Please enter number for position.")
	}
	c, err := strconv.Atoi(parts[1])
	if err != nil {
		return [2]int{}, errors.New("Invalid input. Please enter number for position.")
	}
	if c < 0 || r < 0 {
		return [2]int{}, errors.New("Invalid input. Please enter in the number > -1.")
	}
	if c >= t.oneDimSize || r >= t.oneDimSize {
		return [2]int{}, errors.New("Invalid input. Please enter the number within the range.")
	}
	if _, isTaken := t.grid[[2]int{r, c}]; isTaken {
		return [2]int{}, errors.New("Invalid input. The position is taken.")
	}
	return [2]int{r, c}, nil
}
