package main

import "fmt"

func merge(arr [][]int) [][]int{

	res := arr
	i := 0
	for {
		if i == len(res) - 1 {
			break
		}
		if res[i][1] >= res[i+1][0] {
			res[i][1] = res[i+1][1]
			if len(res)-1 == i {
				res = res[:len(res)-1]
			} else {
				res = append(res[:i+1], res[i+2:]...)
			}
		} else {
			i++
		}
	}
	return res
}

func main() {
	arr := [][]int{
		{1, 5},
		{3, 8},
		{4, 10},
		{10, 19},
	}

	res := merge(arr)
	fmt.Println(res)
}