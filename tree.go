package main

import (
	"errors"
	"github.com/rs/zerolog/log"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

// Node Scope
func (n *Node) Insert(value, data string) error {

	if n == nil {
		err := "Cannot insert a value into a nil tree"
		log.Print(err)
		return errors.New(err)
	}

	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s string) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		err := "replaceNode() not allowed on a nil node"
		log.Print(err)
		return errors.New(err)
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		err := "Value to be deleted does not exist in the tree"
		log.Print(err)
		return errors.New(err)
	}

	// Search the node to be deleted.
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value
		n.Data = replacement.Data

		return replacement.Delete(replacement.Value, replParent)
	}
}

// Tree scope
func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		err := "Cannot delete from an empty tree"
		log.Print(err)
		return errors.New(err)
	}

	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		log.Print(err)
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}
