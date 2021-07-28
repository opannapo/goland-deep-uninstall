package dirpattern

type Darwin struct {
}

func (d *Darwin) V_2020_2() map[string]string {
	return map[string]string{
		"./.config/JetBrains/GoLand2020.2":      "./.config/JetBrains/GoLand2020.2",
		"./.cache/JetBrains/GoLand2020.2":       "./.cache/JetBrains/GoLand2020.2",
		"./.java/.userPrefs/jetbrains/goland":   "./.java/.userPrefs/jetbrains/goland",
		"./.local/share/JetBrains/GoLand2020.2": "./.local/share/JetBrains/GoLand2020.2",
		"./.local/share/JetBrains/Goland":       "./.local/share/JetBrains/Goland",
	}
}

func (d *Darwin) V_2021_1() map[string]string {
	return map[string]string{
		"./Library/Application Support/JetBrains/GoLand2021.1":              "./Library/Application Support/JetBrains/GoLand2021.1",
		"./Library/Caches/JetBrains/GoLand2021.1":                           "./Library/Caches/JetBrains/GoLand2021.1",
		"./Library/Logs/JetBrains/GoLand2021.1":                             "./Library/Logs/JetBrains/GoLand2021.1",
		"./Library/Saved Application State/com.jetbrains.goland.savedState": "./Library/Saved Application State/com.jetbrains.goland.savedState",
		"./Library/Preferences/jetbrains.goland.4f014ddc.plist":             "./Library/Preferences/jetbrains.goland.4f014ddc.plist",
		"./Library/Preferences/com.jetbrains.goland.plist":                  "./Library/Preferences/com.jetbrains.goland.plist",
		"./Library/Preferences/jetbrains.jetprofile.asset.plist":            "./Library/Preferences/jetbrains.jetprofile.asset.plist",
	}
}
