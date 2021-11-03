package main

import "fmt"

type RBTree struct {
	root     *Node
	nullNode *Node
}

func NewRBTree() (self *RBTree) {
	self = &RBTree{}
	self.nullNode = NewNode(0, nil)
	self.nullNode.colour = false
	self.root = self.nullNode
	return
}

type Node struct {
	data                int
	colour              bool // true for red colour, false for black colour
	left, right, parent *Node
}

func NewNode(d int, nullNode *Node) (nd *Node) {
	nd = &Node{}
	nd.data = d
	nd.left = nullNode
	nd.right = nullNode
	nd.parent = nullNode
	nd.colour = true // New node are red in colour.
	return
}

// To check whether node is of colour red or not.
func (self *RBTree) isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return (node.colour == true)
}

func (self *RBTree) uncle(node *Node) *Node {
	// If no parent or grandparent, then no uncle
	if node.parent == self.nullNode || node.parent.parent == self.nullNode {
		return nil
	}
	if node.parent == node.parent.parent.left {
		// uncle on right
		return node.parent.parent.right
	} else {
		// uncle on left
		return node.parent.parent.left
	}
}

// Function to right rotate subtree rooted with x
func (self *RBTree) rightRotate(x *Node) *Node {
	y := x.left
	T := y.right

	// Rotation
	y.parent = x.parent
	y.right = x
	x.parent = y
	x.left = T
	if T != self.nullNode {
		T.parent = x
	}
	if x == self.root {
		self.root = y
		return y
	}
	if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
	}
	// Return new root
	return y
}

// Function to left rotate subtree rooted with x
func (self *RBTree) leftRotate(x *Node) *Node {
	y := x.right
	T := y.left

	// Rotation
	y.parent = x.parent
	y.left = x
	x.parent = y
	x.right = T
	if T != self.nullNode {
		T.parent = x
	}
	if x == self.root {
		self.root = y
		return y
	}
	if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
	}
	// Return new root
	return y
}

func (self *RBTree) rightLeftRotate(node *Node) *Node {
	node.right = self.rightRotate(node.right)
	return self.leftRotate(node)
}

func (self *RBTree) leftRightRotate(node *Node) *Node {
	node.left = self.leftRotate(node.left)
	return self.rightRotate(node)
}

func (self *RBTree) Find(data int) *Node {
	curr := self.root
	for curr != self.nullNode {
		if curr.data == data {
			return curr
		} else if curr.data > data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return nil
}

func (self *RBTree) PrintTree() {
	self.printTreeUtil(self.root, "", false)
	fmt.Println()
}

func (self *RBTree) printTreeUtil(node *Node, indent string, isLeft bool) {
	if node == self.nullNode {
		return
	}
	if isLeft {
		fmt.Print(indent, "L:")
		indent += "|  "
	} else {
		fmt.Print(indent, "R:")
		indent += "   "
	}
	fmt.Println(node.data, "(", node.colour, ")")
	self.printTreeUtil(node.left, indent, true)
	self.printTreeUtil(node.right, indent, false)
}

func (self *RBTree) Insert(data int) {
	self.root = self.insertUtil(self.root, data)
	temp := self.Find(data)
	self.fixRedRed(temp)
}

func (self *RBTree) insertUtil(node *Node, data int) *Node {
	if node == self.nullNode {
		node = NewNode(data, self.nullNode)
	} else if node.data > data {
		node.left = self.insertUtil(node.left, data)
		node.left.parent = node
	} else if node.data < data {
		node.right = self.insertUtil(node.right, data)
		node.right.parent = node
	}
	return node
}

func (self *RBTree) fixRedRed(x *Node) {
	// if x is root colour it black and return
	if x == self.root {
		x.colour = false
		return
	}
	if x.parent == self.nullNode || x.parent.parent == self.nullNode {
		return
	}
	// Initialize parent, grandparent, uncle
	parent := x.parent
	grandparent := parent.parent
	uncle := self.uncle(x)

	var mid *Node = nil
	if parent.colour == false {
		return
	}
	// parent colour is red. gp is black.
	if uncle != self.nullNode && uncle.colour == true {
		// uncle and parent is red.
		parent.colour = false
		uncle.colour = false
		grandparent.colour = true
		self.fixRedRed(grandparent)
		return
	}
	// parent is red, uncle is black and gp is black.
	// Perform LR, LL, RL, RR
	if parent == grandparent.left && x == parent.left {
		mid = self.rightRotate(grandparent)
	} else if parent == grandparent.left && x == parent.right {
		mid = self.leftRightRotate(grandparent)
	} else if parent == grandparent.right && x == parent.left {
		mid = self.rightLeftRotate(grandparent)
	} else if parent == grandparent.right && x == parent.right {
		mid = self.leftRotate(grandparent)
	}
	mid.colour = false
	mid.left.colour = true
	mid.right.colour = true
}

func (self *RBTree) Delete(data int) {
	self.deleteUtil(self.root, data)
}

func (self *RBTree) deleteUtil(node *Node, key int) {
	z := self.nullNode
	var x *Node
	var y *Node
	for node != self.nullNode {
		if node.data == key {
			z = node
			break
		} else if node.data <= key {
			node = node.right
		} else {
			node = node.left
		}
	}
	if z == self.nullNode {
		fmt.Println("Couldn't find key in the tree")
		return
	}
	y = z
	yColour := y.colour
	if z.left == self.nullNode {
		x = z.right
		self.joinParentChild(z, z.right)
	} else if z.right == self.nullNode {
		x = z.left
		self.joinParentChild(z, z.left)
	} else {
		y = self.minimum(z.right)
		yColour = y.colour
		z.data = y.data
		self.joinParentChild(y, y.right)
		x = y.right
	}
	if yColour == false {
		if x.colour == true {
			x.colour = false
			return
		} else {
			self.fixDoubleBlack(x)
		}
	}
}

func (self *RBTree) fixDoubleBlack(x *Node) {
	if x == self.root { // Root node.
		return
	}
	sib := self.sibling(x)
	parent := x.parent
	if sib == self.nullNode {
		// No sibling double black shifted to parent.
		self.fixDoubleBlack(parent)
	} else {
		if sib.colour == true {
			// Sibling colour is red.
			parent.colour = true
			sib.colour = false
			if sib.parent.left == sib {
				// Sibling is left child.
				self.rightRotate(parent)
			} else {
				// Sibling is right child.
				self.leftRotate(parent)
			}
			self.fixDoubleBlack(x)
		} else {
			// Sibling colour is black
			// At least one child is red.
			if sib.left.colour == true || sib.right.colour == true {
				if sib.parent.left == sib {
					// Sibling is left child.
					if sib.left != self.nullNode && sib.left.colour == true {
						// left left case.
						sib.left.colour = sib.colour
						sib.colour = parent.colour
						self.rightRotate(parent)
					} else {
						// left right case.
						sib.right.colour = parent.colour
						self.leftRotate(sib)
						self.rightRotate(parent)
					}
				} else {
					// Sibling is right child.
					if sib.left != self.nullNode && sib.left.colour == true {
						// right left case.
						sib.left.colour = parent.colour
						self.rightRotate(sib)
						self.leftRotate(parent)
					} else {
						// right right case.
						sib.right.colour = sib.colour
						sib.colour = parent.colour
						self.leftRotate(parent)
					}
				}
				parent.colour = false
			} else {
				// Both children black.
				sib.colour = true
				if parent.colour == false {
					self.fixDoubleBlack(parent)
				} else {
					parent.colour = false
				}
			}
		}
	}
}

func (self *RBTree) sibling(node *Node) *Node {
	// sibling null if no parent
	if node.parent == self.nullNode {
		return nil
	}
	if node.parent.left == node {
		return node.parent.right
	}
	return node.parent.left
}

func (self *RBTree) joinParentChild(u *Node, v *Node) {
	if u.parent == self.nullNode {
		self.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (self *RBTree) minimum(node *Node) *Node {
	for node.left != self.nullNode {
		node = node.left
	}
	return node
}

func main() {
	tree := NewRBTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(9)
	tree.PrintTree()
	tree.Delete(4)
	tree.PrintTree()
	tree.Delete(7)
	tree.PrintTree()
}

/*
R:4 ( false )
   L:2 ( true )
   |  L:1 ( false )
   |  R:3 ( false )
   R:6 ( true )
      L:5 ( false )
      R:8 ( false )
         L:7 ( true )
         R:9 ( true )

R:5 ( false )
   L:2 ( true )
   |  L:1 ( false )
   |  R:3 ( false )
   R:7 ( true )
      L:6 ( false )
      R:8 ( false )
         R:9 ( true )

R:5 ( false )
   L:2 ( true )
   |  L:1 ( false )
   |  R:3 ( false )
   R:8 ( true )
      L:6 ( false )
      R:9 ( false )
*/
