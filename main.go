package main

import "fmt"

func main() {
	question := 1

	// strings questions
	s1 := "dd"
	s2 := "ahbdge"

	fmt.Println(StringSolutions(s1, s2, question))

	// arrays questions
	nums := []int{1, 2, 0, 3, 4, 0, 0, 5}
	n := 2
	fmt.Println(ArraySolutions(nums, n, question))

}
