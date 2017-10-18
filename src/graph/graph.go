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
		for _, w := range dg.adj[v] {
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
	s := strconv.Itoa(dg.v) + " vertices, " + strconv.Itoa(dg.e) + " edges\n"
	for i, a := range dg.adj {
		s += strconv.Itoa(i) + " -> "
		for _, w := range a {
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}

func NewGraphFromEdges(v int, edges [][]int) Digraph {
	dg := NewDigraph(v)
	for _, edge := range edges {
		dg.AddEdge(edge[0], edge[1])
	}
	return dg
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func BuildOrder(v int, dependents [][]int) (bool, []int) {
	//dependents [3,5] mean 5 dependent on 3,in graph 3->5
	dg := NewGraphFromEdges(v, dependents)
	result := make([]int, 0)
	for true {
		rdg := dg.Reverse()
		noDepNum := 0
		for i, a := range rdg.adj {
			if len(a) == 0 && !contains(result, i) { //rdg[i] no out edge,dg[i] no in edge
				result = append(result, i)
				noDepNum++
			}
		}
		if noDepNum == 0 {
			return false, result
		}
		if len(result) == dg.v {
			return true, result
		}
		for _, i := range result { //clear in result vertex out edge
			dg.adj[i] = make([]int, 0)
		}
	}
	return false, result
}

type SymbolDigraph struct {
	Digraph
	indexes map[string]int //vertex name to index
	names   map[int]string //vertex index to name
}

func NewSymbolDigraphFromFile(f, sp string) SymbolDigraph {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("fail to open")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	indexes := make(map[string]int)
	names := make(map[int]string)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		symbols := strings.Split(str, sp)
		for _, s := range symbols {
			if _, ok := indexes[s]; !ok {
				indexes[s] = len(indexes)
			}
		}
	}

	for k, v := range indexes {
		names[v] = k
	}

	dg := NewDigraph(len(indexes))
	file.Seek(0, 0) //rewind file pointer
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		symbols := strings.Split(str, sp)
		v := indexes[symbols[0]]
		for i := 1; i < len(symbols); i++ {
			dg.AddEdge(v, indexes[symbols[i]])
		}
	}

	return SymbolDigraph{dg, indexes, names}
}

func (sdg SymbolDigraph) String() string {
	s := strconv.Itoa(sdg.v) + " vertices, " + strconv.Itoa(sdg.e) + " edges\n"
	for i, a := range sdg.adj {
		s += sdg.names[i] + " -> "
		for _, w := range a {
			s += sdg.names[w] + " "
		}
		s += "\n"
	}
	return s
}

func (sdg SymbolDigraph) BuildOrder() ([]string, bool) {
	result := make([]string, 0)
	for true {
		rdg := sdg.Reverse()
		noDepNum := 0
		for i, a := range rdg.adj {
			//rdg[i] no out edge,dg[i] no in edge,i no depend that not in result
			if len(a) == 0 && !containsStr(result, sdg.names[i]) {
				result = append(result, sdg.names[i])
				noDepNum++
			}
		}
		if len(result) == sdg.v {
			return result, true
		}
		if noDepNum == 0 {
			return result, false
		}
		for _, i := range result { //clear in result vertex out edge
			sdg.adj[sdg.indexes[i]] = make([]int, 0)
		}
	}

	return result, false
}

func (sdg *SymbolDigraph) AddVertex(v string) {
	if _, ok := sdg.indexes[v]; ok {
		return
	}
	vNum := len(sdg.indexes)
	sdg.indexes[v] = vNum
	sdg.names[vNum] = v
	sdg.adj = append(sdg.adj, make([]int, 0))
	sdg.v++
}
