package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var patternGoland2020_2 = map[string]string{
	"./.config/JetBrains/GoLand2020.2":      "./.config/JetBrains/GoLand2020.2",
	"./.cache/JetBrains/GoLand2020.2":       "./.cache/JetBrains/GoLand2020.2",
	"./.java/.userPrefs/jetbrains/goland":   "./.java/.userPrefs/jetbrains/goland",
	"./.local/share/JetBrains/GoLand2020.2": "./.local/share/JetBrains/GoLand2020.2",
	"./.local/share/JetBrains/Goland":       "./.local/share/JetBrains/Goland",
}

var patternGoland2021_1 = map[string]string{
	"./.config/JetBrains/GoLand2020.2":      "./.config/JetBrains/GoLand2020.2",
	"./.cache/JetBrains/GoLand2020.2":       "./.cache/JetBrains/GoLand2020.2",
	"./.java/.userPrefs/jetbrains/goland":   "./.java/.userPrefs/jetbrains/goland",
	"./.local/share/JetBrains/GoLand2020.2": "./.local/share/JetBrains/GoLand2020.2",
	"./.local/share/JetBrains/Goland":       "./.local/share/JetBrains/Goland",
}

var patternTest = map[string]string{
	"./DeppUninstallTestGoland":  "./DeppUninstallTestGoland",
	"./.DeppUninstallTestGoland": "./.DeppUninstallTestGoland",
}

const (
	version2020_2 = "2020.2"
	version2021_1 = "2021.1"
	versionTest   = "test"
)

func findAllDirectoryFiles(param []string) (retry bool) {
	if len(param) < 1 {
		fmt.Printf("%sError :: Empty Version %s\n", colorRed, colorReset)
		fmt.Printf("%sExample :: scan 2020.2 %s\n", colorYellow, colorReset)
		return true
	}

	param[0] = strings.ToLower(param[0])

	strCmd := "find"
	argsCmd := []string{
		".",
		"-type",
		"d",
		"-iname",
		"*goland*",
		"-print",
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v %s(%s)%s\n", fmt.Sprintf(mainCommands["scan"].Label, param[0]), colorCyan, dir, colorReset)

	cmd := exec.Command(strCmd, argsCmd...)
	cmd.Dir = dir
	findOut, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	dirToRemove := map[int]string{}
	splitStr := strings.Split(string(findOut), "\n")

	selectedMapPattern := map[string]string{}
	switch param[0] {
	case version2020_2:
		selectedMapPattern = patternGoland2020_2
		break
	case version2021_1:
		selectedMapPattern = patternGoland2021_1
		break
	case versionTest:
		selectedMapPattern = patternTest
	}

	for i, s := range splitStr {
		if len(s) > 0 {
			if len(selectedMapPattern[s]) != 0 {
				fmt.Printf("%s[✔] %s %s\n", colorCyan, s, colorReset)
				dirToRemove[i] = s
			} else {
				fmt.Printf("[x] %s\n", s)
			}
		}
	}

	if len(dirToRemove) < 1 {
		fmt.Printf("%sVersion %s Not Found!%s\n", colorYellow, param[0], colorReset)
		return true
	}

	ok := confirmMessage()
	if ok {
		fmt.Println("Deleting...")

	} else {
		fmt.Println("Exit")
	}
	_, _ = reader.Discard(0)
	return false
}
