package main

import "fmt"

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println(x, sum)
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		sum(i)
	}
	sum(99)
}