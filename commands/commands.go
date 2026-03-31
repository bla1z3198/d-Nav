package commands

import (
	"bufio"
	core "dnav/core"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Handle commands
func HandleCommands() {
	input := ""
	splitcom := make([]string, 2)
	waypoints := make([]core.Point, 0)
	data := make([]core.Upload, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("-> ")
		scanner.Scan()
		input = scanner.Text()
		splitcom = strings.Fields(input)

		switch splitcom[0] {
		case "init":
			if len(splitcom) > 1 {
				waypoints = Init(New(), splitcom[1])
			} else {
				fmt.Println("error: you need to type init [filename]")
			}
		case "nav":
			data = core.Nav(waypoints)
		case "exit":
			fmt.Println("Bye.")
			os.Exit(0)
		case "help":
			fmt.Println("help: init [filename] - init json with filename")
			fmt.Println("help: nav - calculate trajectory from current json")
			fmt.Println("help: upload [filename] - upload current flight plan to file")
			fmt.Println("help: exit - exit from dnav")
		case "upload":
			if len(splitcom) > 1 {
				core.UploadIntoFile(splitcom[1], data)
			} else {
				fmt.Println("error: you need to type upload [filename]")
			}
		default:
			fmt.Println("Command", splitcom[0], "not found")
			fmt.Println("Try to type 'help'")
		}
	}
}

// Creating new waypoints list
func New() []core.Point {
	points := make([]core.Point, 0)
	return points
}

// Load waypoints config form ".json"
func Init(points []core.Point, filename string) []core.Point {

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("E01 - Can't read file...")
		return nil
	}
	e := json.Unmarshal(f, &points)
	if e != nil {
		fmt.Println("E02 - Can't parse JSON...")
		return nil
	}
	fmt.Println("Loaded", len(points), "points")

	return points
}
