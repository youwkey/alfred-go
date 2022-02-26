// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

import (
	"fmt"
	"testing"
)

func TestItems_Length(t *testing.T) {
	t.Parallel()

	type test struct {
		in  Items
		out int
	}

	tests := []test{
		{in: make(Items, 0), out: 0},
		{in: Items{}, out: 0},
		{in: make(Items, 1), out: 1},
		{in: Items{&Item{}}, out: 1},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:Length", i), func(t *testing.T) {
			t.Parallel()
			if test.in.Length() != test.out {
				t.Errorf("#%d: got: %d want: %d", i, test.in.Length(), test.out)
			}
		})
	}
}

func TestItems_IsEmpty(t *testing.T) {
	t.Parallel()
	type Test struct {
		in  Items
		out bool
	}

	tests := []Test{
		{in: make(Items, 0), out: true},
		{in: Items{}, out: true},
		{in: make(Items, 1), out: false},
		{in: Items{&Item{}}, out: false},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:IsEmpty", i), func(t *testing.T) {
			t.Parallel()
			if test.in.IsEmpty() != test.out {
				t.Errorf("#%d: got: %t want: %t", i, test.in.IsEmpty(), test.out)
			}
		})
	}
}

func TestItems_Append(t *testing.T) {
	t.Parallel()

	type Test struct {
		in Items
	}

	tests := []Test{
		{in: make(Items, 0)},
		{in: make(Items, 0)},
		{in: make(Items, 1)},
		{in: make(Items, 1)},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:Append", i), func(t *testing.T) {
			t.Parallel()
			length := test.in.Length()
			test.in.Append(&Item{})
			if test.in.Length() != length+1 {
				t.Errorf("#%d: got: %d want: %d", i, test.in.Length(), length+1)
			}
			test.in.Append(make([]*Item, 2)...)
			if test.in.Length() != length+1+2 {
				t.Errorf("#%d: got: %d want: %d", i, test.in.Length(), length+1+2)
			}
		})
	}
}
