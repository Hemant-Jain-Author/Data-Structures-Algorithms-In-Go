package main

import (
	"fmt"
	"strings"
)

type Node struct {
	isLastChar bool
	child      [26](*Node)
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(s string) {
	if s == "" {
		return
	}

	str := strings.ToLower(s)
	t.root = t.InsertUtil(t.root, str, 0)
}

func (t *Trie) InsertUtil(curr *Node, str string, index int) *Node {
	if curr == nil {
		curr = new(Node)
	}
	if len(str) == index {
		curr.isLastChar = true
	} else {
		curr.child[str[index]-'a'] = t.InsertUtil(curr.child[str[index]-'a'], str, index+1)
	}
	return curr
}

func (t *Trie) Remove(s string) {
	if s == "" {
		return
	}

	str := strings.ToLower(s)
	t.RemoveUtil(t.root, str, 0)
}

func (t *Trie) RemoveUtil(curr *Node, str string, index int) {
	if curr == nil {
		return
	}

	if len(str) == index {
		if curr.isLastChar {
			curr.isLastChar = false
		}
		return
	}
	t.RemoveUtil(curr.child[str[index]-'a'], str, index+1)
}

func (t *Trie) Find(s string) bool {
	if s == "" {
		return false
	}
	str := strings.ToLower(s)
	return t.FindUtil(t.root, str, 0)
}

func (t *Trie) FindUtil(curr *Node, str string, index int) bool {
	if curr == nil {
		return false
	}
	if len(str) == index {
		return curr.isLastChar
	}
	return t.FindUtil(curr.child[str[index]-'a'], str, index+1)
}

func main() {
	t := new(Trie)
	a := "apple"
	b := "app"
	c := "appletree"
	d := "tree"

	t.Insert(a)
	t.Insert(d)

	fmt.Println(t.Find(a))
	fmt.Println(t.Find(b))
	fmt.Println(t.Find(c))
	fmt.Println(t.Find(d))
}
