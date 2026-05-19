package main

import "fmt"

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[2*i] = nums[i]
		ans[(2*i)+1] = nums[i+n]

	}

	return ans
}

func main() {
	input := []int{2, 5, 1, 3, 4, 7}
	n := 3
	ekspektasi := []int{2, 3, 5, 4, 1, 7}

	hasil := shuffle(input, n)

	fmt.Println("input: ", input)
	fmt.Println("ekspektasi: ", ekspektasi)
	fmt.Println("hasil: ", hasil)
}
