package main

import (
	"fmt"
	"math"
)

//function variation
// func circleArea(c *Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// func main() {
// 	//initialization options
// 	//1// var circle Circle
// 	//2// circle := new(Circle)
// 	//3// circle := Circle{x: 0, y: 0, r: 5}
// 	circle := Circle{0, 0, 5}
// 	// fmt.Println(circle.x, circle.y, circle.r)
// 	circle.x = 10
// 	circle.y = 2
// 	// fmt.Println(circle.x, circle.y, circle.r)
// 	fmt.Println(circleArea(&circle))
// }

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

//methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) distance() float64 {
	return r.x1 + r.x2 + r.y1 + r.y2
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return r.x1 + r.x2 + r.y1 + r.y2
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

//embedded types
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is, ", p.Name)
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, p := range shapes {
		perimeter += p.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(c.area())
	fmt.Println(r.distance())
	fmt.Println(r.area())
	// person := new(Person)
	// person.Name = "Chuck"
	// person.Talk()
	// a := new(Android)
	// a.Name = "turkey"
	// a.Person.Talk()
	// fmt.Println(totalArea(&c, &r))
	// fmt.Println(totalPerimeter(&c, &r))

}
