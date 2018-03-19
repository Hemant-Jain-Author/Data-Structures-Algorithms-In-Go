package main

import (
	"fmt"
	"github.com/golang-collections/go-datastructures/queue"
	"math"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(value int) {
	t.root = Add(t.root, value)
}

func Add(n *Node, value int) *Node {
	if n == nil {
		// Equivalent to return &Tree{value: value}.
		n = new(Node)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = Add(n.left, value)
	} else {
		n.right = Add(n.right, value)
	}
	return n
}

func (t *Tree) InOrder() {
	InOrder(t.root)
	fmt.Println()
}

func InOrder(n *Node) {
	if n == nil {
		return
	}
	InOrder(n.left)
	fmt.Print(n.value)
	InOrder(n.right)
}

func (t *Tree) PreOrder() {
	PreOrder(t.root)
	fmt.Println()
}

func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value)
	PreOrder(n.left)
	PreOrder(n.right)
}

func (t *Tree) PostOrder() {
	PostOrder(t.root)
	fmt.Println()
}

func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Print(n.value)
}

// Sort sorts values in place.
func Sort(values []int) {
	t := new(Tree)
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

func (t *Tree) PrintBredthFirst() {
	que := new(queue.Queue)
	var temp *Node

	if t.root != nil {
		que.Put(t.root)
	}

	for que.Empty() == false {
		temp2, _ := que.Get(1)
		temp = temp2[0].(*Node)

		fmt.Print(temp.value, " ")

		if temp.left != nil {
			que.Put(temp.left)
		}
		if temp.right != nil {
			que.Put(temp.right)
		}
	}
	fmt.Println()
}

func (t *Tree) NthPreOrder(index int) {
	var counter int = 0
	NthPreOrder(t.root, index, &counter)
}

func NthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
		NthPreOrder(node.left, index, counter)
		NthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int = 0
	NthPostOrder(t.root, index, &counter)
}

func NthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		NthPostOrder(node.left, index, counter)
		NthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int = 0
	NthInOrder(t.root, index, &counter)
}

func NthInOrder(node *Node, index int, counter *int) {
	if node != nil {
		NthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		NthInOrder(node.right, index, counter)
	}
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

func FindMax(curr *Node) *Node {
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

func FindMin(curr *Node) *Node {
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
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			} else {
				if node.left == nil {
					temp = node.right
					return temp
				}

				if node.right == nil {
					temp = node.left
					return temp
				}

				maxNode := FindMax(node.left)
				maxValue := maxNode.value
				node.value = maxValue
				node.left = DeleteNode(node.left, maxValue)
			}
		} else {
			if node.value > value {
				node.left = DeleteNode(node.left, value)
			} else {
				node.right = DeleteNode(node.right, value)
			}
		}
	}
	return node
}

func (t *Tree) TreeDepth() int {
	return TreeDepth(t.root)
}

func TreeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := TreeDepth(root.left)
	rDepth := TreeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

func (t *Tree) isEqual(t2 *Tree) bool {
	return Identical(t.root, t2.root)
}

func Identical(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			Identical(node1.left, node2.left) &&
			Identical(node1.right, node2.right))
	}
}

func (t *Tree) Ancestor(first int, second int) *Node {
	if first > second {
		temp := first
		first = second
		second = temp
	}
	return Ancestor(t.root, first, second)
}

func Ancestor(curr *Node, first int, second int) *Node {
	if curr == nil {
		return nil
	}
	if curr.value > first && curr.value > second {
		return Ancestor(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return Ancestor(curr.right, first, second)
	}
	return curr
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = CopyTree(t.root)
	return Tree2
}

func CopyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = CopyTree(curr.left)
		temp.right = CopyTree(curr.right)
		return temp
	} else {
		return nil
	}
}

func (t *Tree) CopyMirrorTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = CopyMirrorTree(t.root)
	return Tree2
}

func CopyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.right = CopyMirrorTree(curr.left)
		temp.left = CopyMirrorTree(curr.right)
		return temp
	} else {
		return nil
	}
}

func (t *Tree) NumNodes() int {
	return NumNodes(t.root)
}

func NumNodes(curr *Node) int {
	if curr == nil {
		return 0
	} else {
		return (1 + NumNodes(curr.right) + NumNodes(curr.left))
	}
}

func (t *Tree) NumFullNodesBT() int {
	return NumNodes(t.root)
}

func NumFullNodesBT(curr *Node) int {
	var count int
	if curr == nil {
		return 0
	}

	count = NumFullNodesBT(curr.right) + NumFullNodesBT(curr.left)
	if curr.right != nil && curr.left != nil {
		count++
	}

	return count
}

func (t *Tree) MaxLengthPathBT() int {
	return MaxLengthPathBT(t.root)
}

func MaxLengthPathBT(curr *Node) int {
	var max, leftPath, rightPath, leftMax, rightMax int

	if curr == nil {
		return 0
	}

	leftPath = TreeDepth(curr.left)
	rightPath = TreeDepth(curr.right)
	max = leftPath + rightPath + 1

	leftMax = MaxLengthPathBT(curr.left)
	rightMax = MaxLengthPathBT(curr.right)
	if leftMax > max {
		max = leftMax
	}

	if rightMax > max {
		max = rightMax
	}
	return max
}

func (t *Tree) NumLeafNodes() int {
	return NumLeafNodes(t.root)
}

func NumLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	} else {
		return (NumLeafNodes(curr.right) + NumLeafNodes(curr.left))
	}
}

func (t *Tree) SumAllBT() int {
	return SumAllBT(t.root)
}

func SumAllBT(curr *Node) int {
	var sum, leftSum, rightSum int
	if curr == nil {
		return 0
	}

	rightSum = SumAllBT(curr.right)
	leftSum = SumAllBT(curr.left)
	sum = rightSum + leftSum + curr.value
	return sum
}

func IsBST3(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && FindMax(root.left).value > root.value {
		return false
	}
	if root.right != nil && FindMin(root.right).value <= root.value {
		return false
	}
	return (IsBST3(root.left) && IsBST3(root.right))
}

func (t *Tree) IsBST() bool {
	return IsBST(t.root, math.MinInt32, math.MaxInt32)
}

func IsBST(curr *Node, min int, max int) bool {
	if curr == nil {
		return true
	}

	if curr.value < min || curr.value > max {
		return false
	}

	return IsBST(curr.left, min, curr.value) && IsBST(curr.right, curr.value, max)
}

func (t *Tree) IsBST2() bool {
	var c int
	return IsBST2(t.root, &c)
}

func IsBST2(root *Node, count *int) bool {
	var ret bool
	if root != nil {
		ret = IsBST2(root.left, count)
		if !ret {
			return false
		}

		if *count > root.value {
			return false
		}

		*count = root.value

		ret = IsBST2(root.right, count)
		if !ret {
			return false
		}
	}
	return true
}

type Stack struct {
	s []int
}

func (s *Stack) Push(value int) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (t *Tree) PrintAllPath() {
	stk := new(Stack)
	PrintAllPath(t.root, stk)
}

func PrintAllPath(curr *Node, stk *Stack) {
	if curr == nil {
		return
	}

	stk.Push(curr.value)
	if curr.left == nil && curr.right == nil {
		stk.Print()
		stk.Pop()
		return
	}

	PrintAllPath(curr.right, stk)
	PrintAllPath(curr.left, stk)
	stk.Pop()
}

func (t *Tree) LCA(first int, second int) (int, bool) {
	ans := LCA(t.root, first, second)
	if ans != nil {
		return ans.value, true
	} else {
		fmt.Println("NotFoundException")
		return 0, false
	}
}

func LCA(curr *Node, first int, second int) *Node {
	var left, right *Node

	if curr == nil {
		return nil
	}

	if curr.value == first || curr.value == second {
		return curr
	}

	left = LCA(curr.left, first, second)
	right = LCA(curr.right, first, second)

	if left != nil && right != nil {
		return curr
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
	if curr == nil {
		fmt.Println("NotFoundException")
		return 0, false
	}

	if curr.value > first && curr.value > second {
		return LcaBST(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return LcaBST(curr.right, first, second)
	}
	return curr.value, true
}

func (t *Tree) TrimOutsidedataRange(min int, max int) {
	t.root = TrimOutsidedataRange(t.root, min, max)
}

func TrimOutsidedataRange(curr *Node, min int, max int) *Node {
	if curr == nil {
		return nil
	}

	curr.left = TrimOutsidedataRange(curr.left, min, max)
	curr.right = TrimOutsidedataRange(curr.right, min, max)

	if curr.value < min {
		return curr.right
	}

	if curr.value > max {
		return curr.left
	}

	return curr
}

func (t *Tree) PrintIndataRange(min int, max int) {
	PrintIndataRange(t.root, min, max)
}

func PrintIndataRange(root *Node, min int, max int) {
	if root == nil {
		return
	}

	PrintIndataRange(root.left, min, max)

	if root.value >= min && root.value <= max {
		fmt.Print(root.value, " ")
	}

	PrintIndataRange(root.right, min, max)
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
	return FindMaxBT(t.root)
}

func FindMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}

	max := curr.value
	left := FindMaxBT(curr.left)
	right := FindMaxBT(curr.right)
	if left > max {
		max = left
	}

	if right > max {
		max = right
	}

	return max
}

func SearchBT(root *Node, value int) bool {
	var left, right bool

	if root == nil || root.value == value {
		return false
	}

	left = SearchBT(root.left, value)
	if left {
		return true
	}

	right = SearchBT(root.right, value)
	if right {
		return true
	}

	return false
}

func CreateBinaryTree(arr []int, size int) *Tree {
	t := new(Tree)
	t.root = CreateBinaryTreeUtil(arr, 0, size-1)
	return t
}

func CreateBinaryTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	curr := new(Node)
	curr.value = arr[mid]
	curr.left = CreateBinaryTreeUtil(arr, start, mid-1)
	curr.right = CreateBinaryTreeUtil(arr, mid+1, end)
	return curr
}

func main() {
	t := new(Tree)
	t.Add(2)
	t.Add(1)
	t.Add(3)
	t.Add(4)
	//t.InOrder()
	//t.PreOrder()
	//t.PostOrder()
	//lst := []int{2, 1, 3, 4}
	//Sort(lst)
	//fmt.Println(lst)
	//t.PrintBredthFirst()
	//t.NthPreOrder(2)
	//t.NthPostOrder(2)
	//t.NthInOrder(2)
	//fmt.Println(t.Find(10))
	//fmt.Println(t.Find(3))
	//fmt.Println(t.FindMax())
	//fmt.Println(t.FindMaxNode())
	//fmt.Println(t.FindMin())
	//fmt.Println(t.FindMinNode())
	//t.Free()
	//t.InOrder()
	//fmt.Println()
	t.PrintAllPath()
}