package main

import "fmt"

func hasDuplicate(nums []int) bool {
	check := make(map[int]bool)
	for _, val := range nums {
		if check[val] {
			return false
		}
		check[val] = true
	}
	return true
}

func main(){
	nums := []int{2, 2, 3, 4}
	fmt.Println(hasDuplicate(nums))
}
