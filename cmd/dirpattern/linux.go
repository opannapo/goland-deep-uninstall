package dirpattern

type Linux struct {
}

func (l *Linux) V_2020_2() map[string]string {
	return map[string]string{
		"./.config/JetBrains/GoLand2020.2":      "./.config/JetBrains/GoLand2020.2",
		"./.cache/JetBrains/GoLand2020.2":       "./.cache/JetBrains/GoLand2020.2",
		"./.java/.userPrefs/jetbrains/goland":   "./.java/.userPrefs/jetbrains/goland",
		"./.local/share/JetBrains/GoLand2020.2": "./.local/share/JetBrains/GoLand2020.2",
		"./.local/share/JetBrains/Goland":       "./.local/share/JetBrains/Goland",
	}
}

func (l *Linux) V_2021_1() map[string]string {
	return map[string]string{
		"./.config/JetBrains/GoLand2021.1":      "./.config/JetBrains/GoLand2021.1",
		"./.cache/JetBrains/GoLand2021.1":       "./.cache/JetBrains/GoLand2021.1",
		"./.java/.userPrefs/jetbrains/goland":   "./.java/.userPrefs/jetbrains/goland",
		"./.local/share/JetBrains/GoLand2021.1": "./.local/share/JetBrains/GoLand2021.1",
		"./.local/share/JetBrains/Goland":       "./.local/share/JetBrains/Goland",
	}
}
