package main
import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length, width float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

type circle struct {
	radius float64
}

func print(s shape) {
	fmt.Println(s.area())
	fmt.Println(r.perimeter())
}

func main() {
	r := rect {
		length: 8,
		width: 3,
	}
	c := circle {
		radius: 5,
	}
}
