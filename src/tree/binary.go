package tree

type node struct {
	key         int
	value       string
	left, right *node
	n           int
}
