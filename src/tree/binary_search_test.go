package tree

import (
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	var bt *node
	bt = bt.put(10, "")
	bt.put(5, "")
	bt.put(20, "")
	bt.put(8, "")
	bt.put(3, "")
	bt.put(15, "")
	seqs := bt.Sequences()
	for _, seq := range seqs {
		for _, node := range seq {
			fmt.Printf("%d,", node.key)
		}
		fmt.Println()
	}

}
