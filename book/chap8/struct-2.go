package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}
<<<<<<< HEAD
//
//func circleArea(c *Circle) float64 {
//	return math.Pi * c.r*c.r
//}

func (c *Circle) area() float64 {
=======

func circleArea(c *Circle) float64 {
>>>>>>> b1297ab1f5a67286a11be9a22c1debfac166f3e2
	return math.Pi * c.r*c.r
}

func main() {
	var rx1, ry1 float64 = 0,0
	var rx2, ry2 float64 = 6,6
//
	c := Circle{0,0,10}
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
<<<<<<< HEAD
	fmt.Println(c.area())
=======
	fmt.Println(circleArea(&c))
>>>>>>> b1297ab1f5a67286a11be9a22c1debfac166f3e2
}