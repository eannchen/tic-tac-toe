package tictactoe

import "testing"

func TestSetUserMove(t *testing.T) {
	type condition struct {
		oneDimSize int
	}
	type useCase struct {
		name      string
		condition condition
		steps     [][2]int
		winner    player
		isEnd     bool
	}
	tests := []useCase{
		useCase{
			name: "Unfinished game",
			condition: condition{
				oneDimSize: 3,
			},
			// X O
			//   X
			//   O
			steps: [][2]int{
				[2]int{0, 1},
				[2]int{0, 0},
				[2]int{2, 1},
				[2]int{1, 1},
			},
			winner: PlayerNone,
			isEnd:  false,
		},
		useCase{
			name: "Draw",
			condition: condition{
				oneDimSize: 3,
			},
			// X O X
			// O X O
			// O X O
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{0, 1},
				[2]int{0, 2},
				[2]int{1, 1},
				[2]int{1, 0},
				[2]int{1, 2},
				[2]int{2, 1},
				[2]int{2, 0},
				[2]int{2, 2},
			},
			winner: PlayerNone,
			isEnd:  true,
		},
		useCase{
			name: "User 1 wins horizontally (row 0)",
			condition: condition{
				oneDimSize: 3,
			},
			// O O O
			// X X
			//
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{1, 0},
				[2]int{0, 1},
				[2]int{1, 1},
				[2]int{0, 2},
			},
			winner: PlayerOne,
			isEnd:  true,
		},
		useCase{
			name: "User 1 wins vertically (column 1)",
			condition: condition{
				oneDimSize: 3,
			},
			// X O
			// X O
			//   O
			steps: [][2]int{
				[2]int{0, 1},
				[2]int{0, 0},
				[2]int{1, 1},
				[2]int{1, 0},
				[2]int{2, 1},
			},
			winner: PlayerOne,
			isEnd:  true,
		},
		useCase{
			name: "User 1 Wins diagonally (from top left)",
			condition: condition{
				oneDimSize: 3,
			},
			// O   X
			// X O
			// O X O
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{0, 2},
				[2]int{1, 1},
				[2]int{1, 0},
				[2]int{2, 0},
				[2]int{2, 1},
				[2]int{2, 2},
			},
			winner: PlayerOne,
			isEnd:  true,
		},
		useCase{
			name: "User 1 wins diagonally (from top right)",
			condition: condition{
				oneDimSize: 3,
			},
			//   X O
			//   O X
			// O X O
			steps: [][2]int{
				[2]int{0, 2},
				[2]int{0, 1},
				[2]int{1, 1},
				[2]int{1, 2},
				[2]int{2, 2},
				[2]int{2, 1},
				[2]int{2, 0},
			},
			winner: PlayerOne,
			isEnd:  true,
		},
		useCase{
			name: "User 2 wins horizontally (row 0)",
			condition: condition{
				oneDimSize: 3,
			},
			// X X X
			// O O
			//   O
			steps: [][2]int{
				[2]int{1, 0},
				[2]int{0, 0},
				[2]int{1, 1},
				[2]int{0, 1},
				[2]int{2, 1},
				[2]int{0, 2},
			},
			winner: PlayerTwo,
			isEnd:  true,
		},
		useCase{
			name: "User 2 wins horizontally (row 1)",
			condition: condition{
				oneDimSize: 3,
			},
			//   O O
			// X X X
			//   O
			steps: [][2]int{
				[2]int{0, 1},
				[2]int{1, 0},
				[2]int{0, 2},
				[2]int{1, 1},
				[2]int{2, 1},
				[2]int{1, 2},
			},
			winner: PlayerTwo,
			isEnd:  true,
		},
		useCase{
			name: "User 2 wins horizontally (row 2)",
			condition: condition{
				oneDimSize: 3,
			},
			// O O
			//   O
			// X X X
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{2, 0},
				[2]int{0, 1},
				[2]int{2, 1},
				[2]int{1, 1},
				[2]int{2, 2},
			},
			winner: PlayerTwo,
			isEnd:  true,
		},
		useCase{
			name: "User 2 wins diagonally (from top left)",
			condition: condition{
				oneDimSize: 3,
			},
			// X O O
			// O X
			//     X
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{0, 2},
				[2]int{0, 1},
				[2]int{1, 1},
				[2]int{1, 0},
				[2]int{2, 0},
			},
			winner: PlayerTwo,
			isEnd:  true,
		},
		useCase{
			name: "User 2 wins diagonally (from top right)",
			condition: condition{
				oneDimSize: 3,
			},
			// O O X
			// O X
			// X
			steps: [][2]int{
				[2]int{0, 0},
				[2]int{0, 2},
				[2]int{0, 1},
				[2]int{1, 1},
				[2]int{1, 0},
				[2]int{2, 0},
			},
			winner: PlayerTwo,
			isEnd:  true,
		},
	}
	for _, tt := range tests {
		game := NewTicTacToe()
		game.setOptions(option{oneDimSize: tt.condition.oneDimSize})
		t.Run(tt.name, func(t *testing.T) {
			var winner player
			var isEnd bool
			for _, pos := range tt.steps {
				winner, isEnd = game.setUserMove(pos[0], pos[1])
				game.takeTurn()
			}
			game.printGrid()
			if isEnd != tt.isEnd {
				t.Errorf("setUserMove() returned %t, expected %t", isEnd, tt.isEnd)
			}
			if winner != tt.winner {
				t.Errorf("setUserMove() returned %v, expected %v", winner, tt.winner)
			}
		})
	}
}

func TestValidateSize(t *testing.T) {
	type useCase struct {
		name   string
		input  string
		size   int
		hasErr bool
	}
	tests := []useCase{
		useCase{
			name:   "Invalid number",
			input:  "!@#$%^&*",
			size:   -1,
			hasErr: true,
		},
		useCase{
			name:   "Invalid size",
			input:  "2",
			size:   -1,
			hasErr: true,
		},
		useCase{
			name:   "Ok, use default size",
			input:  "   ",
			size:   DefaultSize,
			hasErr: false,
		},
		useCase{
			name:   "Ok",
			input:  "5",
			size:   5,
			hasErr: false,
		},
	}
	for _, tt := range tests {
		game := NewTicTacToe()
		t.Run(tt.name, func(t *testing.T) {
			size, err := game.validateSize(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("validateSize() returned nil, error is expected")
			}
			if size != tt.size {
				t.Errorf("validateSize() returned %d, expected %d", size, tt.size)
			}
		})
	}
}

func TestValidatePosition(t *testing.T) {
	type condition struct {
		occupiedPosition [2]int
		oneDimSize       int
	}
	type useCase struct {
		name      string
		condition condition
		input     string
		position  [2]int
		hasErr    bool
	}

	tests := []useCase{
		useCase{
			name:     "Invalid format",
			input:    "!@#$%^&*",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name:     "Invalid format (row)",
			input:    "a,0",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name:     "Invalid format (column)",
			input:    "0,b",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name:     "Out of bound (row < 0)",
			input:    "-1,0",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name:     "Out of bound (column < 0)",
			input:    "0,-1",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name: "Out of bound (row >= range)",
			condition: condition{
				oneDimSize:       5,
				occupiedPosition: [2]int{1, 1},
			},
			input:    "99,0",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name: "Out of bound (column >= range)",
			condition: condition{
				oneDimSize:       5,
				occupiedPosition: [2]int{1, 1},
			},
			input:    "0,99",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name: "The position is taken",
			condition: condition{
				oneDimSize:       5,
				occupiedPosition: [2]int{1, 1},
			},
			input:    "0,99",
			position: [2]int{},
			hasErr:   true,
		},
		useCase{
			name: "Ok",
			condition: condition{
				oneDimSize:       5,
				occupiedPosition: [2]int{1, 1},
			},
			input:    "3,3",
			position: [2]int{3, 3},
			hasErr:   false,
		},
	}
	for _, tt := range tests {
		game := NewTicTacToe()
		game.setOptions(option{oneDimSize: tt.condition.oneDimSize})
		game.grid[tt.condition.occupiedPosition] = 1

		t.Run(tt.name, func(t *testing.T) {
			pos, err := game.validatePosition(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("validatePosition() returned nil, error is expected")
			}
			if pos != tt.position {
				t.Errorf("validatePosition() returned %v, expected %v", pos, tt.position)
			}
		})
	}
}
