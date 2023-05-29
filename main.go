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
	a := r.Intn(1000)
	b := r.Intn(1000)

	fmt.Printf("%v + %v = ", a, b)
	now := time.Now()
	fmt.Scan(&n)
	if n == a+b {
		fmt.Println("Correct!")
		fmt.Printf("Time: %.2fs\n", time.Since(now).Seconds())
	} else {
		fmt.Println(a + b)
	}
}
