package other

import "fmt"

func Hannoi(from, to, assist string, n int) {
	if n == 1 {
		fmt.Printf("move disc from %s to %s\n", from, to)
		return
	}
	Hannoi(from, assist, to, n-1)
	Hannoi(from, to, assist, 1)
	Hannoi(assist, to, from, n-1)
}
