package board

var isGameOverTestCases = []struct {
	description string
	board       Board
	expected    bool
}{
	{
		"Test player 2 ends the game",
		Board{
			[14]int{0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0},
			7,
			0,
			1,
		},
		true,
	},
	{
		"Test player 1 ends the game",
		Board{
			[14]int{15, 0, 0, 0, 0, 0, 0, 15, 18, 0, 0, 0, 0, 0},
			7,
			0,
			2,
		},
		true,
	},
	{
		"Test game is not over",
		Board{
			[14]int{15, 9, 0, 0, 0, 0, 0, 15, 9, 0, 0, 0, 0, 0},
			7,
			0,
			2,
		},
		false,
	},
}

var endGameTestCases = []struct {
	description string
	board       Board
	expected    Board
}{
	{
		"Test a game ends for P2",
		Board{
			[14]int{10, 4, 4, 4, 4, 4, 4, 2, 0, 0, 0, 0, 0, 0},
			7,
			0,
			1,
		},
		Board{
			[14]int{10, 0, 0, 0, 0, 0, 0, 26, 0, 0, 0, 0, 0, 0},
			7,
			0,
			1,
		},
	},
	{
		"Test a game ends for P1",
		Board{
			[14]int{2, 0, 0, 0, 0, 0, 0, 10, 4, 4, 4, 4, 4, 4},
			7,
			0,
			1,
		},
		Board{
			[14]int{26, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0},
			7,
			0,
			1,
		},
	},
}
