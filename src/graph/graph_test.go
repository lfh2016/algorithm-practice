package graph

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	dg := NewDigraphFromFile("tinyDG.txt")
	fmt.Println(dg)
}

func TestBuildOrder(t *testing.T) {
	dependents := [][]int{
		{0, 3},
		{5, 1},
		{1, 3},
		{5, 0},
		{3, 2}}
	if ok, result := BuildOrder(6, dependents); ok {
		fmt.Println(result)
	}
}

func TestDigraphReverse(t *testing.T) {
	dg := NewDigraphFromFile("nanoDG.txt")
	fmt.Println(dg.Reverse())
}

func TestNewSymbolDigraphFromFile(t *testing.T) {
	sg := NewSymbolDigraphFromFile("dependence2.txt", ",")
	sg.AddVertex("e")
	fmt.Println(sg.BuildOrder())
}
