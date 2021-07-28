package dirpattern

type Demo struct {
}

func (s *Demo) Dummy() map[string]string {
	return map[string]string{
		"./DemoDeppUninstallGoland":     "./DemoDeppUninstallGoland",
		"./.DemoDeppUninstallGoland":    "./.DemoDeppUninstallGoland",
		"./DemoDeppUninstallGoland.txt": "./DemoDeppUninstallGoland.txt",
	}
}
