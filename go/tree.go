package main

// type TreeInterface interface {
// 	CountNodes() int
// 	PrintTree()
// 	Search(x Element) *Node
// 	InsertNode(x Element)
// 	DeleteNode(x Element) *Node
// 	RightSubTree(x Element) *Tree
// 	FindMostLeft() *Node
// 	DeleteAllNodes() *Node
// }

// type Element int

// type Node struct {
// 	data  Element
// 	left  *Node
// 	right *Node
// }

// type Tree struct {
// 	root *Node
// }

// func IsLeaf(node *Node) bool {
// 	if node == nil {
// 		return true
// 	}
// 	return node.left == nil && node.right == nil
// }

// func (n *Node) ChildTree() *Tree {
// 	if n == nil {
// 		return nil
// 	}
// 	return &Tree{root: n}
// }

// func (t *Tree) CountNodes() int {
// 	if t == nil || t.root == nil {
// 		return 0
// 	}
// 	if IsLeaf(t.root) {
// 		return 1
// 	}
// 	return 1 + t.root.left.ChildTree().CountNodes() + t.root.right.ChildTree().CountNodes()
// }

// func (t *Tree) PrintTree() {
// 	if t == nil || t.root == nil {
// 		fmt.Printf("nothing!")
// 		return
// 	}
// 	if t.root.left != nil {
// 		t.root.left.ChildTree().PrintTree()
// 	}
// 	fmt.Printf("%d ", t.root.data)
// 	if t.root.right != nil {
// 		t.root.right.ChildTree().PrintTree()
// 	}
// }

// func (t *Tree) Search(x Element) *Node {
// 	if t == nil {
// 		return nil
// 	}
// 	if t.root.data == x {
// 		return t.root
// 	}
// 	if t.root.data < x {
// 		return t.root.right.ChildTree().Search(x)
// 	} else {
// 		return t.root.left.ChildTree().Search(x)
// 	}
// }

// func (t *Tree) InsertNode(x Element) {
// 	node := &Node{data: x}
// 	if t.root == nil {
// 		// fmt.Println("root insert")
// 		t.root = node
// 	} else {
// 		// fmt.Println("not root insert!!!!!!!!")
// 		if t.root.data < x {
// 			if rightChild := t.root.right.ChildTree(); rightChild != nil {
// 				fmt.Println("rightChild.InsertNode(x)", x)
// 				rightChild.InsertNode(x)
// 			} else {
// 				fmt.Println("add root right", x)
// 				t.root.right = node
// 			}
// 		} else {
// 			if leftChild := t.root.left.ChildTree(); leftChild != nil {
// 				fmt.Println("leftChild.InsertNode(x)", x)
// 				leftChild.InsertNode(x)
// 			} else {
// 				fmt.Println("add root left", x)
// 				t.root.left = node
// 			}
// 		}
// 	}
// }

// func (t *Tree) DeleteNode(x Element) *Node {
// 	if t == nil {
// 		return nil
// 	}
// 	if t.root.data < x {
// 		t.root.right = t.root.right.ChildTree().DeleteNode(x)
// 		return t.root
// 	} else if t.root.data > x {
// 		t.root.left = t.root.left.ChildTree().DeleteNode(x)
// 		return t.root
// 	} else {
// 		if IsLeaf(t.root) {
// 			t.root = nil
// 		} else if t.root.left == nil && t.root.right != nil {
// 			t.root = t.root.right
// 		} else if t.root.left != nil && t.root.right == nil {
// 			t.root = t.root.left
// 		} else {
// 			rightMin := t.RightSubTree(x).FindMostLeft().data
// 			t.root.data = rightMin
// 			t.root.right = t.root.right.ChildTree().DeleteNode(rightMin)
// 		}
// 		return t.root
// 	}
// }

// func (t *Tree) RightSubTree(x Element) *Tree {
// 	if t == nil {
// 		return nil
// 	}
// 	if t.root.data == x {
// 		return t.root.right.ChildTree()
// 	} else if t.root.data < x {
// 		return t.root.right.ChildTree().RightSubTree(x)
// 	} else {
// 		return t.root.left.ChildTree().RightSubTree(x)
// 	}
// }

// func (t *Tree) FindMostLeft() *Node {
// 	if t == nil {
// 		return nil
// 	}
// 	if t.root.left == nil {
// 		return t.root
// 	}
// 	return t.root.left.ChildTree().FindMostLeft()
// }

// func (t *Tree) DeleteAllNodes() *Node {
// 	if t != nil {
// 		t.root.left = t.root.left.ChildTree().DeleteAllNodes()
// 		t.root.right = t.root.right.ChildTree().DeleteAllNodes()
// 		t.root = nil
// 		return t.root
// 	}
// 	return nil
// }

// func NewTree() TreeInterface {
// 	return &Tree{}
// }

// func main() {
// 	tree := NewTree()
// 	integers := []Element{5, 2, 3, 8, 9, 1, 4, 6, 7}
// 	//     5
// 	// 	2------8
// 	// 1  3  6  9
// 	//     4  7

// 	for _, v := range integers {
// 		tree.InsertNode(v)
// 	}

// 	fmt.Printf("count: %d\n", tree.CountNodes())
// 	tree.PrintTree()

// 	tree.DeleteNode(2)
// 	fmt.Printf("\nafter delete 2: ")
// 	tree.PrintTree()

// 	tree.DeleteAllNodes()
// 	fmt.Printf("\nafter delete all nodes: ")
// 	tree.PrintTree()
// }
