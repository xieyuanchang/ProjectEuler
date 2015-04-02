// prj53 project main.go
package main

import (
	"fmt"
)

func main() {

	count := 0
	for n := uint64(1); n <= 100; n++ {
		for r := uint64(1); r < n-1; r++ {
			result := C(n, r)
			if result >= 1000000 {
				fmt.Println("C(", n, ",", r, ") = ", result)
				count++
			}
		}
	}
	fmt.Println("count:", count)

	//fmt.Println(C(23, 10))
}

func C(n uint64, r uint64) uint64 {
	arr := make([]uint64, 0, 100)
	for a := n; a > r; a-- {
		arr = append(arr, a)
	}

	A_arr := A(n - r)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(A_arr); j++ {
			if arr[i] == A_arr[j] {
				arr[i] = 1
				A_arr[j] = 1
			}
		}
	}

	for {
		flgBreak := true
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(A_arr); j++ {
				if arr[i] == 1 || A_arr[j] == 1 {
					continue
				}
				if arr[i]%A_arr[j] == 0 {
					arr[i] = arr[i] / A_arr[j]
					A_arr[j] = 1
					flgBreak = false
				}
				if A_arr[j]%arr[i] == 0 {
					A_arr[j] = A_arr[j] / arr[i]
					arr[i] = 1
					flgBreak = false
				}
			}
		}
		if flgBreak {
			break
		}
	}

	result := uint64(1)

	for i := 0; i < len(arr); i++ {
		if arr[i] > 1 {
			result = result * arr[i]
		}

	}

	if result < 0 {
		fmt.Println(arr)
		fmt.Println(A_arr)
		fmt.Println(result)
	}

	for j := 0; j < len(A_arr); j++ {
		if A_arr[j] > 1 {
			result = result / A_arr[j]
		}
	}

	return result
}

func A(r uint64) []uint64 {
	arr := make([]uint64, 0, 100)
	for i := uint64(1); i <= r; i++ {
		arr = append(arr, i)
	}

	return arr
}
