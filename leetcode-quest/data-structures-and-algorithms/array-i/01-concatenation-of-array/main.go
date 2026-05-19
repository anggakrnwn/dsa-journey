package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}

func main() {
	input := []int{1, 2, 1}
	ekspektasi := []int{1, 2, 1, 1, 2, 1}

	hasil := getConcatenation(input)

	fmt.Println("input: ", input)
	fmt.Println("ekspektasi: ", ekspektasi)
	fmt.Println("hasil: ", hasil)

}
