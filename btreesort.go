package Hal

type Node struct {
	left *Node
	right *Node
	v int
}


func (n *Node) insert(v int) {
	if n == nil {
		return
	} else if v <= n.v {
		if n.left == nil {
			n.left = &Node {
				left: nil,
				right: nil,
				v: v,
			}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &Node {
				left: nil,
				right: nil,
				v: v,
			}
		} else {
			n.right.insert(v)
		}
	}
}

type BTree struct {
	root *Node
}

func (t *BTree) insert(v int) *BTree{
	if t.root == nil {
		t.root = &Node {
			left: nil,
			right: nil,
			v: v,
		}
	} else {
		t.root.insert(v)
	}
	return t
}

func fetchInOrder(root *Node, a []int, i *int) {
	if root != nil {
		fetchInOrder(root.left, a, i)
		a[*i] = root.v
		*i += 1
		a[*i] = root.v
		fetchInOrder(root.right, a, i)
	}
}

func BTreeSort(a []int) []int {
	n := len(a)
	tree := &BTree{}

	for i := 1; i < n; i ++ {
		tree.insert(a[i])
	}

	i := 0
	fetchInOrder(tree.root, a, &i)
	// fmt.Println(a)
	return a
}
