// prj47 project main.go
package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 0, 100)
	flg := false
	count := 0
	for i := 2; true; i++ {
		arr = MakePrimeArray(arr, i)
		if ShowFactor(arr, i) == 4 {
			fmt.Println(i)
			if flg {
				count++
			} else {
				flg = true
				count = 1
			}
			if count == 4 {
				break
			}
		} else {
			count = 0
			flg = false
		}
	}
}

func ShowFactor(arr []int, d int) int {
	tmp := d
	preFactor := 0
	count := 0
	factor := make([]int, 0, 100)
	for _, v := range arr {
		for d%v == 0 {
			d = d / v
			factor = append(factor, v)
			if preFactor != v {
				count++
				preFactor = v
			}
		}
	}
	if count == 4 {
		fmt.Println(tmp, " = ", factor)
	}
	return count
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
