package main

import (
	"fmt"
	"mancala/board"
	"mancala/interactive"
)

func main() {
	b := board.Board{}
	b.Init()
	interactive.Main()
	fmt.Println(b.Print())
	fmt.Println("%+v", b)
	b.Move(3)
	fmt.Println(b.Print())
	//fmt.Println("%+v", b)
	b.Move(4)
	fmt.Println(b.Print())
	b.Move(9)
	fmt.Println(b.Print())
	// fmt.Println("%+v", b)
	b.Move(6)
	fmt.Println(b.Print())
	b.Move(6)
	fmt.Println(b.Print())
	b.Move(5)
	fmt.Println(b.Print())
	b.Move(3)
	// fmt.Println("%+v", b)
	fmt.Println(b.Print())
}
