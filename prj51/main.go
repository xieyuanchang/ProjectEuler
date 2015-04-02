// prj51 project main.go
package main

import (
	"fmt"
	"math"
)

const (
	MAX_SIZE = 1000
)

var result int
var max_i int

func main() {
	arr := make([]int, 0, MAX_SIZE)
	arr = MakePrimeArray(arr, MAX_SIZE)

	for _, v := range arr {
		if v%10 == 7 {
			fmt.Println(v)
		}
	}

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
