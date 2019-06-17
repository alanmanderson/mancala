package board

var moveTestCases = []struct {
	description string
	board       Board
	move        int
	expected    Board
	err         error
}{
	{
		"Test a valid move that ends in the P1s pit",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		3,
		Board{
			[14]int{0, 4, 4, 0, 5, 5, 5, 1, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		nil,
	},
	{
		"Test a valid move that ends in the P1s pit",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			2,
		},
		10,
		Board{
			[14]int{1, 4, 4, 4, 4, 4, 4, 0, 4, 4, 0, 5, 5, 5},
			7,
			0,
			2,
		},
		nil,
	},
	{
		"Test an invalid move from an empty pit",
		Board{
			[14]int{0, 4, 0, 5, 5, 5, 5, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		2,
		Board{
			[14]int{0, 4, 0, 5, 5, 5, 5, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		ErrInvalidMove,
	},
	{
		"Test an invalid move from a number off the board",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		14,
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		ErrInvalidMove,
	},
	{
		"Test an invalid move from the wrong player P1",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		12,
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		ErrInvalidMove,
	},
	{
		"Test an invalid move from the wrong player P2",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			2,
		},
		2,
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			2,
		},
		ErrInvalidMove,
	},
	{
		"Test an invalid move from the scoring Pit",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		0,
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		ErrInvalidMove,
	},
	{
		"Test a standard move",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		2,
		Board{
			[14]int{0, 4, 0, 5, 5, 5, 5, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			2,
		},
		nil,
	},
	{
		"Test a player 2 move that wraps around the board",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
			7,
			0,
			2,
		},
		12,
		Board{
			[14]int{1, 5, 5, 4, 4, 4, 4, 0, 4, 4, 4, 4, 0, 5},
			7,
			0,
			1,
		},
		nil,
	},
	{
		"Test a player 2 move that wraps all the way around the board",
		Board{
			[14]int{10, 0, 0, 0, 0, 0, 1, 10, 24, 0, 0, 0, 0, 0},
			7,
			0,
			2,
		},
		8,
		Board{
			[14]int{12, 2, 2, 2, 2, 2, 2, 10, 1, 2, 2, 2, 2, 2},
			7,
			0,
			1,
		},
		nil,
	},
	{
		"Test a player 2 move that ends the game",
		Board{
			[14]int{10, 0, 0, 0, 0, 0, 5, 20, 0, 0, 0, 0, 0, 1},
			7,
			0,
			2,
		},
		13,
		Board{
			[14]int{11, 0, 0, 0, 0, 0, 0, 25, 0, 0, 0, 0, 0, 0},
			7,
			0,
			2,
		},
		nil,
	},
}
