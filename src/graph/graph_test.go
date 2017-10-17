package graph

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	dg := NewDigraphFromFile("tinyDG.txt")
	fmt.Println(dg)
}
