package graph

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Digraph struct {
	v, e int
	adj  [][]int
}

func NewDigraph(v int) Digraph {
	adj := make([][]int, v)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	dg := Digraph{v: v, adj: adj}
	return dg
}

func (dg *Digraph) AddEdge(v, w int) {
	dg.adj[v] = append(dg.adj[v], w)
	dg.e++
}

func (dg *Digraph) Reverse() Digraph {
	rdg := NewDigraph(dg.v)
	for v := 0; v < dg.v; v++ {
		for w := range dg.adj[v] {
			rdg.AddEdge(w, v)
		}
	}
	return rdg
}

func NewDigraphFromFile(f string) Digraph {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("fail to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("fail to Atoi")
	}
	dg := NewDigraph(i)

	scanner.Scan() //skip edge num ,because Addedge will update it

	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		s := strings.Fields(str)
		v, err := strconv.Atoi(s[0])
		fmt.Println()
		if err != nil {
			log.Fatalf("err in atoi %s", s[0])
		}
		w, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("err in atoi %s", s[1])
		}
		dg.AddEdge(v, w)
	}
	return dg

}

func (dg Digraph) String() string {
	s := strconv.Itoa(dg.v) + " vertices, " + strconv.Itoa(dg.e) + " deges\n"
	for i, a := range dg.adj {
		s += strconv.Itoa(i) + " -> "
		for _, w := range a {
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}
