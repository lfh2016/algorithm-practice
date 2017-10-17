package main

import "fmt"

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func main() {
	x := string(42)
	fmt.Println(x)
}
