package main

import "fmt"

type Dimension struct {
	length int
	width int
	height int
}

func (d *Dimension) volume() int {
	d.length = 5
	return d.length * d.width * d.height
}

func (d Dimension) area() int {
	d.length = 20
	return d.length * d.width
}

func main () {
	fmt.Println("Hello World 😍")

	x, y := 5,10

	n := &x

	fmt.Println(*n);

	*n = 50

	fmt.Println(x);

	t := &y
	*t = *t + 30

	fmt.Println(t);
	fmt.Println(*t);
	fmt.Println(y);


	d := Dimension{length: 10, width: 20, height: 30}
	fmt.Println(d.volume())
	fmt.Println(d)
	fmt.Println(d.area())
	fmt.Println(d)
}
