// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package foo implements a set of simple mathematical functions.These comments are for
demonstration purpose only.Nothing more.

If you have any question, please don't hesitate to add yourself to
golang-nuts@googlegroups.com

You can also visit golang.org for full Go documentation.
*/
package foo

// Foo compares the two input values and returns the larger
// value. If the two values are equal , returns 0.
func Foo(a, b int) (ret int, err error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
	return 0, nil
}

// BUG(jack): #1 : I'm sorry but this code has an issue to be solved.
// BUG(tom): #2 : An issue assigned to another person.
