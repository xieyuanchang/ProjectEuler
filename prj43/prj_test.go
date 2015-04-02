package main

import "testing"

func TestIsSpecil(t *testing.T) {
	arr_ok := []int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}
	arr_fail := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	if IsSpecil(arr_ok) != true {
		t.Fail()
	}

	if IsSpecil(arr_fail) != false {
		t.Fail()
	}
}

func TestToint(t *testing.T) {
	arr := []int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}
	if 1406357289 != toInt(arr) {
		t.Fail()
	}
}
