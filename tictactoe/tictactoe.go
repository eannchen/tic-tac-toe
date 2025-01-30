package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TicTacToe struct {
	grid       map[[2]int]int
	oneDimSize int
	curTurn    int // 0 (none) / 1 (user 1) / 2 (user 2)
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

func NewTicTacToe() TicTacToe {
	return TicTacToe{}
}

func (t TicTacToe) Start() error {
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
			if winner == 0 {
				fmt.Println("Draw!")
			} else {
				fmt.Printf("The winner is User %d!\n", winner)
			}
			break
		}
		t.printGrid()
		t.takeTurn()
	}
	return nil
}

func (t TicTacToe) readOptions() (option, error) {
	fmt.Print("Enter the size (number) for TicTacToe game: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return option{}, err
		}
		size, err := t.praseSize(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if size < 3 {
			fmt.Println("Invalid input. The size needs to be >= 3.")
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
	t.grid = make(map[[2]int]int)
	t.curTurn = 1
}

func (t TicTacToe) readUserMove() ([2]int, error) {
	fmt.Printf("Enter the User %d position (0-index) in the format 'row,column': ", t.curTurn)

	var pos [2]int
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return [2]int{}, err
		}
		pos, err = t.prasePosition(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if _, isTaken := t.grid[[2]int{pos[0], pos[1]}]; isTaken {
			fmt.Println("Invalid input. The position is taken.")
			continue
		}
		break
	}

	return pos, nil
}

func (t *TicTacToe) setUserMove(r, c int) (int, bool) {

	if t.curTurn == 1 {
		t.grid[[2]int{r, c}] = 1

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
		t.grid[[2]int{r, c}] = 2

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
		return 0, true
	}

	return 0, false
}

func (t *TicTacToe) takeTurn() {
	t.curTurn = t.curTurn%2 + 1
}

func (t TicTacToe) printGrid() {
	for r := 0; r < t.oneDimSize; r++ {
		var colStr string
		for c := 0; c < t.oneDimSize; c++ {
			if val, ok := t.grid[[2]int{r, c}]; ok {
				if val == 1 {
					colStr += "O "
				} else {
					colStr += "X "
				}
			} else {
				colStr += "  "
			}
		}
		fmt.Printf("%s\n", colStr)
	}
}

func (t TicTacToe) praseSize(input string) (int, error) {
	size, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1, errors.New("Invalid input. Please enter number for size.")
	}
	return size, nil
}

func (t TicTacToe) prasePosition(input string) ([2]int, error) {
	parts := strings.Split(strings.TrimSpace(input), ",")
	if len(parts) != 2 {
		return [2]int{}, errors.New("Invalid input. Please enter in the format 'row,column'.")
	}
	c, err := strconv.Atoi(parts[0])
	if err != nil {
		return [2]int{}, errors.New("Invalid input. Please enter number for position.")
	}
	r, err := strconv.Atoi(parts[1])
	if err != nil {
		return [2]int{}, errors.New("Invalid input. Please enter number for position.")
	}
	if c < 0 || r < 0 {
		return [2]int{}, errors.New("Invalid input. Please enter in the number > -1.")
	}
	if c >= t.oneDimSize || r >= t.oneDimSize {
		return [2]int{}, errors.New("Invalid input. Please enter the number within the range.")
	}
	return [2]int{c, r}, nil
}
