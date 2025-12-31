package main

import "fmt"

func twosum(nums []int, target int) []int {
	var result int
	for i := 0; i < len(nums); i++ {
		for j := i + 1 ; j < len(nums);j++ {
			result = nums[i] + nums[j]
			if result == target {
				return []int{i, j};
			}
		}
	}
	return []int{} 
}

func main() {
	arr1 := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twosum(arr1, target))
}
