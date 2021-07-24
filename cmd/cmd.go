package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type argsAttributes struct {
	Label    string
	Function func(param []string) bool
}

var mainCommands = map[string]*argsAttributes{}
var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
	mainCommands = map[string]*argsAttributes{
		"scan": {
			Label:    "Scanning GOLAND directory for %s version",
			Function: findAllDirectoryFiles,
		},
		"help": {
			Label:    "Help",
			Function: showHelp,
		},
		"clear": {
			Label:    "Clear",
			Function: clear,
		},
	}

	//alias
	mainCommands["s"] = mainCommands["scan"]
	mainCommands["h"] = mainCommands["help"]
	mainCommands["c"] = mainCommands["clear"]
}
func Execute() {
	for {
		fmt.Print("» ")
		in, _ := reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", -1)

		splitArgParam := strings.Split(in, " ")
		arg := mainCommands[splitArgParam[0]]
		if arg != nil {
			retry := arg.Function(splitArgParam[1:])
			if retry {
				Execute()
			}
			break
		} else {
			fmt.Printf("Wrong Args '%+v' \n", in)
		}
	}
}
