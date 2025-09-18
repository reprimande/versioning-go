package versioninggo

import "fmt"

func GetVersion() string {
	return "v1.0.2"
}

func PrintVersion() {
	fmt.Println(GetVersion())
}
