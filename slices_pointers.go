package main

import (
	"fmt"
	"sort"
)

// append работает только по указателю. Воздействие append на срез без указателя не принесет рез-та
func main() {
	//var number int
	//fmt.Scanf("%d", &number)

	nums := []int{5, 1, 3, 9}
	sortSlice(nums)
	fmt.Println(nums)

	appendByVal(nums, 43)
	fmt.Println(nums)

	appendByPtr(&nums, 43)
	fmt.Println(nums)

}

func sortSlice(nums []int) {
	sort.Ints(nums)
}

func appendByVal(nums []int, n int) {
	nums = append(nums, n)
}

func appendByPtr(nums *[]int, n int) {
	*nums = append(*nums, n)
}
