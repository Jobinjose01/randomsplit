package common

import (
	"math/rand"
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
