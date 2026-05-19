package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxNumber := 0
	hitung := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			hitung++

			if hitung > maxNumber {
				maxNumber = hitung
			}
		} else {
			hitung = 0
		}
	}
	return maxNumber
}

func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	ekspektasi := 3

	hasil := findMaxConsecutiveOnes(input)

	fmt.Println("input: ", input)
	fmt.Println("ekspektasi: ", ekspektasi)
	fmt.Println("hasil: ", hasil)
}

/*
ans := []int{}

	// [1,1,0,1,1,1]

	ans[0] = nums[1]
	ans[1] = nums[1]
	ans[2] = nums[0]
	ans[3] = nums[1]
	ans[4] = nums[1]
	ans[5] = nums[1]

	return ans

*/
