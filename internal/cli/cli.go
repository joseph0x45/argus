package cli

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Get some help")
}

func DispatchCommands(args []string) int {
	if len(args) == 1 {
		printUsage()
		return 0
	}
	cmd := args[1]
	switch cmd {
	case "help":
		printUsage()
		return 0
	case "add-channel":
		return addChannel(args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command '%s'\n", cmd)
		printUsage()
		return 1
	}
}
