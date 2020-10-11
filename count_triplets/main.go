// Copyright 2020 Vasyl Naumenko. All rights reserved.
package main

import (
	"fmt"
)

// one of the possible solutions for triplets count problem,
// more info https://www.hackerrank.com/challenges/count-triplets-1/problem
type Triplet interface {
	AddValue(value int64)
	Count() int64
}

type RFactor struct {
	R int64
}

func (r *RFactor) Mul(value int64) int64 {
	return r.R * value
}

// each node has index in nodes tree and value
// children are corresponds to r factor calculation
type Node struct {
	Index    int
	Children []*Node
	Value    int64
}

// nodes tree
type Nodes struct {
	Nodes      []*Node // nodes tree
	RFactor    RFactor // R factor to calculate triplets
	Cursor     *Node   // helper
	ChildIndex int     // helper
}

func (n *Nodes) Init(R int64) Triplet {
	nodes := &Nodes{}
	nodes.RFactor.R = R
	nodes.Nodes = []*Node{}

	return nodes
}

// adds value as node
// if value can be used in triplet, it will be added as child
func (n *Nodes) Insert(val int64) {
	if n.FindParent(val) {
		child := n.AddChild(val)
		n.SetChildIndex(child.Index)
	}
	if n.SetCursorToNext() {
		n.Insert(val)
	}
}

// sets cursor to first node
func (n *Nodes) ResetCursor() {
	n.Cursor = n.Nodes[0]
}

// clears stored child index
func (n *Nodes) ClearChildIndex() {
	n.ChildIndex = 0 // index 0 node can`t be a child, it always base node
}

// stores node index upon child creation
func (n *Nodes) SetChildIndex(i int) {
	n.ChildIndex = i
}

// inserts a new value to the nodes tree
func (n *Nodes) AddValue(val int64) {
	if len(n.Nodes) == 0 {
		n.Nodes = append(n.Nodes, &Node{
			Index: 0,
			Value: val,
		})
		return
	}

	n.ResetCursor()
	n.ClearChildIndex()
	n.Insert(val)
}

// moves Cursor to Parent
func (n *Nodes) FindParent(val int64) bool {
	if n.RFactor.Mul(n.Cursor.Value) == val {
		return true
	}

	if n.SetCursorToNext() {
		return n.FindParent(val)
	}

	return false
}

func (n *Nodes) SetCursorToNext() bool {
	if n.IsNextNode() {
		n.Cursor = n.Nodes[n.Cursor.Index+1]
		return true
	}

	return false
}

// adds node as child and sets
// sets child index
func (n *Nodes) AddChild(val int64) *Node {
	cur := *n.Cursor
	parent := n.Nodes[cur.Index]
	var child *Node

	// store child as parent for next Insert
	if n.ChildIndex == 0 {
		child = &Node{
			Index: len(n.Nodes),
			Value: val,
		}

		n.Nodes = append(n.Nodes, child)
	} else {
		child = n.Nodes[n.ChildIndex]
	}

	if parent.Children == nil {
		parent.Children = []*Node{}
	}

	parent.Children = append(parent.Children, n.Nodes[child.Index])
	return n.Nodes[child.Index]
}

// true if next node available
func (n *Nodes) IsNextNode() bool {
	return n.Cursor.Index+1 < len(n.Nodes)
}

// checks each node and counts triplets
func (n *Nodes) CountTriplets() int64 {
	c := int64(0)

	if len(n.Cursor.Children) > 0 {
		c += n.ChildHasTriplets(0)
	}

	if n.SetCursorToNext() {
		c += n.CountTriplets()
	}

	return c
}

// checks each child of selected node and count triplets
func (n *Nodes) ChildHasTriplets(i int) int64 {
	c := int64(0)
	ch := n.Cursor.Children[i]
	if len(ch.Children) > 0 {
		c += int64(len(ch.Children))
	}

	i++
	if len(n.Cursor.Children) > i {
		c += n.ChildHasTriplets(i)
	}

	return c
}

// start count triplets from first node
func (n *Nodes) Count() int64 {
	n.ResetCursor()

	return n.CountTriplets()
}

// outputs Nodes structure
func (n *Nodes) String() string {
	var list string
	for _, node := range n.Nodes {
		list += fmt.Sprintf("👩nodes[%v]:%d", node.Index, node.Value)

		if node.Children != nil {
			list += fmt.Sprintf("\n\t📂children: ")

			for _, c := range node.Children {
				list += fmt.Sprintf("(👤nodes[%v]:%d)", c.Index, c.Value)
			}
			list += fmt.Sprintf("\n")
		}
	}

	return fmt.Sprintf("Nodes: \n%s", list)
}

// build Nodes tree and count Triplets
func countTriplets(arr []int64, r int64) int64 {
	nodes := Nodes{}
	triplets := nodes.Init(r)
	for _, val := range arr {
		triplets.AddValue(val)
	}

	fmt.Println(triplets)
	return triplets.Count()
}

func main() {
	var arr []int64
	var r, ans int64

	// helper log
	log := func(a, b, c interface{}) {
		fmt.Printf("\n📜arr is %v \n🔺Triplets: %v should be %v\n", a, b, c)
		fmt.Printf("---\n\n")
	}

	arr = []int64{1, 2, 2, 4}
	r = int64(2)
	ans = countTriplets(arr, r)
	log(arr, ans, 2)

	arr = []int64{1, 3, 9, 9, 27, 81}
	r = int64(3)
	ans = countTriplets(arr, r)
	log(arr, ans, 6)

	arr = []int64{1, 5, 5, 25, 125}
	r = int64(5)
	ans = countTriplets(arr, r)
	log(arr, ans, 4)
}
