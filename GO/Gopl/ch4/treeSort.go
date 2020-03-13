package main

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values{
		root = Add(root, v)
	}

	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}

	if t.value > v {
		t.left = Add(t.left, v)
	}else {
		t.right = Add(t.right, v)
	}

	return t
}