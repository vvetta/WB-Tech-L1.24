package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y,}
}

func(p *Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	kremlinPoint := NewPoint(55.751806, 37.619058)
	moscowCityPoint := NewPoint(55.748499, 37.536791)

	fmt.Printf("KremlinPoint: %v \nMoscowCityPoint: %v\n", kremlinPoint, moscowCityPoint)

	distance := kremlinPoint.Distance(*moscowCityPoint)

	fmt.Printf("Расстояние от точки до точки на плоскости: %.6f", distance)
}
