package cmd

import (
	"fmt"
	"log"
	"strings"
)

func confirmMessage() bool {
	fmt.Print("Delete ? y/n » ")
	in, err := reader.ReadString('\n')
	_, _ = reader.Discard(0)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(strings.ReplaceAll(in, "\n", "")) == "y" {
		return true
	} else {
		return false
	}
}
