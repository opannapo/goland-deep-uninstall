package cmd

import (
	"os"
	"os/exec"
	"runtime"
)

var clearCmd map[string]func()

func init() {
	clearCmd = make(map[string]func())
	clearCmd["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearCmd["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clear(param []string) (retry bool) {
	value, ok := clearCmd[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported!")
	}

	return true
}
