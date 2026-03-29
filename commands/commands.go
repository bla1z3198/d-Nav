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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("-> ")
		scanner.Scan()
		input = scanner.Text()
		splitcom = strings.Fields(input)

		switch splitcom[0] {
		case "init":
			waypoints = Init(New(), splitcom[1])
		case "nav":
			core.Nav(waypoints)
		case "exit":
			fmt.Println("Bye.")
			os.Exit(0)
		default:
			fmt.Println("Command", splitcom[0], "not found")
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
