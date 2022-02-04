package common

import (
	"math/rand"
	"sort"
	"time"
)

func ShuffleMe(a []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func SumSlice(result []int) int {
	var sum = 0
	for _, v := range result {
		sum += v
	}
	return sum
}

func DistributeVals(result []int, last_num int, split int) []int {

	sort.IntsAreSorted(result)

	largest_num := result[len(result)-1]

	diff := last_num - largest_num
	result[len(result)-1] = largest_num + 1
	diff = diff - 1

	if diff >= split && diff%split == 0 {
		extra := (diff / split)

		for key, val := range result {
			result[key] = int(val + extra)
		}
	}
	if diff >= split && diff%split != 0 {

		remainder := diff % split
		extra := (diff / split)

		for key, val := range result {
			result[key] = int(val+extra) + remainder
			remainder = 0
		}
	}

	if diff < split {
		result[split-1] = last_num
	}

	return result
}
