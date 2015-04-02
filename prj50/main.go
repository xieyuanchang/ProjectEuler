// prj50 project main.go
package main

import (
	"fmt"
	"math"
)

const (
	MAX_SIZE = 10000
)

var result int
var max_i int

func main() {
	arr := make([]int, 0, MAX_SIZE)
	arr = MakePrimeArray(arr, MAX_SIZE)
	fmt.Println(arr)

	for i := 1; i < len(arr) && len(arr)-1 > max_i; i++ {
		startWith(arr, i)
	}

}

func startWith(arr []int, start int) {
	sum := 0
	max := 0
	for i := start; i < len(arr); i++ {
		sum += arr[i]
		max += 1
		if sum > MAX_SIZE {
			break
		}
		if IsPrime(arr, sum) {
			if max_i < max {
				fmt.Println(sum, max)
				result = sum
				max_i = max
			}
		}
	}
}

func IsPrime(arr []int, d int) bool {
	arr_len := len(arr)
	for i := arr_len - 1; arr[i] >= d; i-- {
		if arr[i] == d {
			return true
		}
	}
	return false
}

func MakePrimeArray(arr []int, End int) []int {
	arr_len := len(arr)

	if len(arr) == 0 {
		arr = append(arr, 2)
		arr_len = 1
	}
	if arr[arr_len-1] <= End {
		for i := arr[arr_len-1] + 1; i <= End; i++ {
			isPrime := true
			sqr_i := int(math.Sqrt(float64(i)))
			for _, v := range arr {
				if v <= sqr_i {
					if i%v == 0 {
						isPrime = false
						break
					}
				} else {
					break
				}
			}
			if isPrime {
				arr = append(arr, i)
				if i > MAX_SIZE {
					break
				}
			}
		}
	}

	return arr
}
