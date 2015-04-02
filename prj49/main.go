// prj49 project main.go
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := make([]int, 0, 1000)
	arr = MakePrimeArray(arr, 9999)
	prMap := make(map[string][]int)
	//fmt.Println(arr)
	for _, v := range arr {
		if v > 999 {
			pk := key(v)
			d, ok := prMap[pk]
			if ok {
				d = append(d, v)
				prMap[pk] = d
			} else {
				prMap[pk] = []int{v}
			}
		}
	}

	for _, v := range prMap {
		if len(v) >= 3 {
			for i := 0; i < len(v)-2; i++ {
				if v[i+1]-v[i] == v[i+2]-v[i+1] {
					fmt.Println(v[i:])
				}
			}
		}
	}
}

func key(d int) string {

	if d > 999 {
		retS := ""
		arr := make([]int, 0, 4)
		for d > 0 {
			arr = append(arr, d%10)
			d = d / 10
		}
		sort.Ints(arr)

		for _, v := range arr {
			retS += strconv.Itoa(v)
		}
		return retS
	} else {
		return "0"
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
