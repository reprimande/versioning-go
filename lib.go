package versioninggo

import "fmt"

func GetVersion() string {
	return "v3.0.0-beta"
}

func PrintVersion() {
	fmt.Println(GetVersion())
}
