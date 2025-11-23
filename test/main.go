package main

import (
	"errors"
	"fmt"
)

func meanCalc  (nums [] int ) (*float64, error) {
	if len(nums) == 0 {
		return nil, errors.New("empty or nil array")
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	mean := float64(sum) / float64(len(nums))
	return &mean, nil
	//return nil, errors.New("not implemented")
}


func main () {

	fmt.Print("sid")

	mean, err := meanCalc([] int {1,2,3,4,5,6,7,8,9})
	if err == nil {
		fmt.Println("mean : ", *mean)
	}
	mean2, err2 := meanCalc([] int {})
	if err2 == nil {
		fmt.Println("mean : ", *mean2)
	} else {
		fmt.Println("error : ", err2, " for array : ")
	}

}
