package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	half := len(arr) / 2

	first := mergeSort(arr[:half])
	second := mergeSort(arr[half:])

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	res := []int{}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}

	return res
}
