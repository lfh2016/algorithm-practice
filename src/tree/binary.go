package tree

type node struct {
	key         int
	value       string
	left, right *node
	n           int
}

func FirstCommonAncestor(root, p, q *node) *node {
	if !isChild(p, root) || isChild(q, root) { //p or q is child of root
		return nil
	}
	return firstCommonAncestor(root, p, q)
}

func isChild(n, root *node) bool {
	if root == nil {
		return false
	}
	if root == n {
		return true
	}
	return isChild(n, root.left) || isChild(n, root.right)
}

func firstCommonAncestor(root, p, q *node) *node {
	if p == root || q == root {
		return root
	}
	pInleft := isChild(p, root.left)
	qInleft := isChild(q, root.left)
	if pInleft != qInleft {
		return root
	}
	if pInleft {
		return firstCommonAncestor(root.left, p, q)
	} else {
		return firstCommonAncestor(root.right, p, q)
	}
}

func IsSubTree(t1, t2 *node) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t1.value == t2.value && matchTree(t1, t2) {
		return true
	}
	return IsSubTree(t1.left, t2) || IsSubTree(t1.right, t2)
}

func matchTree(t1, t2 *node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.value != t2.value {
		return false
	}
	return matchTree(t1.left, t2.left) && matchTree(t1.right, t2.right)
}

//The path does not need to start or end at the root or a leaf, but it must go downwards
//(traveling only from parent nodes to child nodes).
func (n *node) PathsWithNum(targetSum int) int {
	runningSumCounter := make(map[int]int)
	runningSumCounter[0] = 1
	return pathsWithNum(n, targetSum, 0, runningSumCounter)
}

func pathsWithNum(n *node, targetSum, runningSum int, runningSumCounter map[int]int) int {
	if n == nil {
		return 0
	}
	runningSum += n.key
	paths := runningSumCounter[runningSum-targetSum]

	runningSumCounter[runningSum] += 1
	paths += pathsWithNum(n.left, targetSum, runningSum, runningSumCounter)
	paths += pathsWithNum(n.right, targetSum, runningSum, runningSumCounter)
	runningSumCounter[runningSum] -= 1

	return paths
}
