package main

import "fmt"

func Printer(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i] - 100)
	}
}
