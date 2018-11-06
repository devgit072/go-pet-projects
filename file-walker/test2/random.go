package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {
	count := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i:= 0 ;i<1000 ;i++ {
		val := rand.Intn(1000)
		if val == 1 {
			count++
		}

	}
	fmt.Println("Count: ", count)

}
