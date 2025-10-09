package versioninggo

import "fmt"

func GetVersion() string {
	return "v3.0.0-beta2"
}

func PrintVersion() {
	fmt.Println(GetVersion())
}
