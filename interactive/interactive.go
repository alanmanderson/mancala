package interactive

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"strconv"
)

// Main does stuff
func Main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Sample Interactive Shell")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "play",
		Help: "play <pit> plays from that pit",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				fmt.Println("Invalid number of arguments")
				return
			}
			selectedPit, _ := strconv.Atoi(c.Args[0])
			c.Println("Hello " + fmt.Sprintf("%d", selectedPit))
		},
	})

	// run shell
	shell.Run()
}
