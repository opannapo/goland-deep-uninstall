package cmd

import (
	"fmt"
)

func showHelp(param []string) (retry bool) {
	str := `
	%sThis is for personal use!%s

	Scanning Directory Commands
	-----------------------------------
  	%s-s <version>, scan <version>
  	Example »  scan 2020.2, s 2020.2%s

	Help Commands 
	-----------------------------------
  	%s-h, help%s 
`

	fmt.Printf(str, colorRed, colorReset, colorCyan, colorReset, colorCyan, colorReset)
	return true
}
