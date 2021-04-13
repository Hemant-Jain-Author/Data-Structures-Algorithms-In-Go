package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	isLastChar bool
	child      [26](*TrieNode)
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(s string) {
	if s == "" {
		return
	}

	str := strings.ToLower(s)
	t.root = t.InsertUtil(t.root, str, 0)
}

func (t *Trie) InsertUtil(curr *TrieNode, str string, index int) *TrieNode {
	if curr == nil {
		curr = new(TrieNode)
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

func (t *Trie) RemoveUtil(curr *TrieNode, str string, index int) {
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

func (t *Trie) FindUtil(curr *TrieNode, str string, index int) bool {
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
	t.Insert("banana")
	t.Insert("apple")
	t.Insert("mango")
	fmt.Println("Apple Found :", t.Find("apple"))
	fmt.Println("Grapes Found :", t.Find("grapes"))
	fmt.Println("Banana Found :", t.Find("banana"))
}

/*
Apple Found : true
Grapes Found : false
Banana Found : true
*/