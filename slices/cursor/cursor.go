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
var ErrUnderflow = errors.New("underflow")
var ErrOverflow = errors.New("overflow")

type Cursor[T any] struct {
	buff   []T
	pos    int
	cap    int
	lessFn func(i, j int) bool
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

type hasLess interface {
	Less(i, j int) bool
}

func (c *Cursor[T]) Less(i, j int) bool {
	if c.lessFn != nil {
		return c.lessFn(i, j)
	}

	var cc T
	v := any(cc)
	if h, ok := v.(hasLess); ok {
		return h.Less(i, j)
	}

	return false
}

func (c *Cursor[T]) Swap(i, j int) {
	if c.isValidPOS(i) && c.isValidPOS(j) {
		c.buff[i], c.buff[j] = c.buff[j], c.buff[i]
	}
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

func (c *Cursor[T]) Skip(i int) (*Cursor[T], error) {
	if c.isValidPOS(c.pos + i) {
		return New(c.buff[c.pos+i:]), nil
	}

	return New[T](nil), ErrIndexOutOfRange
}

// Take takes the next X from the cursor if they exist, or returns an error
// if there are not enough elements
func (c *Cursor[T]) Take(i int) ([]T, *Cursor[T], error) {
	if i < 0 {
		return nil, c, ErrIndexOutOfRange
	}

	if c.pos+i > len(c.buff) {
		return nil, c, ErrUnderflow
	}

	out := make([]T, i)
	copy(out, c.buff[c.pos:c.pos+i])

	return out, New(c.buff[c.pos+i:]), nil
}

func (c *Cursor[T]) Copy() *Cursor[T] {
	out := New(c.buff)
	out.pos = c.pos
	return out
}

// Replace replaces the next X elements with the given values
// and returns a new cursor with the updated buffer and the current
// cursor position
func (c *Cursor[T]) Replace(values ...T) (*Cursor[T], error) {
	return c.ReplaceAt(c.pos, values...)
}

func (c *Cursor[T]) ReplaceAt(pos int, values ...T) (*Cursor[T], error) {
	if !c.isValidPOS(pos) {
		return nil, ErrIndexOutOfRange
	}

	if pos+len(values) > len(c.buff) {
		return nil, ErrOverflow
	}

	out := New(append(
		c.buff[:pos],
		append(
			values,
			c.buff[pos+len(values):]...,
		)...,
	))

	out.pos = c.pos

	return out, nil
}

func (c *Cursor[T]) Delete() {
	c.DeleteAt(c.pos)
}

func (c *Cursor[T]) DeleteAt(pos int) {
	if c.isValidPOS(pos) {
		c.buff = append(c.buff[:pos], c.buff[pos+1:]...)
	}
}

// Insert inserts the given values at the current position
// and returns an error if there is not enough space
// This function modifies the cursor rather than returning a new one
func (c *Cursor[T]) Insert(values ...T) error {
	return c.InsertAt(c.pos, values...)
}

func (c *Cursor[T]) InsertAt(pos int, values ...T) error {
	if !c.isValidPOS(pos) {
		return ErrIndexOutOfRange
	}

	if pos+len(values) > c.cap {
		return ErrOverflow
	}

	c.buff = append(
		c.buff[:pos],
		append(
			values,
			c.buff[pos:]...,
		)...,
	)

	return nil
}

func (c *Cursor[T]) Append(values ...T) {
	c.buff = append(c.buff, values...)
}

func (c *Cursor[T]) Prepend(values ...T) {
	c.buff = append(
		append([]T{}, values...),
		c.buff...,
	)

	// shift position
	c.pos += len(values)
}

type Option[T any] func(*Cursor[T])

func LessFn[T any](fn func(i, j int) bool) Option[T] {
	return func(c *Cursor[T]) {
		c.lessFn = fn
	}
}

func Cap[T any](capacity int) Option[T] {
	return func(c *Cursor[T]) {
		c.cap = capacity
	}
}

func New[T any](buff []T, opts ...Option[T]) *Cursor[T] {
	if buff == nil {
		buff = make([]T, 0)
	}

	copyBuff := make([]T, len(buff))
	copy(copyBuff, buff)

	out := &Cursor[T]{
		buff: copyBuff,
		pos:  0,

		// max value for int
		cap: int(^uint(0) >> 1),
	}

	for _, opt := range opts {
		opt(out)
	}

	return out
}
