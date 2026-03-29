package core

import (
	"fmt"
	"math"
)

// Struct for waypoint
type Point struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
	ID int     `json:"id"`
	R  float64 `json:"r"`
}

func Nav(points []Point) { // Func for calculate length of way
	// Previous sin and cos
	var prev_sin float64
	var prev_cos float64
	// Length of way
	var result float64
	// Cycle for Points
	for i := range points {
		// Calculate length for pairs of points (not single)
		if i == len(points)-1 {
			break
		}
		// Define the now and next points
		Ax := points[i].X
		Ay := points[i].Y
		Bx := points[i+1].X
		By := points[i+1].Y
		Radius := points[i].R
		// Calculate length between A and B
		line_len := math.Sqrt(math.Pow(Bx-Ax, 2) +
			math.Pow(By-Ay, 2))

		sin := (By - Ay) / line_len
		cos := math.Sqrt(1 - sin*sin)
		// Calculate length between tangent points
		prepare := math.Pow((Ax-sin*Radius)-(Ax-prev_sin*Radius), 2) +
			math.Pow((Ay+cos*Radius)-(Ay+prev_cos*Radius), 2)
		// Argument for arccos of gamma angle
		arg := ((prepare) / (2 * points[i].R * points[i].R)) - 1
		gamma := math.Acos(arg)
		gamma = ((gamma / math.Pi) * 180) - 180
		// Calculate length off circle sector
		sector_len := (math.Pi * points[i].R * gamma) / 180
		// Result length of way
		result += line_len + sector_len
		// Previous sin and cos
		prev_sin = sin
		prev_cos = cos
	}
	fmt.Println("Result length:", result)
}
