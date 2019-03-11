package main

import (
	"fmt"
	"math"

	"github.com/golang-collections/go-datastructures/queue"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Add(value int) {
	t.root = Add(t.root, value)
}

func Add(n *TreeNode, value int) *TreeNode {
	if n == nil {
		// Equivalent to return &Tree{value: value}.
		n = new(TreeNode)
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

func InOrder(n *TreeNode) {
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

func PreOrder(n *TreeNode) {
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

func PostOrder(n *TreeNode) {
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
func appendValues(values []int, t *TreeNode) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func (t *Tree) PrintBredthFirst() {
	que := new(queue.Queue)
	var temp *TreeNode

	if t.root != nil {
		que.Put(t.root)
	}

	for que.Empty() == false {
		temp2, _ := que.Get(1)
		temp = temp2[0].(*TreeNode)

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

func NthPreOrder(TreeNode *TreeNode, index int, counter *int) {
	if TreeNode != nil {
		(*counter)++
		if *counter == index {
			fmt.Println(TreeNode.value)
		}
		NthPreOrder(TreeNode.left, index, counter)
		NthPreOrder(TreeNode.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int = 0
	NthPostOrder(t.root, index, &counter)
}

func NthPostOrder(TreeNode *TreeNode, index int, counter *int) {
	if TreeNode != nil {
		NthPostOrder(TreeNode.left, index, counter)
		NthPostOrder(TreeNode.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println(TreeNode.value)
		}
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int = 0
	NthInOrder(t.root, index, &counter)
}

func NthInOrder(TreeNode *TreeNode, index int, counter *int) {
	if TreeNode != nil {
		NthInOrder(TreeNode.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(TreeNode.value)
		}
		NthInOrder(TreeNode.right, index, counter)
	}
}

func (t *Tree) Find(value int) bool {
	var curr *TreeNode = t.root

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
	var TreeNode *TreeNode = t.root
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for TreeNode.left != nil {
		TreeNode = TreeNode.left
	}
	return TreeNode.value, true
}

func (t *Tree) FindMax() (int, bool) {
	var TreeNode *TreeNode = t.root
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for TreeNode.right != nil {
		TreeNode = TreeNode.right
	}
	return TreeNode.value, true
}

func (t *Tree) FindMaxTreeNode() *TreeNode {
	var TreeNode *TreeNode = t.root
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for TreeNode.right != nil {
		TreeNode = TreeNode.right
	}
	return TreeNode
}

func (t *Tree) FindMinTreeNode() *TreeNode {
	var TreeNode *TreeNode = t.root
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for TreeNode.left != nil {
		TreeNode = TreeNode.left
	}
	return TreeNode
}

func FindMax(curr *TreeNode) *TreeNode {
	var TreeNode *TreeNode = curr
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for TreeNode.right != nil {
		TreeNode = TreeNode.right
	}
	return TreeNode
}

func FindMin(curr *TreeNode) *TreeNode {
	var TreeNode *TreeNode = curr
	if TreeNode == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for TreeNode.left != nil {
		TreeNode = TreeNode.left
	}
	return TreeNode
}

func (t *Tree) Free() {
	t.root = nil
}

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(TreeNode *TreeNode, value int) *TreeNode {
	if TreeNode != nil {
		if TreeNode.value == value {
			if TreeNode.left == nil && TreeNode.right == nil {
				return nil
			} else {
				if TreeNode.left == nil {
					return TreeNode.right
				}

				if TreeNode.right == nil {
					return TreeNode.left
				}

				maxTreeNode := FindMax(TreeNode.left)
				maxValue := maxTreeNode.value
				TreeNode.value = maxValue
				TreeNode.left = DeleteNode(TreeNode.left, maxValue)
			}
		} else {
			if TreeNode.value > value {
				TreeNode.left = DeleteNode(TreeNode.left, value)
			} else {
				TreeNode.right = DeleteNode(TreeNode.right, value)
			}
		}
	}
	return TreeNode
}

func (t *Tree) TreeDepth() int {
	return TreeDepth(t.root)
}

func TreeDepth(root *TreeNode) int {
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

func Identical(TreeNode1 *TreeNode, TreeNode2 *TreeNode) bool {
	if TreeNode1 == nil && TreeNode2 == nil {
		return true
	} else if TreeNode1 == nil || TreeNode2 == nil {
		return false
	} else {
		return ((TreeNode1.value == TreeNode2.value) &&
			Identical(TreeNode1.left, TreeNode2.left) &&
			Identical(TreeNode1.right, TreeNode2.right))
	}
}

func (t *Tree) Ancestor(first int, second int) *TreeNode {
	if first > second {
		temp := first
		first = second
		second = temp
	}
	return Ancestor(t.root, first, second)
}

func Ancestor(curr *TreeNode, first int, second int) *TreeNode {
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

func CopyTree(curr *TreeNode) *TreeNode {
	var temp *TreeNode
	if curr != nil {
		temp = new(TreeNode)
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

func CopyMirrorTree(curr *TreeNode) *TreeNode {
	var temp *TreeNode
	if curr != nil {
		temp = new(TreeNode)
		temp.value = curr.value
		temp.right = CopyMirrorTree(curr.left)
		temp.left = CopyMirrorTree(curr.right)
		return temp
	} else {
		return nil
	}
}

func (t *Tree) NumTreeNodes() int {
	return NumTreeNodes(t.root)
}

func NumTreeNodes(curr *TreeNode) int {
	if curr == nil {
		return 0
	} else {
		return (1 + NumTreeNodes(curr.right) + NumTreeNodes(curr.left))
	}
}

func (t *Tree) NumFullTreeNodesBT() int {
	return NumTreeNodes(t.root)
}

func NumFullTreeNodesBT(curr *TreeNode) int {
	var count int
	if curr == nil {
		return 0
	}

	count = NumFullTreeNodesBT(curr.right) + NumFullTreeNodesBT(curr.left)
	if curr.right != nil && curr.left != nil {
		count++
	}

	return count
}

func (t *Tree) MaxLengthPathBT() int {
	return MaxLengthPathBT(t.root)
}

func MaxLengthPathBT(curr *TreeNode) int {
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

func (t *Tree) NumLeafTreeNodes() int {
	return NumLeafTreeNodes(t.root)
}

func NumLeafTreeNodes(curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	} else {
		return (NumLeafTreeNodes(curr.right) + NumLeafTreeNodes(curr.left))
	}
}

func (t *Tree) SumAllBT() int {
	return SumAllBT(t.root)
}

func SumAllBT(curr *TreeNode) int {
	var sum, leftSum, rightSum int
	if curr == nil {
		return 0
	}

	rightSum = SumAllBT(curr.right)
	leftSum = SumAllBT(curr.left)
	sum = rightSum + leftSum + curr.value
	return sum
}

func IsBST3(root *TreeNode) bool {
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

func IsBST(curr *TreeNode, min int, max int) bool {
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

func IsBST2(root *TreeNode, count *int) bool {
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

func PrintAllPath(curr *TreeNode, stk *Stack) {
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

func LCA(curr *TreeNode, first int, second int) *TreeNode {
	var left, right *TreeNode

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

func LcaBST(curr *TreeNode, first int, second int) (int, bool) {
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

func TrimOutsidedataRange(curr *TreeNode, min int, max int) *TreeNode {
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

func PrintIndataRange(root *TreeNode, min int, max int) {
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

func FindMaxBT(curr *TreeNode) int {
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

func SearchBT(root *TreeNode, value int) bool {
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

func CreateBinaryTreeUtil(arr []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	curr := new(TreeNode)
	curr.value = arr[mid]
	curr.left = CreateBinaryTreeUtil(arr, start, mid-1)
	curr.right = CreateBinaryTreeUtil(arr, mid+1, end)
	return curr
}

func main38() {
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
	//fmt.Println(t.FindMaxTreeNode())
	//fmt.Println(t.FindMin())
	//fmt.Println(t.FindMinTreeNode())
	//t.Free()
	//t.InOrder()
	//fmt.Println()
	t.PrintAllPath()
}
