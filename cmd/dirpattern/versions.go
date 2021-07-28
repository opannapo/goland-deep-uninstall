package dirpattern

const (
	Version2020_2 = "2020.2"
	Version2021_1 = "2021.1"
	VersionDemo   = "demo"
)

var darwin = new(Darwin)
var linux = new(Linux)
var demo = new(Demo)

func DarwinDirMap(version string) map[string]string {
	switch version {
	case Version2020_2:
		return darwin.V_2020_2()
	case Version2021_1:
		return darwin.V_2021_1()
	}

	return nil
}

func LinuxDirMap(version string) map[string]string {
	switch version {
	case Version2020_2:
		return linux.V_2020_2()
	case Version2021_1:
		return linux.V_2021_1()
	}

	return nil
}

func DemoMap() map[string]string {
	return demo.Dummy()
}
