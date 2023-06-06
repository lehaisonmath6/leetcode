package main

import (
	"fmt"
	"sort"
)

func isDupTriplet(tripleA, tripleB []int) bool {
	mapA := make(map[int]int)
	for _, i := range tripleA {
		mapA[i]++
	}
	mapB := make(map[int]int)

	for _, i := range tripleB {
		mapB[i]++
	}

	for k, v := range mapA {
		if v != mapB[k] {
			return false
		}
	}

	return true
}

func threeSum(nums []int) [][]int {
	target := 0
	l := [][]int{}
	n := len(nums)
	sort.Ints(nums)

	// Remove the duplicates from the items array.
	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// Now take the two pointer approach to solve the linear array problem.
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// if sum  is greater than target means the larger
			// value(from right as items is sorted i.e, k at right)
			// is taken and it is not able to sum up to the target
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				l = append(l, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j-1] == nums[j] && j < k {
					j++
				}
			}

		}
	}

	return l
}

func main() {
	fmt.Println(threeSum([]int{-7, -1, -13, 2, 13, 2, 12, 3, -11, 3, 7, -15, 2, -9, -13, -13, 11, -10, 5, -13, 2, -12, 0, -8, 8, -1, 4, 10, -13, -5, -6, -4, 9, -12, 5, 8, 5, 3, -4, 9, 13, 10, 10, -8, -14, 4, -6, 5, 10, -15, -1, -3, 10, -15, -4, 3, -1, -15, -10, -6, -13, -9, 5, 11, -6, -13, -4, 14, -3, 8, 1, -4, -5, -12, 3, -11, 7, 13, 9, 2, 13, -7, 6, 0, -15, -13, -11, -8, 9, -14, 1, 11, -7, 13, 0, -6, -15, 11, -6, -2, 4, 2, 9, -15, 5, -11, -11, -11, -13, 5, 7, 7, 5, -10, -7, 6, -7, -11, 13, 9, -10, -9}))
}
