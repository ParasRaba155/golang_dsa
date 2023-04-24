package search

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type arrays[k any] struct {
	arr  []k
	size int
}

func generateRandomIntArrayofSize(s int) []int {
	var result []int
	for i := 0; i < s; i++ {
		result = append(result, rand.Int())
	}
	return result
}

func BenchmarkLinearSearchWhenItemExist(b *testing.B) {
	noOfTest := 10
	for i := 0; i < noOfTest; i++ {
		size := rand.Intn(100000)
		b.Run(fmt.Sprintf("array of size : %d", size), func(b *testing.B) {
			b.StopTimer() // Pause the benchmark for setup
			arr := generateRandomIntArrayofSize(size)
			itemIndex := rand.Intn(len(arr))
			b.StartTimer()
			for j := 0; j < b.N; j++ {
				LinearSearch(arr, arr[itemIndex])
			}
		})
	}
}

func BenchmarkLinearSearchWhenItemDoesNotExist(b *testing.B) {
	noOfTest := 10
	for i := 0; i < noOfTest; i++ {
		size := rand.Intn(100000)
		b.Run(fmt.Sprintf("array of size : %d", size), func(b *testing.B) {
			b.StartTimer()
			arr := generateRandomIntArrayofSize(size)
			b.StartTimer()
			for j := 0; j < b.N; j++ {
				LinearSearch(arr, -5)
			}
		})
	}
}

func BenchmarkBinarySearchWhenItemExist(b *testing.B) {
	noOfTest := 10
	for i := 0; i < noOfTest; i++ {
		size := rand.Intn(100000)
		b.Run(fmt.Sprintf("array of size : %d", size), func(b *testing.B) {
			b.StopTimer() // Pause the benchmark for setup
			arr := generateRandomIntArrayofSize(size)
			sort.Ints(arr)
			itemIndex := rand.Intn(len(arr))
			b.StartTimer()
			for j := 0; j < b.N; j++ {
				BinarySearch(arr, arr[itemIndex])
			}
		})
	}
}

func BenchmarkBinarySearchWhenItemDoesNotExist(b *testing.B) {
	noOfTest := 10
	for i := 0; i < noOfTest; i++ {
		size := rand.Intn(100000)
		b.Run(fmt.Sprintf("array of size : %d", size), func(b *testing.B) {
			b.StartTimer()
			arr := generateRandomIntArrayofSize(size)
			sort.Ints(arr)
			b.StartTimer()
			for j := 0; j < b.N; j++ {
				BinarySearch(arr, -5)
			}
		})
	}
}
