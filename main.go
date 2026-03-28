package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
)

type Point struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
	ID int     `json:"id"`
	R  float64 `json:"r"`
}

func New() []Point {
	points := make([]Point, 0)
	return points
}

func (p *Point) Init(points []Point) []Point {

	f, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatalln("E01 - Can't read file...")
	}
	e := json.Unmarshal(f, &points)
	if e != nil {
		log.Fatalln("E02 - Cant't parse JSON...")
	}
	fmt.Println("Загружено", len(points), "точек")

	return points
}

func Nav(points []Point) {
	var prev_sin float64
	var prev_cos float64
	var result float64
	for i := range points {
		if i == len(points)-1 {
			break
		}

		Ax := points[i].X
		Ay := points[i].Y
		Bx := points[i+1].X
		By := points[i+1].Y
		Radius := points[i].R

		line_len := math.Sqrt(math.Pow(Bx-Ax, 2) +
			math.Pow(By-Ay, 2))

		sin := (By - Ay) / line_len
		cos := math.Sqrt(1 - sin*sin)

		prepare := math.Pow((Ax-sin*Radius)-(Ax-prev_sin*Radius), 2) +
			math.Pow((Ay+cos*Radius)-(Ay+prev_cos*Radius), 2)

		//fmt.Println(i+1, "x:", Ax-prev_sin*Radius,
		//	i+1, "y:", Ay+prev_cos*Radius)
		//fmt.Println(i+2, "x:", Ax-sin*Radius,
		//	i+2, "y:", Ay+cos*Radius)
		fmt.Println()

		arg := ((prepare) / (2 * points[i].R * points[i].R)) - 1
		gamma := math.Acos(arg)
		gamma = ((gamma / math.Pi) * 180) - 180

		sector_len := (math.Pi * points[i].R * gamma) / 180

		result += line_len + sector_len

		prev_sin = sin
		prev_cos = cos
	}
	fmt.Println("Общий путь:", result)
}

func main() {
	fmt.Println("Запускаюсь...")
	var wg sync.WaitGroup
	data := (&Point{}).Init(New())

	for i := range 4 {
		wg.Go(func() {
			Nav(data[i : i+2])
			fmt.Println("Запущен процесс:", i+9000)
		})
	}
	wg.Wait()
	fmt.Println("Завершил", 4, "процессов")
}
