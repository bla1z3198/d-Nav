package core

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// Struct for waypoint
type Point struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
	ID int     `json:"id"`
	R  float64 `json:"r"`
}

type Upload struct {
	Line  float64 `json:"line"`
	Curve float64 `json:"curve"`
	ID    int     `json:"id"`
}

func Nav(points []Point) []Upload { // Func for calculate length of way
	// Previous sin and cos
	var prev_sin float64
	var prev_cos float64
	// Length of way
	var result float64
	// Length of first tangent
	var tangent float64
	// Distance between tangent points
	var tangent_line float64
	// Array for lines and curves
	upload := make([]Upload, 0)
	// Length of circle sector
	var sector float64
	// Angle between 2 tangent points
	var gamma float64
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
			angle0 := math.Asin(Radius/math.Sqrt(help_distance)) +
				math.Asin(Ay/math.Sqrt(help_distance))
			// Coordinates of first point
			x := math.Cos(angle0) * tangent
			y := math.Sin(angle0) * tangent

			tangent_line = math.Pow((Ax-sin*Radius)-(x), 2) +
				math.Pow((Ay+cos*Radius)-(y), 2)
			// Upload first part of way
			upload = append(upload, Upload{
				math.Sqrt(help_distance),
				0,
				i + 4999, // ID start at 4999
			})
		} else {
			tangent_line = math.Pow((Ax-sin*Radius)-(Ax-prev_sin*Radius), 2) +
				math.Pow((Ay+cos*Radius)-(Ay+prev_cos*Radius), 2)
		}
		// Argument for arccos of gamma angle
		arg := (1 - (tangent_line / (2 * Radius * Radius)))
		gamma = math.Acos(arg)
		if i > 0 {
			if (points[i].Y < points[i-1].Y) &&
				(points[i].Y < points[i+1].Y) {
				// Gamma for pit points
				gamma = 6.28319 - gamma
			}
		}
		// Calculate length off circle sector
		sector = Radius * gamma
		// Result length of way
		result += distanceAB + sector
		// Previous sin and cos
		prev_sin = sin
		prev_cos = cos
		// Append to upload
		upload = append(upload, Upload{
			distanceAB,
			sector,
			i + 5000, // ID start at 5000
		})
	}
	// Add first length
	result += tangent
	fmt.Println("Result length:", result)
	return upload
}

func UploadIntoFile(filename string, data []Upload) {
	// Sum of lines
	line_sum := 0.0
	// Sum of curves
	curve_sum := 0.0
	// Total string
	total := ""
	// Create (or re-write) file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("E03 - Can't create/read file...")
	}

	f.WriteString("dnav flight plan (for simulation use only!)\n\n")
	// Make newline symbol between data[i]
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "")

	for i := range data {
		encoder.Encode(data[i])
		line_sum += data[i].Line
		curve_sum += data[i].Curve
	}
	// Info string
	total = fmt.Sprintf("\nSum of lines = %.4f Sum of curves = %.4f\nTotal = %.4f",
		line_sum, curve_sum, line_sum+curve_sum)
	f.WriteString(total)
	// Take number of written symbols
	num, _ := f.Seek(0, 1)
	// Close file
	defer f.Close()

	fmt.Println("Success! Written", num, "bytes")
}
