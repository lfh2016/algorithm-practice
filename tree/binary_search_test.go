package tree

import (
	"fmt"
	"testing"
)
import (
	"log"
	"math/rand"
	"time"
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

func TestNode_RandomNode(t *testing.T) {
	var bt *node
	rand.Seed(time.Now().UnixNano())
	bt = bt.put(10, "")
	bt.put(5, "")
	bt.put(20, "")
	bt.put(8, "")
	bt.put(3, "")
	bt.put(15, "")
	fmt.Println(bt.RandomNode().key)

	bt = nil
	bt = bt.put(9, "")
	if bt.RandomNode().key != 9 {
		log.Fatal("should be 9")
	}
}

func TestNode_PathsWithNum(t *testing.T) {
	var bt *node
	rand.Seed(time.Now().UnixNano())
	bt = bt.put(10, "")
	bt.put(5, "")
	bt.put(11, "")
	bt.put(6, "")
	bt.put(3, "")
	bt.put(15, "")
	bt.put(2, "")
	bt.put(1, "")
	fmt.Println(bt.PathsWithNum(11))
}
