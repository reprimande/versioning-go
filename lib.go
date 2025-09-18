package versioninggo

import "fmt"

func GetVersion() string {
	return "v2.0.1"
}

func PrintVersion() {
	fmt.Println(GetVersion())
}
