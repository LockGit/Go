/**
 * Created by lock
 */
package main

import (
	"fmt"
	"math"
)

func countingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	var max int = math.MinInt32
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}
	copy(a, r)
}

func main() {
	arr := []int{8, 1, 2, 9, 4, 1, 2, 12, 5, 12}
	countingSort(arr, len(arr))
	fmt.Println("optimize count sort arr after is:", arr)
}
