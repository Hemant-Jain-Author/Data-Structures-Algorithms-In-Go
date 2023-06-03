package main

import (
	"fmt"
	"math"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func CreateCompleteBinaryTree(arr []int) *Tree {
	tree := &Tree{}
	tree.root = createCompleteBinaryTree(arr, 0, len(arr))
	return tree
}

func createCompleteBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = createCompleteBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = createCompleteBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) Add(value int) {
	t.root = t.addUtil(t.root, value)
}

func (t *Tree) addUtil(n *Node, value int) *Node {
	if n == nil {
		return &Node{value, nil, nil}
	}
	if value < n.value {
		n.left = t.addUtil(n.left, value)
	} else {
		n.right = t.addUtil(n.right, value)
	}
	return n
}

func (t *Tree) PrintPreOrder() {
	fmt.Print("Pre Order:")
	t.printPreOrder(t.root)
	fmt.Println()
}

func (t *Tree) printPreOrder(n *Node) {
	if n != nil {
		fmt.Print(n.value, " ")
		t.printPreOrder(n.left)
		t.printPreOrder(n.right)
	}
}

func (t *Tree) PrintPostOrder() {
	fmt.Print("Post Order:")
	t.printPostOrder(t.root)
	fmt.Println()
}

func (t *Tree) printPostOrder(n *Node) {
	if n != nil {
		t.printPostOrder(n.left)
		t.printPostOrder(n.right)
		fmt.Print(n.value, " ")
	}
}

func (t *Tree) PrintInOrder() {
	fmt.Print("In Order:")
	t.printInOrder(t.root)
	fmt.Println()
}

func (t *Tree) printInOrder(n *Node) {
	if n != nil {
		t.printInOrder(n.left)
		fmt.Print(n.value, " ")
		t.printInOrder(n.right)
	}
}

func (t *Tree) NthPreOrder(index int) {
	var counter int
	t.nthPreOrder(t.root, index, &counter)
}

func (t *Tree) nthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Println("NthPreOrder at index :", index, "is :", node.value)
			return
		}
		t.nthPreOrder(node.left, index, counter)
		t.nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	t.nthPostOrder(t.root, index, &counter)
}

func (t *Tree) nthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		t.nthPostOrder(node.left, index, counter)
		t.nthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println("NthPostOrder at index ", index, "is :", node.value)
			return
		}
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	t.nthInOrder(t.root, index, &counter)
}

func (t *Tree) nthInOrder(node *Node, index int, counter *int) {
	if node != nil {
		t.nthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println("NthInOrder at index :", index, "is :", node.value)
			return
		}
		t.nthInOrder(node.right, index, counter)
	}
}

func (t *Tree) PrintBreadthFirst() {
	que := new(Queue)
	var temp *Node

	if t.root != nil {
		que.Add(t.root)
	}
	fmt.Print("Breadth First : ")
	for que.Len() != 0 {
		temp2 := que.Remove()
		temp = temp2.(*Node)

		fmt.Print(temp.value, " ")

		if temp.left != nil {
			que.Add(temp.left)
		}
		if temp.right != nil {
			que.Add(temp.right)
		}
	}
	fmt.Println()
}

func (t *Tree) PrintLevelOrderLineByLine() {
	que1 := new(Queue)
	que2 := new(Queue)
	var temp *Node

	if t.root != nil {
		que1.Add(t.root)
	}

	for que1.Len() != 0 || que2.Len() != 0 {
		for que1.Len() != 0 {
			temp2 := que1.Remove()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")

			if temp.left != nil {
				que2.Add(temp.left)
			}
			if temp.right != nil {
				que2.Add(temp.right)
			}
		}
		fmt.Println(" ")
		for que2.Len() != 0 {
			temp2 := que2.Remove()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Add(temp.left)
			}
			if temp.right != nil {
				que1.Add(temp.right)
			}
		}
		fmt.Println(" ")
	}
}

func (t *Tree) PrintLevelOrderLineByLine2() {
	que := new(Queue)
	var count int

	if t.root != nil {
		que.Add(t.root)
	}

	for que.Len() != 0 {
		count = que.Len()
		for count > 0 {
			temp := que.Remove().(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que.Add(temp.left)
			}
			if temp.right != nil {
				que.Add(temp.right)
			}
			count -= 1
		}
		fmt.Println(" ")
	}
}

func (t *Tree) PrintDepthFirst() {
	stk := new(Stack)

	if t.root != nil {
		stk.Push(t.root)
	}
	fmt.Print("Depth First : ")
	for stk.Len() != 0 {
		temp := stk.Pop().(*Node)
		fmt.Print(temp.value, " ")
		if temp.right != nil {
			stk.Push(temp.right)
		}
		if temp.left != nil {
			stk.Push(temp.left)
		}
	}
}

func (t *Tree) PrintSpiralTree() {
	stk1 := new(Stack)
	stk2 := new(Stack)
	var temp *Node

	if t.root != nil {
		stk1.Push(t.root)
	}

	for stk1.Len() != 0 || stk2.Len() != 0 {
		for stk1.Len() != 0 {
			temp = stk1.Pop().(*Node)
			fmt.Print(temp.value, " ")
			if temp.right != nil {
				stk2.Push(temp.right)
			}
			if temp.left != nil {
				stk2.Push(temp.left)
			}
		}
		for stk2.Len() != 0 {
			temp2 := stk2.Pop()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				stk1.Push(temp.left)
			}
			if temp.right != nil {
				stk1.Push(temp.right)
			}
		}
	}
	fmt.Println(" ")
}

func (t *Tree) Find(value int) bool {
	var curr *Node = t.root
	for curr != nil {
		if curr.value == value {
			return true
		} else if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *Tree) FindMin() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for node.left != nil {
		node = node.left
	}
	return node.value, true
}

func (t *Tree) FindMax() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for node.right != nil {
		node = node.right
	}
	return node.value, true
}

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *Tree) FindMinNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *Tree) findMax(curr *Node) *Node {
	var node *Node = curr
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *Tree) findMin(curr *Node) *Node {
	var node *Node = curr
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *Tree) Free() {
	t.root = nil
}

func (t *Tree) DeleteNode(value int) {
	t.root = t.deleteNode(t.root, value)
}

func (t *Tree) deleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}

			maxNode := t.findMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = t.deleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = t.deleteNode(node.left, value)
			} else {
				node.right = t.deleteNode(node.right, value)
			}
		}
	}
	return node
}

func (t *Tree) TreeDepth() int {
	return t.treeDepth(t.root)
}

func (t *Tree) treeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := t.treeDepth(root.left)
	rDepth := t.treeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return t.isEqual(t.root, t2.root)
}

func (t *Tree) isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			t.isEqual(node1.left, node2.left) &&
			t.isEqual(node1.right, node2.right))
	}
}

func (t *Tree) Ancestor(first int, second int) *Node {
	if first > second {
		first, second = second, first
	}
	return t.ancestor(t.root, first, second)
}

func (t *Tree) ancestor(curr *Node, first int, second int) *Node {
	if curr == nil {
		return nil
	}
	if curr.value > first && curr.value > second {
		return t.ancestor(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return t.ancestor(curr.right, first, second)
	}
	return curr
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := &Tree{}
	Tree2.root = t.copyTree(t.root)
	return Tree2
}

func (t *Tree) copyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = &Node{curr.value,
			t.copyTree(curr.left),
			t.copyTree(curr.right)}
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := &Tree{}
	tree.root = t.copyMirrorTree(t.root)
	return tree
}

func (t *Tree) copyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = &Node{curr.value,
			t.copyMirrorTree(curr.right),
			t.copyMirrorTree(curr.left)}
		return temp
	}
	return nil
}

func (t *Tree) NumNodes() int {
	return t.numNodes(t.root)
}

func (t *Tree) numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return 1 + t.numNodes(curr.right) + t.numNodes(curr.left)
}

func (t *Tree) NumFullNodesBT() int {
	return t.numFullNodesBT(t.root)
}

func (t *Tree) numFullNodesBT(curr *Node) int {
	var count int
	if curr == nil {
		return 0
	}

	count = t.numFullNodesBT(curr.right) + t.numFullNodesBT(curr.left)
	if curr.right != nil && curr.left != nil {
		count++
	}
	return count
}

func (t *Tree) MaxLengthPathBT() int {
	return t.maxLengthPathBT(t.root)
}

func (t *Tree) maxLengthPathBT(curr *Node) int {
	var max, leftPath, rightPath, leftMax, rightMax int

	if curr == nil {
		return 0
	}

	leftPath = t.treeDepth(curr.left)
	rightPath = t.treeDepth(curr.right)
	max = leftPath + rightPath + 1

	leftMax = t.maxLengthPathBT(curr.left)
	rightMax = t.maxLengthPathBT(curr.right)
	if leftMax > max {
		max = leftMax
	}

	if rightMax > max {
		max = rightMax
	}
	return max
}

func (t *Tree) NumLeafNodes() int {
	return t.numLeafNodes(t.root)
}

func (t *Tree) numLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return t.numLeafNodes(curr.right) + t.numLeafNodes(curr.left)
}

func (t *Tree) SumAllBT() int {
	return t.sumAllBT(t.root)
}

func (t *Tree) sumAllBT(curr *Node) int {
	if curr == nil {
		return 0
	}
	return t.sumAllBT(curr.right) + t.sumAllBT(curr.left) + curr.value
}

func (t *Tree) TreeToListRec() *Node {
	return t.treeToListRec(t.root)
}

func (t *Tree) treeToListRec(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	var Head, Tail, tempHead *Node

	if curr.left == nil && curr.right == nil {
		curr.left = curr
		curr.right = curr
		return curr
	}

	if curr.left != nil {
		Head = t.treeToListRec(curr.left)
		Tail = Head.left
		curr.left = Tail
		Tail.right = curr
	} else {
		Head = curr
	}

	if curr.right != nil {
		tempHead = t.treeToListRec(curr.right)
		Tail = tempHead.left
		curr.right = tempHead
		tempHead.left = curr
	} else {
		Tail = curr
	}

	Head.left = Tail
	Tail.right = Head
	return Head
}

func (root *Node) PrintDLL() {
	if root == nil {
		return
	}
	curr := root
	tail := curr.left
	fmt.Print(`DLL nodes are : `)
	for curr != tail {
		fmt.Print(curr.value, " ")
		curr = curr.right
	}
	fmt.Println(curr.value)
}

func (t *Tree) IsBST3() bool {
	return t.isBST3(t.root)
}

func (t *Tree) isBST3(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && t.findMax(root.left).value > root.value {
		return false
	}
	if root.right != nil && t.findMin(root.right).value <= root.value {
		return false
	}
	return t.isBST3(root.left) && t.isBST3(root.right)
}

func (t *Tree) IsBST() bool {
	return t.isBST(t.root, math.MinInt32, math.MaxInt32)
}

func (t *Tree) isBST(curr *Node, min int, max int) bool {
	if curr == nil {
		return true
	}
	if curr.value < min || curr.value > max {
		return false
	}
	return t.isBST(curr.left, min, curr.value) && t.isBST(curr.right, curr.value, max)
}

func (t *Tree) IsBST2() bool {
	var c int
	return t.isBST2(t.root, &c)
}

func (t *Tree) isBST2(root *Node, count *int) bool {
	var ret bool
	if root != nil {
		ret = t.isBST2(root.left, count)
		if !ret {
			return false
		}

		if *count > root.value {
			return false
		}

		*count = root.value

		ret = t.isBST2(root.right, count)
		if !ret {
			return false
		}
	}
	return true
}

func (t *Tree) IsCompleteTree() bool {
	return t.isCompleteTree(t.root)
}

func (t *Tree) isCompleteTree(root *Node) bool {
	que := new(Queue)
	var temp *Node
	var noChild = false

	if root != nil {
		que.Add(root)
	}

	for que.Len() != 0 {
		temp = que.Remove().(*Node)
		if temp.left != nil {
			if noChild == true {
				return false
			}
			que.Add(temp.left)
		} else {
			noChild = true
		}

		if temp.right != nil {
			if noChild == true {
				return false
			}
			que.Add(temp.right)
		} else {
			noChild = true
		}
	}
	return true
}

func (t *Tree) IsCompleteTree2() bool {
	count := t.NumNodes()
	return t.isCompleteTreeUtil(t.root, 0, count)
}

func (t *Tree) isCompleteTreeUtil(curr *Node, index int, count int) bool {

	if curr == nil {
		return true
	}
	if index > count {
		return false
	}

	return t.isCompleteTreeUtil(curr.left, index*2+1, count) && t.isCompleteTreeUtil(curr.right, index*2+2, count)
}

func (t *Tree) IsHeap() bool {
	parentValue := math.MinInt32
	return t.IsCompleteTree() && t.isHeapUtil(t.root, parentValue)
}

func (t *Tree) isHeapUtil(curr *Node, parentValue int) bool {
	if curr == nil {
		return true
	}
	if curr.value < parentValue {
		return false
	}
	return t.isHeapUtil(curr.left, curr.value) && t.isHeapUtil(curr.right, curr.value)
}

func (t *Tree) IsHeap2() bool {
	count := t.NumNodes()
	parentValue := math.MinInt32
	return t.isHeapUtil2(t.root, 0, count, parentValue)
}

func (t *Tree) isHeapUtil2(curr *Node, index int, count int, parentValue int) bool {
	if curr == nil {
		return true
	}
	if index > count {
		return false
	}
	if curr.value < parentValue {
		return false
	}
	return t.isHeapUtil2(curr.left, index*2+1, count, curr.value) &&
		t.isHeapUtil2(curr.right, index*2+2, count, curr.value)
}

func (t *Tree) PrintAllPath() {
	stk := new(Stack)
	t.printAllPath(t.root, stk)
}

func (t *Tree) printAllPath(curr *Node, stk *Stack) {
	if curr == nil {
		return
	}
	stk.Push(curr.value)
	if curr.left == nil && curr.right == nil {
		stk.Print()
		stk.Pop()
		return
	}

	t.printAllPath(curr.right, stk)
	t.printAllPath(curr.left, stk)
	stk.Pop()
}

func (t *Tree) LCA(first int, second int) (int, bool) {
	ans := t.LCAUtil(t.root, first, second)
	if ans != nil {
		return ans.value, true
	}
	fmt.Println("NotFoundException")
	return 0, false
}

func (t *Tree) LCAUtil(curr *Node, first int, second int) *Node {
	var left, right *Node

	if curr == nil {
		return nil
	}

	if curr.value == first || curr.value == second {
		return curr
	}

	left = t.LCAUtil(curr.left, first, second)
	right = t.LCAUtil(curr.right, first, second)

	if left != nil && right != nil {
		return curr
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	return t.LcaBSTUtil(t.root, first, second)
}

func (t *Tree) LcaBSTUtil(curr *Node, first int, second int) (int, bool) {
	if curr == nil {
		fmt.Println("NotFoundException")
		return 0, false
	}

	if curr.value > first && curr.value > second {
		return t.LcaBSTUtil(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return t.LcaBSTUtil(curr.right, first, second)
	}
	if t.Find(first) && t.Find(second) {
		return curr.value, true
	}
	return 0, false
}

func (t *Tree) TrimOutsidedataRange(min int, max int) {
	t.root = t.trimOutsidedataRange(t.root, min, max)
}

func (t *Tree) trimOutsidedataRange(curr *Node, min int, max int) *Node {
	if curr == nil {
		return nil
	}
	curr.left = t.trimOutsidedataRange(curr.left, min, max)
	curr.right = t.trimOutsidedataRange(curr.right, min, max)
	if curr.value < min {
		return curr.right
	}
	if curr.value > max {
		return curr.left
	}
	return curr
}

func (t *Tree) PrintDataInRange(min int, max int) {
	t.printDataInRange(t.root, min, max)
}

func (t *Tree) printDataInRange(root *Node, min int, max int) {
	if root == nil {
		return
	}
	t.printDataInRange(root.left, min, max)
	if root.value >= min && root.value <= max {
		fmt.Print(root.value, " ")
	}
	t.printDataInRange(root.right, min, max)
}

func (t *Tree) FloorBST(val int) int {
	curr := t.root
	floor := math.MaxInt32

	for curr != nil {
		if curr.value == val {
			floor = curr.value
			break
		} else if curr.value > val {
			curr = curr.left
		} else {
			floor = curr.value
			curr = curr.right
		}
	}
	return floor
}

func (t *Tree) CeilBST(val int) int {
	curr := t.root
	ceil := math.MinInt32

	for curr != nil {
		if curr.value == val {
			ceil = curr.value
			break
		} else if curr.value > val {
			ceil = curr.value
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return ceil
}

func (t *Tree) FindMaxBT() int {
	return t.findMaxBT(t.root)
}

func (t *Tree) findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}

	max := curr.value
	left := t.findMaxBT(curr.left)
	right := t.findMaxBT(curr.right)
	if left > max {
		max = left
	}

	if right > max {
		max = right
	}
	return max
}

func (t *Tree) SearchBT(value int) bool {
	return t.searchBT(t.root, value)
}

func (t *Tree) searchBT(root *Node, value int) bool {
	var left, right bool
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	}

	left = t.searchBT(root.left, value)
	if left {
		return true
	}

	right = t.searchBT(root.right, value)
	if right {
		return true
	}
	return false
}

func CreateBinarySearchTree(arr []int) *Tree {
	t := &Tree{}
	size := len(arr)
	t.root = createBinarySearchTreeUtil(arr, 0, size-1)
	return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &Node{arr[mid],
		createBinarySearchTreeUtil(arr, start, mid-1),
		createBinarySearchTreeUtil(arr, mid+1, end)}
}

func isBSTArray(preorder []int, size int) bool {
	stk := new(Stack)
	var value int
	root := -999999
	for i := 0; i < size; i++ {
		value = preorder[i]

		// If value of the right child is less than root.
		if value < root {
			return false
		}
		// First left child values will be popped
		// Last popped value will be the root.
		for stk.Len() > 0 && stk.Top().(int) < value {
			root = stk.Pop().(int)
		}
		// add current value to the stack.
		stk.Push(value)
	}
	return true
}

// Sort sorts values in place.
func Sort(values []int) {
	t := &Tree{}
	for _, v := range values {
		t.Add(v)
	}
	appendValues(values[:0], t.root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Node) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

/* Testing Code */
func main1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateCompleteBinaryTree(arr)
	t.PrintPreOrder()
	t.PrintPostOrder()
	t.PrintInOrder()
	t.PrintBreadthFirst()
	t.PrintDepthFirst()
	t.PrintLevelOrderLineByLine()
	t.PrintLevelOrderLineByLine2()
	t.PrintSpiralTree()
	t.NthPreOrder(5)
	t.NthPostOrder(5)
	t.NthInOrder(5)
	t.PrintAllPath()
}

/*
Pre Order:1 2 4 8 9 5 10 3 6 7
Post Order:8 9 4 10 5 2 6 7 3 1
In Order:8 4 9 2 10 5 1 6 3 7
Breadth First : 1 2 3 4 5 6 7 8 9 10
Depth First : 1 2 4 8 9 5 10 3 6 7

1
2 3
4 5 6 7
8 9 10
1
2 3
4 5 6 7
8 9 10

1 2 3 7 6 5 4 8 9 10

NthPreOrder at index : 5 is : 9
NthPostOrder at index  5 is : 5
NthInOrder at index : 5 is : 10
[1 3 7]
[1 3 6]
[1 2 5 10]
[1 2 4 9]
[1 2 4 8]
*/

func main2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateCompleteBinaryTree(arr)
	fmt.Println(t.NumNodes())
	fmt.Println(t.SumAllBT())
	fmt.Println(t.NumLeafNodes())
	fmt.Println(t.NumFullNodesBT())
	fmt.Println(t.SearchBT(9))
	fmt.Println(t.FindMaxBT())
	fmt.Println(t.TreeDepth())
	fmt.Println(t.MaxLengthPathBT())
}

/*
10
55
5
4
true
10
4
6
*/

func main3() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateCompleteBinaryTree(arr)
	t2 := t.CopyTree()
	t3 := t.CopyMirrorTree()
	t.PrintInOrder()
	t2.PrintInOrder()
	t3.PrintInOrder()
	fmt.Println(t.IsEqual(t2))
	fmt.Println(t.IsCompleteTree())
	fmt.Println(t.IsCompleteTree2())
	fmt.Println(t.IsHeap())
	fmt.Println(t.IsHeap2())
}

/*
In Order:8 4 9 2 10 5 1 6 3 7
In Order:8 4 9 2 10 5 1 6 3 7
In Order:7 3 6 1 5 10 2 9 4 8
true
true
true
true
true
*/

func main4() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateCompleteBinaryTree(arr)

	t.PrintInOrder()
	t3 := t.TreeToListRec()
	t3.PrintDLL()
}

/*
In Order:8 4 9 2 10 5 1 6 3 7
DLL nodes are : 8 4 9 2 10 5 1 6 3 7
*/

func main5() {
	t := &Tree{}
	t.Add(6)
	t.Add(4)
	t.Add(2)
	t.Add(5)
	t.Add(1)
	t.Add(3)
	t.Add(8)
	t.Add(7)
	t.Add(9)
	t.Add(10)

	t.PrintInOrder()
}

/*
In Order:1 2 3 4 5 6 7 8 9 10

*/

func main6() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	t.PrintInOrder()

	fmt.Println(t.Find(6))
	fmt.Println(t.FindMin())
	fmt.Println(t.FindMax())
	fmt.Println(t.IsBST())
	fmt.Println(t.IsBST2())
	fmt.Println(t.IsBST3())
}

/*
In Order:1 2 3 4 5 6 7 8 9 10
true
1 true
10 true
true
true
true
*/
func main7() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	t.PrintInOrder()
	t.DeleteNode(8)
	t.PrintInOrder()

	fmt.Println(t.LcaBST(3, 4))
	fmt.Println(t.LcaBST(1, 4))
	fmt.Println(t.LcaBST(10, 4))

	t.PrintDataInRange(3, 9)
	t.TrimOutsidedataRange(3, 9)
	t.PrintInOrder()

}

/*
In Order:1 2 3 4 5 6 7 8 9 10
In Order:1 2 3 4 5 6 7 9 10
3 true
2 true
5 true
3 4 5 6 7 9
In Order:3 4 5 6 7 9
*/

func main8() {
	arr3 := []int{5, 2, 4, 6, 9, 10}
	fmt.Println(isBSTArray(arr3, len(arr3)))

	arr2 := []int{5, 2, 6, 4, 7, 9, 10}
	fmt.Println(isBSTArray(arr2, len(arr2)))
}

/*
true
false
*/

func main9() {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	fmt.Println(t.CeilBST(5))
	fmt.Println(t.FloorBST(5))
}

/*
6
4
*/

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
}

type Stack struct {
	stk []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[:n-1]
	return value
}

func (s Stack) Top() interface{} {
	n := len(s.stk)
	return s.stk[n-1]
}

func (s Stack) Len() int {
	return len(s.stk)
}

func (s Stack) IsEmpty() bool {
	return len(s.stk) == 0
}

func (s Stack) Print() {
	fmt.Println(s.stk)
}

type Queue struct {
	que []interface{}
}

func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func (q Queue) Front() interface{} {
	return q.que[0]
}

func (q Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

func (q Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q Queue) Len() int {
	return len(q.que)
}

func (q Queue) Print() {
	fmt.Println(q.que)
}
