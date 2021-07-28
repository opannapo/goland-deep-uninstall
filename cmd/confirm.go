package cmd

import (
	"fmt"
	"log"
	"strings"
)

func confirmDeleteMessage(len int) bool {
	fmt.Printf("Delete [%s %d selected %s] ? y/n » ", colorRed, len, colorReset)
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

func confirmCreateMessage() bool {
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
