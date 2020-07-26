package main

import "fmt"

// This is an example solution to the Simple Array Sum challenge:
// https://www.hackerrank.com/challenges/simple-array-sum/problem

func main() {
	var count int
	fmt.Scanf("%d\n", &count)

	sum := 0
	for i := 0; i < count; i++ {
		var value int
		fmt.Scanf("%d", &value)
		sum += value
	}

	fmt.Println(sum)
}
