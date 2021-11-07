package cmd

import (
	"bytes"
	"fmt"
	"goland-deep-uninstall/cmd/dirpattern"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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
		"-iname",
		`*goland*`,
		"-print",
	}

	homeUserDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v %s(%s)%s\n", fmt.Sprintf(mainCommands["scan"].Label, param[0]), colorCyan, homeUserDir, colorReset)

	cmd := exec.Command(strCmd, argsCmd...)
	cmd.Dir = homeUserDir
	log.Printf("cmd %s %s\n", homeUserDir, cmd.String())

	var stderr bytes.Buffer
	var findOut bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &findOut

	chCmdRun := make(chan struct{})
	go func() {
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Err %s %s\n", stderr.String(), err)
		}
		close(chCmdRun)
	}()

	<-chCmdRun
	dirToRemove := map[int]string{}
	splitStr := strings.Split(findOut.String(), "\n")

	selectedMapPattern, err := getMapPatternByOs(param[0])
	if err != nil {
		log.Fatal(err)
	}

	chColoring := make(chan struct{})
	go func() {
		var lastInsertPattern string
		for i, s := range splitStr {
			if len(s) > 0 {
				time.Sleep(time.Millisecond * 5)
				indicator := selectedMapPattern[s]
				if len(indicator) > 0 {
					fmt.Printf("%s[✔] %s %s\n", colorLightCyan, s, colorReset)
					dirToRemove[i] = s
					lastInsertPattern = s
				} else {
					if len(lastInsertPattern) > 0 && strings.Contains(s, lastInsertPattern) {
						fmt.Printf("%s    %s %s\n", colorCyan, s, colorReset)
					} else {
						fmt.Printf("[x] %s\n", s)
					}
				}
			}
		}
		close(chColoring)
	}()

	<-chColoring
	if len(dirToRemove) < 1 {
		fmt.Printf("%sVersion %s Not Found!%s\n", colorYellow, param[0], colorReset)

		if param[0] == dirpattern.VersionDemo {
			fmt.Printf("%s\nMakesure this directory has been created on your computer :%s\n", colorCyan, colorReset)
			mapDemo := dirpattern.DemoMap()
			for _, dir := range mapDemo {
				fmt.Printf("%s » %s %s\n", colorCyan, dir, colorReset)
			}
		}

		return true
	}

	ok := confirmDeleteMessage(len(dirToRemove))
	if ok {
		for _, d := range dirToRemove {
			rmCommand(d)
		}

	} else {
		fmt.Println("Exit")
	}
	_, _ = reader.Discard(0)
	return false
}

func rmCommand(dir string) {
	rmCmd := "rm"
	argsRmCmd := []string{
		"-r",
		dir,
	}

	cmd := exec.Command(rmCmd, argsRmCmd...)
	fmt.Printf("%sDeleting...%s %s%+v%s \n", colorRed, colorReset, colorCyan, cmd, colorReset)

	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	homeUserDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Dir = homeUserDir
	err = cmd.Run()
	if err != nil {
		log.Fatal(stderr.String())
	}

	fmt.Printf("%sDeleted » %s %s%+v%s \n", colorGreen, colorReset, colorRed, cmd, colorReset)
}

func getMapPatternByOs(versionParam string) (map[string]string, error) {
	if versionParam == dirpattern.VersionDemo {
		return dirpattern.DemoMap(), nil
	} else {
		switch runtime.GOOS {
		case "darwin":
			return dirpattern.DarwinDirMap(versionParam), nil
		case "linux":
			return dirpattern.LinuxDirMap(versionParam), nil
		}
	}

	return nil, fmt.Errorf("pattern not found")
}
