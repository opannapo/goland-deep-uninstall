package main

import (
	"fmt"
	"goland-deep-uninstall/cmd"
)

const opening = `
                        - MIZQIN -
▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉
▉▉       ▉     ▉▉▉▉   ▉▉         ▉▉         ▉▉▉   ▉▉   ▉▉▉   ▉▉
▉▉   ▉    ▉▉    ▉▉▉▉▉▉▉▉▉▉▉▉▉▉   ▉▉   ▉▉▉   ▉▉▉▉▉▉▉▉     ▉   ▉▉
▉▉   ▉▉    ▉▉   ▉▉▉   ▉▉   ▉▉▉   ▉▉   ▉     ▉▉▉   ▉▉   ▉     ▉▉
▉▉   ▉▉▉   ▉▉   ▉▉▉   ▉▉   ▉▉▉▉▉▉▉▉   ▉▉▉    ▉▉   ▉▉   ▉     ▉▉
▉▉   ▉▉▉▉▉▉▉▉   ▉▉▉   ▉▉         ▉▉         ▉▉▉   ▉▉   ▉▉▉   ▉▉
▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉
`

func main() {
	fmt.Println(opening)
	cmd.Execute()
}
