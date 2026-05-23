package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		num1 := 0
		num2 := 0

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + carry

		digit := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: digit,
		}

		curr.Next = node
		curr = curr.Next
	}

	return dummy.Next
}

func helperList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, n := range nums {
		curr.Next = &ListNode{
			Val: n,
		}
		curr = curr.Next
	}

	return dummy.Next
}

func main() {

	var n1 int
	fmt.Print("jumlah digit l1: ")
	fmt.Scan(&n1)

	arr1 := make([]int, n1)

	fmt.Println("masukkan digit l1:")
	for i := 0; i < n1; i++ {
		fmt.Scan(&arr1[i])
	}

	var n2 int
	fmt.Print("jumlah digit l2: ")
	fmt.Scan(&n2)

	arr2 := make([]int, n2)

	fmt.Println("masukkan digit l2:")
	for i := 0; i < n2; i++ {
		fmt.Scan(&arr2[i])
	}

	l1 := helperList(arr1)
	l2 := helperList(arr2)

	result := addTwoNumbers(l1, l2)

	fmt.Print("[")

	for result != nil {
		fmt.Print(result.Val)

		if result.Next != nil {
			fmt.Print(",")
		}

		result = result.Next
	}

	fmt.Println("]")
}
