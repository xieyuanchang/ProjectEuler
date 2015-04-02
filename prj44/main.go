// prj44 project main.go
package main

import (
	"fmt"
)

func main() {
	maxLen := 10000
	arr := makePentagonal(maxLen)
	flg := false
	for i := 1; !flg && i < maxLen; i++ {
		flg = checkByStep(arr, i)
	}
	fmt.Println("over")
}

func checkByStep(arr []int, step int) bool {
	for i := 1; i+step < len(arr); i++ {
		if IsDiffIn(arr, i, i+step) && IsSumIn(arr, i, i+step) {
			fmt.Println(arr[i], arr[i+step], arr[i+step]-arr[i])
			return true
		}
	}
	return false
}

func makePentagonal(n int) []int {
	arr := make([]int, n, n)
	for i, _ := range arr {
		arr[i] = ToPentagonal(i)
	}
	return arr
}

func IsDiffIn(arr []int, a int, b int) bool {
	for i := a + 1; i < b-1; i++ {
		if arr[i] == arr[b]-arr[a] {
			return true
		}
	}
	return false
}

func IsSumIn(arr []int, a int, b int) bool {
	sum := arr[b] + arr[a]
	for i := b + 1; true; i++ {
		v := ToPentagonal(i)
		if v == sum {
			return true
		} else if v > sum {
			break
		}
	}
	return false
}

func ToPentagonal(i int) int {
	return i * (3*i - 1) / 2
}
