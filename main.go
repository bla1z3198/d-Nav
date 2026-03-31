package main

import (
	com "dnav/commands"
	"fmt"
)

// Main
func main() {
	fmt.Println("Welcome to dnav!")
	fmt.Println("help: init [filename] - init json with filename")
	fmt.Println("help: nav - calculate trajectory from current json")
	fmt.Println("help: upload [filename] - upload current flight plan to file")
	fmt.Println("help: exit - exit from dnav")
	com.HandleCommands()
}
