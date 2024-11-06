package algorithms

import (
	"fmt"
	"sort"
	"strings"
)

func find3Sum(nums []int, target int) bool {
	found := false
	length := len(nums)
	// sort the list in ascending order
	sort.Ints(nums)
	// iterate over the list find 3 distinct values
	// use 3 vars, start, low, high
	start := 0 // start at first item
	low := 1
	high := length - 1
	for start < length-2 {
		// low is start+1 and high is last index
		// unitl low meets high
		for low < high {
			for j := low; j < high; j++ {
				// add start, low and high if sum = target return true
				sum := nums[start] + nums[low] + nums[high]
				if sum == target {
					found = true
					return found
				}
				// if sum is less than target low is too low so bump up index for low
				if sum < target {
					low++
				}
				// if sum is greater that target high is too high so decrease index for high
				if sum > target {
					high--
				}
			}
		}
		start++
		// reset low and high after bummping start
		low = start + 1
		high = length - 1
	}
	return found

}

// Driver code
func RunFind3Sum() {
	numsLists := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{-1, 2, 1, -4, 5, -3},
		{2, 3, 4, 1, 7, 9},
		{1, -1, 0},
		{2, 4, 2, 7, 6, 3, 1},
	}
	tList := []int{10, 7, 20, -1, 8}
	for i, nList := range numsLists {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(nList), " ", ", ", -1))

		if find3Sum(nList, tList[i]) {
			fmt.Printf("   Sum for %d exists\n", tList[i])
		} else {
			fmt.Printf("   Sum for %d does not exist\n", tList[i])
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
