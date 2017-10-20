package tree

func (root *node) get(key int) (string, bool) {
	if root == nil {
		return "", false
	}
	if root.key == key {
		return root.value, true
	} else if key < root.key {
		return root.left.get(key)
	} else {
		return root.right.get(key)
	}
}

func (root *node) put(key int, v string) *node {
	if root == nil {
		return &node{key: key, value: v, n: 1}
	}
	if key == root.key {
		root.value = v
	} else if key < root.key {
		root.left = root.left.put(key, v)
	} else {
		root.right = root.right.put(key, v)
	}
	root.n = root.left.size() + root.right.size() + 1
	return root
}

func (root *node) size() int {
	if root == nil {
		return 0
	}
	return root.n
}

//all sequences of node array that can generate this tree
//get left and right child seq,weave every pair of them,prepend prefix
func (root *node) Sequences() [][]*node {
	result := make([][]*node, 0)
	if root == nil {
		result = append(result, make([]*node, 0))
		return result
	}
	prefix := make([]*node, 1)
	prefix[0] = root

	leftSeqes := root.left.Sequences()
	rightSeqes := root.right.Sequences()
	for _, leftSeq := range leftSeqes {
		for _, rightSeq := range rightSeqes {
			weaved := make([][]*node, 0)
			weave(leftSeq, rightSeq, prefix, &weaved)
			result = append(result, weaved...)
		}
	}
	return result
}

func weave(first, second, prefix []*node, results *[][]*node) {
	if len(first) == 0 || len(second) == 0 {
		result := append([]*node(nil), prefix...)
		result = append(result, first...)
		result = append(result, second...)
		*results = append(*results, result)
		return
	}

	p := append(prefix, first[0])
	firstClone := append([]*node(nil), first...)
	firstClone = append(firstClone[:0], firstClone[1:]...)
	weave(firstClone, second, p, results)

	p = append(prefix, second[0])
	secondClone := append([]*node(nil), second...)
	secondClone = append(secondClone[:0], secondClone[1:]...)
	weave(first, secondClone, p, results)

}
