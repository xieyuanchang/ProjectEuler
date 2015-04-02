// prj46 project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := make([]int, 0, 100)
	for i := 9; true; i += 2 {
		arr = MakePrimeArray(arr, i)
		if IsPrime(arr, i) {
			continue
		} else {
			flg := false
			for _, v := range arr {
				if IsAccept(i, v) {
					flg = true
					break
				}
			}

			if !flg {
				fmt.Println(i)
				break
			}

		}
	}
}

func IsAccept(d int, prime int) bool {
	d -= prime
	d_sqr := math.Sqrt(float64(d) / 2)
	return d_sqr-math.Trunc(d_sqr) == 0
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
			for _, v := range arr {
				if i%v == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				arr = append(arr, i)
			}
		}
	}

	return arr
}
