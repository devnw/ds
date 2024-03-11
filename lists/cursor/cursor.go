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
	"errors"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type Cursor[T any] struct {
	buff []T
	pos  int
}

// Slice returns a new slice with the elements from start to end
// If start or end are out of range, it returns an error
func (c *Cursor[T]) Slice(start, end int) ([]T, error) {
	if !c.isValidPOS(start) || !c.isValidPOS(end) {
		return nil, ErrIndexOutOfRange
	}

	// Create a new buffer and copy the slice
	out := make([]T, end-start)
	copy(out, c.buff[start:end])

	return out, nil
}

// Take returns a new cursor with the elements from start to end
// If start or end are out of range, it returns an error
func (c *Cursor[T]) Take(start, end int) (*Cursor[T], error) {
	slice, err := c.Slice(start, end)
	if err != nil {
		return nil, err
	}

	return New(slice), nil
}

// Chop removes the elements from start to end
// If start or end are out of range, it returns an error
func (c *Cursor[T]) Chop(start, end int) error {
	if !c.isValidPOS(start) || !c.isValidPOS(end) {
		return ErrIndexOutOfRange
	}

	c.buff = append(c.buff[:start], c.buff[end:]...)
	return nil
}

func (c *Cursor[T]) First() (T, error) {
	return c.Seek(0)
}

func (c *Cursor[T]) Last() (T, error) {
	return c.Seek(len(c.buff) - 1)
}

func (c *Cursor[T]) IterFn(f func(T) error) error {
	if !c.validPOS() {
		return ErrIndexOutOfRange
	}

	for c.pos < len(c.buff) {
		err := f(c.buff[c.pos])
		if err != nil {
			return err
		}

		if !c.isValidPOS(c.pos + 1) {
			break
		}

		c.pos++
	}

	return nil
}

func (c *Cursor[T]) Len() int {
	return len(c.buff)
}

func (c *Cursor[T]) Next() (T, error) {
	return c.Seek(c.pos + 1)
}

func (c *Cursor[T]) Prev() (T, error) {
	return c.Seek(c.pos - 1)
}

func (c *Cursor[T]) Get() (T, error) {
	return c.Seek(c.pos)
}

func (c *Cursor[T]) Seek(pos int) (T, error) {
	if c.isValidPOS(pos) {
		c.pos = pos
		return c.buff[c.pos], nil
	}

	var out T
	return out, ErrIndexOutOfRange
}

func (c *Cursor[T]) isValidPOS(pos int) bool {
	return pos >= 0 && pos < len(c.buff)
}

func (c *Cursor[T]) validPOS() bool {
	return c.isValidPOS(c.pos)
}

func (c *Cursor[T]) Set(v T) {
	if c.validPOS() {
		c.buff[c.pos] = v
	}
}

func (c *Cursor[T]) Delete() {
	if c.validPOS() {
		c.buff = append(c.buff[:c.pos], c.buff[c.pos+1:]...)
	}
}

func Append[T any](c *Cursor[T], v T) {
	c.buff = append(c.buff, v)
}

func Prepend[T any](c *Cursor[T], v T) {
	c.buff = append([]T{v}, c.buff...)
}

func New[T any](buff []T) *Cursor[T] {
	if buff == nil {
		buff = make([]T, 0)
	}

	return &Cursor[T]{buff: buff, pos: 0}
}
