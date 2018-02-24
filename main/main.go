package main

import "fmt"

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func main() {
	fmt.Println(lis([]int{10, 22, 9, 33, 21, 50, 41, 60, 80}))
}

func lis(list []int) int {
	d := make([]int, len(list))
	d[0] = 1

	for i := 1; i < len(d); i++ {
		d[i] = 1
		for j := 0; j < i; j++ {
			if list[j] <= list[i] && d[j]+1 > d[i] {
				d[i] = d[j] + 1
			}
		}
	}

	lisLen := 1
	for _, v := range d {
		if v > lisLen {
			lisLen = v
		}
	}
	return lisLen
}
