package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//fmt.Println(os.Args)
	//fileName := string(os.Args[1])
	f, err := os.Open(os.Args[1])
	//bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

/*

	usage : go run main.go myfile.txt
	go build main.go <- in order to print contents of the main.go
	./main main.go <- will print main.go
*/
