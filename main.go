package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	score := 0.0
	for i := 0; i < 10; i++ {
		a := r.Intn(1000)
		b := r.Intn(1000)
		fmt.Printf("%v + %v = ", a, b)
		now := time.Now()
		fmt.Scan(&n)
		score = score + time.Since(now).Seconds()
		if n == a+b {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect: ", a+b)
		}
	}
	fmt.Println("Score: ", score)
}
