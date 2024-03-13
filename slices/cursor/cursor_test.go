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

package cursor

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_New(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})
	if c == nil {
		t.Fatal("expected a new cursor")
	}

	if c.Len() != 5 {
		t.Fatal("expected a cursor with length 5")
	}

	if c.pos != 0 {
		t.Fatal("expected a cursor with position 0")
	}

	if c.buff[0] != 1 {
		t.Fatal("expected a cursor with first element 1")
	}

	if c.buff[4] != 5 {
		t.Fatal("expected a cursor with last element 5")
	}
}

func Test_Cursor_Next(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 3, nil},
		{[]int{1, 2, 3, 4, 5}, 2, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 3, 5, nil},
		{[]int{1, 2, 3, 4, 5}, 4, 0, ErrIndexOutOfRange},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.pos = tt.pos
			got, err := c.Next()

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func Test_Cursor_Prev(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 3, 3, nil},
		{[]int{1, 2, 3, 4, 5}, 2, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 0, 0, ErrIndexOutOfRange},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.pos = tt.pos
			got, err := c.Prev()

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func Test_Cursor_Get(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 2, 3, nil},
		{[]int{1, 2, 3, 4, 5}, 3, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 4, 5, nil},
		{[]int{1, 2, 3, 4, 5}, 5, 0, ErrIndexOutOfRange},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.pos = tt.pos
			got, err := c.Get()

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func Test_Cursor_Seek(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 2, 3, nil},
		{[]int{1, 2, 3, 4, 5}, 3, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 4, 5, nil},
		{[]int{1, 2, 3, 4, 5}, 5, 0, ErrIndexOutOfRange},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			got, err := c.Seek(tt.pos)

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func Test_Cursor_Set(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		val  int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 10, []int{10, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, 10, []int{1, 10, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, 10, []int{1, 2, 10, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, 10, []int{1, 2, 3, 10, 5}},
		{[]int{1, 2, 3, 4, 5}, 4, 10, []int{1, 2, 3, 4, 10}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.pos = tt.pos
			c.Set(tt.val)

			for i, v := range c.buff {
				if tt.want[i] != v {
					t.Fatalf("expected %v, got %v", tt.want[i], v)
				}
			}
		})
	}
}

func Test_Cursor_Take(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		take int
		want []int
		rem  *Cursor[int]
		err  error
	}{
		{
			[]int{1, 2, 3, 4, 5}, 0, 3, []int{1, 2, 3},
			&Cursor[int]{
				buff: []int{4, 5},
				pos:  0,
			},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 5, []int{1, 2, 3, 4, 5},
			&Cursor[int]{
				buff: []int{},
				pos:  0,
			},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 6, []int{},
			&Cursor[int]{
				buff: []int{1, 2, 3, 4, 5},
				pos:  0,
			},
			ErrUnderflow,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 0, []int{},
			&Cursor[int]{
				buff: []int{1, 2, 3, 4, 5},
				pos:  0,
			},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, -1, []int{},
			&Cursor[int]{
				buff: []int{1, 2, 3, 4, 5},
				pos:  0,
			},
			ErrIndexOutOfRange,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 1, []int{1},
			&Cursor[int]{
				buff: []int{2, 3, 4, 5},
				pos:  0,
			},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 2, []int{1, 2},
			&Cursor[int]{
				buff: []int{3, 4, 5},
				pos:  0,
			},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, 4, []int{1, 2, 3, 4},
			&Cursor[int]{
				buff: []int{5},
				pos:  0,
			},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)

			taken, left, err := c.Take(tt.take)
			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			}

			for i, v := range taken {
				if tt.want[i] != v {
					t.Fatalf("expected %v, got %v", tt.want[i], v)
				}
			}

			for i, v := range left.buff {
				if tt.rem.buff[i] != v {
					t.Fatalf("expected %v, got %v", tt.rem.buff[i], v)
				}
			}
		})
	}
}

func Test_Cursor_Replace(t *testing.T) {
	tests := []struct {
		data   []int
		pos    int
		values []int
		want   []int
		err    error
	}{
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{10, 20, 30},
			[]int{10, 20, 30, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 1, []int{10, 20, 30},
			[]int{1, 10, 20, 30, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 2, []int{10, 20, 30},
			[]int{1, 2, 10, 20, 30},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 3, []int{10, 20, 30},
			nil,
			ErrOverflow,
		},
		{
			[]int{1, 2, 3, 4, 5}, 4, []int{10, 20, 30},
			[]int{1, 2, 3, 4, 10},
			ErrOverflow,
		},
		{
			[]int{1, 2, 3, 4, 5}, 5, []int{10, 20, 30},
			nil,
			ErrIndexOutOfRange,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, -1, []int{10, 20, 30},
			nil,
			ErrIndexOutOfRange,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{10},
			[]int{10, 2, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{10, 20},
			[]int{10, 20, 3, 4, 5},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.pos = tt.pos

			newC, err := c.Replace(tt.values...)

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			} else if err != nil {
				return
			}

			diff := cmp.Diff(newC.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}

			if newC.pos != c.pos {
				t.Fatalf("expected %v, got %v", c.pos, newC.pos)
			}

			if newC.Len() != c.Len() {
				t.Fatalf("expected %v, got %v", c.Len(), newC.Len())
			}
		})
	}
}

func Test_Cursor_ReplaceAt(t *testing.T) {
	tests := []struct {
		data   []int
		pos    int
		values []int
		want   []int
		err    error
	}{
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{10, 20, 30},
			[]int{10, 20, 30, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 1, []int{10, 20, 30},
			[]int{1, 10, 20, 30, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 2, []int{10, 20, 30},
			[]int{1, 2, 10, 20, 30},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 3, []int{10, 20, 30},
			nil,
			ErrOverflow,
		},
		{
			[]int{1, 2, 3, 4, 5}, 4, []int{10, 20, 30},
			[]int{1, 2, 3, 4, 10},
			ErrOverflow,
		},
		{
			[]int{1, 2, 3, 4, 5}, 5, []int{10, 20, 30},
			nil,
			ErrIndexOutOfRange,
		},
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, -1, []int{10, 20, 30},
			nil,
			ErrIndexOutOfRange,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)

			newC, err := c.ReplaceAt(tt.pos, tt.values...)

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			} else if err != nil {
				return
			}

			diff := cmp.Diff(newC.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}

			if newC.pos != c.pos {
				t.Fatalf("expected %v, got %v", c.pos, newC.pos)
			}

			if newC.Len() != c.Len() {
				t.Fatalf("expected %v, got %v", c.Len(), newC.Len())
			}
		})
	}
}

func Test_Cursor_Delete(t *testing.T) {
	tests := []struct {
		data []int
		pos  int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, 0,
			[]int{2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, 1,
			[]int{1, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, 2,
			[]int{1, 2, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, 3,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, 4,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5}, 5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, -1,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.DeleteAt(tt.pos)

			diff := cmp.Diff(c.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}

			if c.pos != 0 {
				t.Fatalf("expected %v, got %v", 0, c.pos)
			}

			if c.Len() != len(tt.want) {
				t.Fatalf("expected %v, got %v", len(tt.want), c.Len())
			}
		})
	}
}

func Test_Cursor_InsertAt(t *testing.T) {
	tests := []struct {
		data   []int
		pos    int
		values []int
		want   []int
		err    error
	}{
		{
			[]int{1, 2, 3, 4, 5}, 0, []int{10, 20, 30},
			[]int{10, 20, 30, 1, 2, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 1, []int{10, 20, 30},
			[]int{1, 10, 20, 30, 2, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 2, []int{10, 20, 30},
			[]int{1, 2, 10, 20, 30, 3, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 3, []int{10, 20, 30},
			[]int{1, 2, 3, 10, 20, 30, 4, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 4, []int{10, 20, 30},
			[]int{1, 2, 3, 4, 10, 20, 30, 5},
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5}, 5, []int{10, 20, 30},
			[]int{1, 2, 3, 4, 5, 10, 20, 30},
			ErrIndexOutOfRange,
		},
		{
			[]int{1, 2, 3, 4, 5}, 6, []int{10, 20, 30},
			nil,
			ErrIndexOutOfRange,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			err := c.InsertAt(tt.pos, tt.values...)

			if err != tt.err {
				t.Fatalf("expected %v, got %v", tt.err, err)
			} else if err != nil {
				return
			}

			diff := cmp.Diff(c.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func Test_Cursor_Append(t *testing.T) {
	tests := []struct {
		data   []int
		values []int
		want   []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, []int{10, 20, 30},
			[]int{1, 2, 3, 4, 5, 10, 20, 30},
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{10},
			[]int{1, 2, 3, 4, 5, 10},
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.Append(tt.values...)

			diff := cmp.Diff(c.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func Test_Cursor_Prepend(t *testing.T) {
	tests := []struct {
		data   []int
		values []int
		want   []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, []int{10, 20, 30},
			[]int{10, 20, 30, 1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{10},
			[]int{10, 1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			c := New(tt.data)
			c.Prepend(tt.values...)

			diff := cmp.Diff(c.buff, tt.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
