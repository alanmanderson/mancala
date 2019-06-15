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
