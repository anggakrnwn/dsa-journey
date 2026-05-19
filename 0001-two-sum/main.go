package main

import "fmt"

func twoSum(nums []int, target int) []int {
	catatan := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		angkaSekarang := nums[i]
		jodoh := target - angkaSekarang

		indexJodoh, ada := catatan[jodoh]
		if ada {
			return []int{indexJodoh, i}
		}
		catatan[angkaSekarang] = i
	}

	return nil

}

func main() {
	input := []int{2, 7, 11, 15}
	target := 9
	ekspektasi := []int{0, 1}

	hasil := twoSum(input, target)

	fmt.Println("input      :", input)
	fmt.Println("ekspektasi :", ekspektasi)
	fmt.Println("hasil      :", hasil)
}
