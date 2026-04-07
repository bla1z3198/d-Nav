package core

import (
	"fmt"
	"math"
	"os"
)

// Global variables
var (
	sin      float64
	cos      float64
	distance float64
	tangent  float64
	result   float64
	sector   float64
	gamma    float64
	zero     bool
)

// Struct for waypoint
type Point struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
	ID int     `json:"id"`
	R  float64 `json:"r"`
}

// Struct for upload
type Upload struct {
	FromPoint int
	ToPoint   int
	Line      float64
	Curve     float64
	Rad       float64
}

// Struct for tangent points
type Tangent struct {
	X1      float64
	X2      float64
	Y1      float64
	Y2      float64
	Line    float64
	tangent float64
}

func (tp *Tangent) TangentPoints(
	Ax, Ay, Bx, By, R *float64,
	zero bool) Tangent {
	// Calculate distance between two points
	distance = math.Sqrt(math.Pow(*Bx-*Ax, 2) +
		math.Pow(*By-*Ay, 2))
	// Caclulate tangent points with previous sin and cos
	if !zero {
		tp.X1 = *Ax - sin*(*R)
		tp.Y1 = *Ay + cos*(*R)
	}
	// Calculate first part of way
	if zero {
		dt := math.Sqrt((math.Pow(*Ax, 2) +
			math.Pow(*Ay, 2)))
		// First tangent line
		tangent := math.Sqrt(dt*dt -
			(*R * (*R)))
		// Angle between x axis and tangent
		angle := math.Asin(*R/dt) +
			math.Asin(*Ay/dt)
		// Coordinates of first point
		tp.X1 = math.Cos(angle) * tangent
		tp.Y1 = math.Sin(angle) * tangent
		tp.Line += dt
	}
	// Calculate sin and cos between x-axis and distance line
	sin = (*By - *Ay) / distance
	cos = math.Sqrt(1 - sin*sin)
	// Second pair of tangent points always calculates like this
	tp.X2 = *Ax - sin*(*R)
	tp.Y2 = *Ay + cos*(*R)
	tp.Line = distance

	return *tp
}

func Nav(points []Point) []Upload { // Func for calculate length of way
	// Array for lines and curves
	upload := make([]Upload, 0)
	// Cycle for Points
	for i := range points {
		// Calculate length for pairs of points (not single)
		if i == len(points)-1 {
			break
		}
		// Set zero indicator
		if i == 0 {
			zero = true
		} else {
			zero = false
		}
		// Create Tangent obj and call function for calculate tangent points
		t := Tangent{}
		tp := t.TangentPoints(&points[i].X,
			&points[i].Y,
			&points[i+1].X,
			&points[i+1].Y,
			&points[i].R,
			zero)
		// Calculate tangent line
		tangent = math.Sqrt(math.Pow(tp.X2-tp.X1, 2) +
			math.Pow(tp.Y2-tp.Y1, 2))
		// Argument for arccos of gamma angle
		arg := (1 - (tangent * tangent / (2 * points[i].R * points[i].R)))
		gamma = math.Acos(arg)
		if i > 0 {
			if (points[i].Y < points[i-1].Y) &&
				(points[i].Y < points[i+1].Y) {
				// Gamma for pit points
				gamma = 6.28319 - gamma
			}
		}
		// Calculate length of circle sector
		sector = points[i].R * gamma
		fmt.Println("--sector--", sector)
		// Result length of way
		result += tp.Line + sector
		// Append to upload
		upload = append(upload, Upload{
			i,
			i + 1,
			tp.Line,
			sector,
			(gamma * 180) / math.Pi,
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
	// String for lines
	forward := ""
	// String for turns
	turn := ""
	// Total string
	total := ""
	// Create (or re-write) file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("E03 - Can't create/read file...")
	}

	f.WriteString("dnav flight plan (for simulation use only!)\n\n")

	for i := range data {
		// Format strings
		forward = fmt.Sprintf("Move forward for %.4f metres\n\n",
			data[i].Line)
		turn = fmt.Sprintf("Turn left for %.4f degrees\n",
			data[i].Rad)
		// Write strings
		f.WriteString("Point - " + fmt.Sprint(i) + "\n")
		f.WriteString(turn)
		f.WriteString(forward)

		line_sum += data[i].Line
		curve_sum += data[i].Curve
	}
	// Info string
	total = fmt.Sprintf("\nSum of lines = %.4f metres\nSum of curves = %.4f metres\nTotal = %.4f metres",
		line_sum, curve_sum, line_sum+curve_sum)
	f.WriteString(total)
	// Take number of written symbols
	num, _ := f.Seek(0, 1)
	// Close file
	defer f.Close()

	fmt.Println("Success! Written", num, "bytes")
}
