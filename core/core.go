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
	// Length of first tangent
	var tangent float64
	// Distance between tangent points
	var tangent_line float64
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
		distanceAB := math.Sqrt(math.Pow(Bx-Ax, 2) +
			math.Pow(By-Ay, 2))

		sin := (By - Ay) / distanceAB
		cos := math.Sqrt(1 - sin*sin)
		// Calculate length between tangent points
		if i == 0 {
			// Distance between start and first point
			help_distance := (math.Pow(Ax, 2) + math.Pow(Ay, 2))
			// First tangent
			tangent = math.Sqrt(help_distance -
				(Radius * Radius))
			// Angle between x axis and tangent
			angle := math.Asin(Radius/math.Sqrt(help_distance)) +
				math.Asin(Ay/math.Sqrt(help_distance))
			// Coordinates of first point
			x := math.Cos(angle) * tangent
			y := math.Sin(angle) * tangent

			tangent_line = math.Pow((Ax-sin*Radius)-(x), 2) +
				math.Pow((Ay+cos*Radius)-(y), 2)
		} else {
			tangent_line = math.Pow((Ax-sin*Radius)-(Ax-prev_sin*Radius), 2) +
				math.Pow((Ay+cos*Radius)-(Ay+prev_cos*Radius), 2)
		}
		// Argument for arccos of gamma angle
		arg := (1 - (tangent_line / (2 * Radius * Radius)))
		gamma := math.Acos(arg)
		if i > 0 {
			if (points[i].Y < points[i-1].Y) &&
				(points[i].Y < points[i+1].Y) {
				gamma = 6.28319 - gamma
			}
		}
		// Calculate length off circle sector
		sector := Radius * gamma
		// Result length of way
		result += distanceAB + sector
		// Previous sin and cos
		prev_sin = sin
		prev_cos = cos
		fmt.Println("Angle =", (gamma*180)/3.1415)
	}
	fmt.Println("Result length:", result+tangent)
}
