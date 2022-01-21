package helpers

import (
	"math/rand"
	"time"
)

type SomeType struct {
	TypeName    string
	TypeAddress string
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixMilli())
	value := rand.Intn(n)
	return value
}

func main() {

}
