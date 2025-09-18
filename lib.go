package versioninggo

import "fmt"

func GetVersion() string {
	return "v0.0.1"
}

func PrintVersion() string {
	fmt.Println(GetVersion())
}
