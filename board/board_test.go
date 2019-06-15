package board

import (
	"testing"
)

func TestInit(t *testing.T) {
	b := Board{}
	b.Init()
	expects := Board{
		[14]int{0, 4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4},
		0,
		7,
		1,
	}
	if !b.isEqual(expects) {
		t.Fatal("FAIL: Init failed to create a correct board")
	}
	t.Log("PASS: Correctly created the beginning board")
}

func TestMove(t *testing.T) {
	for _, testCase := range moveTestCases {
		err := testCase.board.Move(testCase.move)
		if err != testCase.err {
			t.Fatalf("FAIL: %s(%v)\nExpected: %q\nActual: %q",
				testCase.description, testCase.board, testCase.err, err)
		}
		if !testCase.board.isEqual(testCase.expected) {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v",
				testCase.description, testCase.board, testCase.expected)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestIsGameOver(t *testing.T) {
	for _, testCase := range isGameOverTestCases {
		actual := testCase.board.IsGameOver()
		if actual != testCase.expected {
			t.Fatalf("FAIL: %s(%v)\nExpected: %t\nActual: %t",
				testCase.description, testCase.board, testCase.expected, actual)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

// func BenchmarkMove(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range testCases {
// 			continue
// 		}
// 	}
// }
