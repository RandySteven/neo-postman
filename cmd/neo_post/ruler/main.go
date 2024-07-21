package main

import (
	"fmt"
	"os"
)

const templateCommand = "create models %s -p -r -u -h -q\n"

/**exp
-p = payload
-r = repositories
-u = usecases
-h = handlers
-q = queries
*/

func main() {
	cliV2()
}

func cliV2() {
	cli := NewCli(os.Args[1:])
	if len(cli.command) == 0 {
		fmt.Println("No command provided")
		return
	}

	switch cli.command[0] {
	case "model":
		if len(cli.command) < 2 {
			fmt.Println("Model name is required")
			return
		}
		cli.createModel()
	case "-r", "--repository":
		if len(cli.command) < 2 {
			fmt.Println("Repository name is required")
			return
		}
		cli.createRepository()
	default:
		fmt.Println("Unknown command")
	}
}
