// prj43 project main.go
package main

import (
	"fmt"
)

var sum int

func main() {
	arr := make([]int, 0)
	sum = 0
	Add(arr, 1)
	fmt.Println(sum)
}

func Add(arr []int, l int) {
	maxLen := 10
	if len(arr) == maxLen {
		if IsSpecil(arr) {
			fmt.Println(arr)
			sum += toInt(arr)
		}
	} else {
		for i := 0; i < maxLen; i++ {
			if IsContained(arr, i) {
				continue
			} else {
				tmp := append(arr, i)
				Add(tmp, l+1)
			}
		}
	}
}

func IsContained(arr []int, t int) bool {
	for _, v := range arr {
		if t == v {
			return true
		}
	}
	return false
}

func toInt(arr []int) int {
	d := 0
	for _, v := range arr {
		d = d*10 + v
	}
	return d
}

func IsSpecil(arr []int) bool {
	sa := []int{1, 2, 3, 5, 7, 11, 13, 17}
	if len(arr) == 10 {
		flg := true
		for i := 1; i < 8; i++ {
			d := arr[i]*100 + arr[i+1]*10 + arr[i+2]
			if d%sa[i] != 0 {
				flg = false
				break
			}
		}
		return flg
	}
	return false
}
