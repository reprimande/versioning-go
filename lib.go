package versioninggo

func GetVersion() string {
	return "v0.0.1"
}

func PrintVersion() string {
	return "Version: " + GetVersion()
}
