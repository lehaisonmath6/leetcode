package main

import (
	"fmt"
	"sort"
)

type NumFreq struct {
	Value int
	Freq  int
}

type NumFreqArr []NumFreq

func (a NumFreqArr) Len() int           { return len(a) }
func (a NumFreqArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NumFreqArr) Less(i, j int) bool { return a[i].Freq < a[j].Freq }

func topKFrequent(nums []int, k int) []int {
	mapFreq := make(map[int]int)
	unSortedArr := NumFreqArr{}
	for _, i := range nums {
		mapFreq[i]++
	}
	for k, v := range mapFreq {
		unSortedArr = append(unSortedArr, NumFreq{
			Value: k,
			Freq:  v,
		})
	}
	sort.Sort(unSortedArr)
	rs := []int{}
	for i := 0; i < k; i++ {
		rs = append(rs, unSortedArr[len(unSortedArr)-1-i].Value)
	}
	return rs
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 5}, 3))
}
