// Copyright 2022 Developer Network
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nary provides a n-ary tree implementation.
package nary

import (
	"go.devnw.com/ds/trees"
)

// New creates a new n-ary tree with the given value as root.
func New[T any](v T) *Tree[T] {
	return &Tree[T]{root: &Node[T]{value: v}}
}

func NewFrom[T any](root *Node[T]) (*Tree[T], error) {
	if root == nil {
		return nil, trees.ErrNilRoot
	}

	return &Tree[T]{root: root}, nil
}

// Tree is a n-ary tree.
type Tree[T any] struct {
	root *Node[T]
}

// Node is a node in a n-ary tree.
type Node[T any] struct {
	value T

	parent   *Node[T]
	children []*Node[T]
}

// Value returns the value of the node.
func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Set(v T) {
	n.value = v
}

// Parent returns the parent of the node.
func (n *Node[T]) Parent() *Node[T] {
	return n.parent
}

// Children returns the children of the node.
func (n *Node[T]) Children() []*Node[T] {
	return n.children
}

// AddChild adds a child to the node.
func (n *Node[T]) AddChildren(c ...*Node[T]) {
	for _, child := range c {
		if child == nil {
			continue
		}

		child.parent = n
		n.children = append(n.children, child)
	}
}

// Root returns the root of the tree.
func (t *Tree[T]) Root() *Node[T] {
	return t.root
}

// Leaves returns the leaves of the tree.
func (t *Tree[T]) Leaves() []*Node[T] {
	var leaves []*Node[T]
	t.root.leaves(&leaves)
	return leaves
}

// leaves appends the leaves of the node to the given slice.
func (n *Node[T]) leaves(leaves *[]*Node[T]) {
	if len(n.children) == 0 {
		*leaves = append(*leaves, n)
		return
	}

	for _, c := range n.children {
		c.leaves(leaves)
	}
}
