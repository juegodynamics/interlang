package main

import (
	"fmt"
	"math"
)

type Circle struct {
	centerX float64
	centerY float64
	radius  float64
}

func (shape *Circle) GetArea() float64 {
	return math.Pi * (math.Pow(shape.radius, 2))
}

type Square struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
}

func (shape *Square) GetArea() float64 {
	return (shape.maxX - shape.minX) * (shape.maxY - shape.minY)
}

type Shape interface {
	GetArea() float64
}

func GetTotalArea(shapes ...Shape) (area float64) {
	for _, shape := range shapes {
		area += shape.GetArea()
	}
	return
}

func main() {
	fmt.Println(GetTotalArea(
		&Circle{centerX: 0, centerY: 0, radius: 1 / (math.Sqrt(math.Pi))},
		&Square{minX: 0, minY: 0, maxX: 1, maxY: 1},
	))
}
