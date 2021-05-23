package main

/* imports */
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

/* main method */
func main() {

	resp, err := http.Get("https://insurance.mobikwik.com/api/home/ping")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	/* bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs)) */

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	sq := square{16}
	tr := triangle{10, 20}

	printArea(sq)
	printArea(tr)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes : ", len(bs))
	return len(bs), nil
}

/* part of the assignment on interface */
type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}
func printArea(s shape) {
	fmt.Println(s.area())
}
