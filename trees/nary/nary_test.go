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

package nary

import (
	"reflect"
	"testing"

	"go.devnw.com/ds/trees"
)

func Test_New(t *testing.T) {
	tree := New(1)
	if tree.root.Value() != 1 {
		t.Fatalf("expected %v, got %v", 1, tree.root.Value())
	}

	if len(tree.root.Children()) != 0 {
		t.Fatalf("expected %v, got %v", 0, len(tree.root.Children()))
	}

	if tree.root.Parent() != nil {
		t.Fatalf("expected %v, got %v", nil, tree.root.Parent())
	}

	if len(tree.Leaves()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(tree.Leaves()))
	}

	if tree.Leaves()[0] != tree.root {
		t.Fatalf("expected %v, got %v", tree.root, tree.Leaves()[0])
	}

	if tree.Root() != tree.root {
		t.Fatalf("expected %v, got %v", tree.root, tree.Root())
	}

	tree.root.AddChildren(New(2).root)
	if len(tree.root.Children()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(tree.root.Children()))
	}

	if tree.root.Children()[0].Value() != 2 {
		t.Fatalf("expected %v, got %v", 2, tree.root.Children()[0].Value())
	}

	if tree.root.Children()[0].Parent() != tree.root {
		t.Fatalf("expected %v, got %v", tree.root, tree.root.Children()[0].Parent())
	}

	if len(tree.root.Children()[0].Children()) != 0 {
		t.Fatalf("expected %v, got %v", 0, len(tree.root.Children()[0].Children()))
	}

	if len(tree.Leaves()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(tree.Leaves()))
	}

	tree.root.AddChildren(New(3).root)
	if len(tree.root.Children()) != 2 {
		t.Fatalf("expected %v, got %v", 2, len(tree.root.Children()))
	}

	if tree.root.Children()[1].Value() != 3 {
		t.Fatalf("expected %v, got %v", 3, tree.root.Children()[1].Value())
	}

	if tree.root.Children()[1].Parent() != tree.root {
		t.Fatalf("expected %v, got %v", tree.root, tree.root.Children()[1].Parent())
	}

	if len(tree.root.Children()[1].Children()) != 0 {
		t.Fatalf("expected %v, got %v", 0, len(tree.root.Children()[1].Children()))
	}

	if len(tree.Leaves()) != 2 {
		t.Fatalf("expected %v, got %v", 2, len(tree.Leaves()))
	}

	tree.root.Children()[0].AddChildren(New(4).root)
	if len(tree.root.Children()[0].Children()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(tree.root.Children()[0].Children()))
	}

	if tree.root.Children()[0].Children()[0].Value() != 4 {
		t.Fatalf("expected %v, got %v", 4, tree.root.Children()[0].Children()[0].Value())
	}

	if tree.root.Children()[0].Children()[0].Parent() != tree.root.Children()[0] {
		t.Fatalf("expected %v, got %v", tree.root.Children()[0], tree.root.Children()[0].Children()[0].Parent())
	}

	if len(tree.root.Children()[0].Children()[0].Children()) != 0 {
		t.Fatalf("expected %v, got %v", 0, len(tree.root.Children()[0].Children()[0].Children()))
	}

	if len(tree.Leaves()) != 2 {
		t.Fatalf("expected %v, got %v", 2, len(tree.Leaves()))
	}

	tree.root.Children()[1].AddChildren(New(5).root)
	if len(tree.root.Children()[1].Children()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(tree.root.Children()[1].Children()))
	}

	if tree.root.Children()[1].Children()[0].Value() != 5 {
		t.Fatalf("expected %v, got %v", 5, tree.root.Children()[1].Children()[0].Value())
	}

	if tree.root.Children()[1].Children()[0].Parent() != tree.root.Children()[1] {
		t.Fatalf("expected %v, got %v", tree.root.Children()[1], tree.root.Children()[1].Children()[0].Parent())
	}

	if len(tree.root.Children()[1].Children()[0].Children()) != 0 {
		t.Fatalf("expected %v, got %v", 0, len(tree.root.Children()[1].Children()[0].Children()))
	}
}

func Test_NewFrom(t *testing.T) {
	tests := map[string]struct {
		root     *Node[int]
		expected *Tree[int]
		err      error
	}{
		"empty": {
			root:     nil,
			expected: nil,
			err:      trees.ErrNilRoot,
		},
		"1": {
			root:     &Node[int]{value: 1},
			expected: &Tree[int]{root: &Node[int]{value: 1}},
		},
		"1-2": {
			root:     &Node[int]{value: 1, children: []*Node[int]{{value: 2}}},
			expected: &Tree[int]{root: &Node[int]{value: 1, children: []*Node[int]{{value: 2}}}},
		},
		"1-2-3": {
			root:     &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3}}}}},
			expected: &Tree[int]{root: &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3}}}}}},
		},
		"1-2-3-4": {
			root:     &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3, children: []*Node[int]{{value: 4}}}}}}},
			expected: &Tree[int]{root: &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3, children: []*Node[int]{{value: 4}}}}}}}},
		},
		"1-2-3-4-5": {
			root:     &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3, children: []*Node[int]{{value: 4, children: []*Node[int]{{value: 5}}}}}}}}},
			expected: &Tree[int]{root: &Node[int]{value: 1, children: []*Node[int]{{value: 2, children: []*Node[int]{{value: 3, children: []*Node[int]{{value: 4, children: []*Node[int]{{value: 5}}}}}}}}}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tree, err := NewFrom(tc.root)
			if err != tc.err {
				t.Fatalf("expected %v, got %v", tc.err, err)
			}

			if !reflect.DeepEqual(tree, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, tree)
			}
		})
	}
}
