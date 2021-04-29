// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

type Items []*Item

func (i *Items) Length() int {
	return len(*i)
}

func (i *Items) IsEmpty() bool {
	return i.Length() == 0
}

func (i *Items) Append(items ...*Item) {
	*i = append(*i, items...)
}
