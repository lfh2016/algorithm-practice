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
