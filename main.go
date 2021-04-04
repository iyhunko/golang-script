package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	number, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Please enter correct number of elements")
		os.Exit(1)
	}
	shouldReturnAll := flag.Arg(1)

	randSource := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(randSource)

	if shouldReturnAll == "all" {
		numbers := make([]int, number)
		for i := range numbers {
			numbers[i] = 1 + i
		}
		customRand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
		fmt.Println(numbers)
	} else {
		fmt.Println(customRand.Intn(number-1) + 1)
	}
}
