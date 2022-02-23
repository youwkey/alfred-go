// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

type Items []*Item

// Length returns the length of the Items.
func (i *Items) Length() int {
	return len(*i)
}

// IsEmpty reports whether Items are empty.
func (i *Items) IsEmpty() bool {
	return i.Length() == 0
}

// Append appends entries to Items.
func (i *Items) Append(items ...*Item) {
	*i = append(*i, items...)
}
