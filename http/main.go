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
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes : ", len(bs))
	return len(bs), nil
}
