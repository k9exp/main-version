package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("could not read build info")
		return
	}
	fmt.Printf("Go Version: %s\n", info.GoVersion)
	fmt.Printf("App Version: %s\n", info.Main.Version)
	fmt.Printf("App Checksum: %s\n", info.Main.Sum)
}
