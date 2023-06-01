package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func generateRandNum(digit int) (int, int) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	max := int(math.Pow10(digit))
	a := r.Intn(max)
	b := r.Intn(max)
	return a, b
}

func askQuestion(question string, answer int) float64 {
	var userAnswer int
	now := time.Now()

	fmt.Printf("%v = ", question)
	fmt.Scan(&userAnswer)
	if userAnswer == answer {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect. The correct answer is", answer)
	}
	return time.Since(now).Seconds()
}

func add() float64 {
	a, b := generateRandNum(3)
	return askQuestion(fmt.Sprintf("%d + %d", a, b), a+b)
}

func sub() float64 {
	a, b := generateRandNum(3)
	return askQuestion(fmt.Sprintf("%d - %d", a, b), a-b)
}

func mul() float64 {
	a, b := generateRandNum(2)
	return askQuestion(fmt.Sprintf("%d * %d", a, b), a*b)
}

func main() {
	fmt.Println(add())
	fmt.Println(sub())
	fmt.Println(mul())
}
