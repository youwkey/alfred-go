// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

import (
	"encoding/json"
	"fmt"
	"testing"
)

func ps(s string) *string {
	return &s
}

func pb(b bool) *bool {
	return &b
}

func pIconType(i IconType) *IconType {
	return &i
}

func pItemType(i ItemType) *ItemType {
	return &i
}

func testMarshalJSON(t *testing.T, n int, in interface{}, out string) {
	t.Helper()

	data, err := json.Marshal(in)

	if err != nil {
		t.Errorf("#%d: marshal error: %v", n, err)
	} else if string(data) != out {
		t.Errorf("#%d: got: %v want: %v", n, string(data), out)
	}
}

func TestIcon_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in1 *Icon
		in2 *Icon
		out string
	}

	tests := []marshalJSONTest{
		// Minimal
		{in1: &Icon{path: "./icon.png"}, in2: NewIcon("./icon.png"), out: `{"path":"./icon.png"}`},
		// With type fileicon
		{
			in1: &Icon{path: "./icon.png", typ: pIconType("fileicon")},
			in2: NewIconWithType("./icon.png", IconTypeFileIcon),
			out: `{"path":"./icon.png","type":"fileicon"}`,
		},
		// With type filetype
		{
			in1: &Icon{path: "./icon.png", typ: pIconType("filetype")},
			in2: NewIcon("./icon.png").Type(IconTypeFileType),
			out: `{"path":"./icon.png","type":"filetype"}`,
		},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalIcon", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in1, test.out)
			testMarshalJSON(t, i, test.in2, test.out)
		})
	}
}

func TestModifier_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in1 *Modifier
		in2 *Modifier
		out string
	}

	tests := []marshalJSONTest{
		// Set subtitle
		{
			in1: &Modifier{subtitle: ps("subtitle")},
			in2: NewModifier().Subtitle("subtitle"),
			out: `{"subtitle":"subtitle"}`,
		},
		// Set arg
		{in1: &Modifier{arg: ps("arg")}, in2: NewModifier().Arg("arg"), out: `{"arg":"arg"}`},
		// Set icon
		{
			in1: &Modifier{icon: &Icon{path: "./icon.png"}},
			in2: NewModifier().Icon(NewIcon("./icon.png")),
			out: `{"icon":{"path":"./icon.png"}}`,
		},
		// Set valid true
		{in1: &Modifier{valid: pb(true)}, in2: NewModifier().Valid(true), out: `{"valid":true}`},
		// Set valid false
		{in1: &Modifier{valid: pb(false)}, in2: NewModifier().Valid(false), out: `{"valid":false}`},
		// Set variables
		{
			in1: &Modifier{variables: map[string]string{"key": "value"}},
			in2: NewModifier().Variables(map[string]string{"key": "value"}),
			out: `{"variables":{"key":"value"}}`,
		},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalModifier", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in1, test.out)
			testMarshalJSON(t, i, test.in2, test.out)
		})
	}
}

func TestModifiers_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in1 *Modifiers
		in2 *Modifiers
		out string
	}

	tests := []marshalJSONTest{
		// Set shift
		{
			in1: &Modifiers{shift: &Modifier{subtitle: ps("shift")}},
			in2: NewModifiers().Shift(NewModifier().Subtitle("shift")),
			out: `{"shift":{"subtitle":"shift"}}`,
		},
		// Set fn
		{
			in1: &Modifiers{fn: &Modifier{subtitle: ps("fn")}},
			in2: NewModifiers().Fn(NewModifier().Subtitle("fn")),
			out: `{"fn":{"subtitle":"fn"}}`,
		},
		// Set ctrl
		{
			in1: &Modifiers{ctrl: &Modifier{subtitle: ps("ctrl")}},
			in2: NewModifiers().Ctrl(NewModifier().Subtitle("ctrl")),
			out: `{"ctrl":{"subtitle":"ctrl"}}`,
		},
		// Set alt
		{
			in1: &Modifiers{alt: &Modifier{subtitle: ps("alt")}},
			in2: NewModifiers().Alt(NewModifier().Subtitle("alt")),
			out: `{"alt":{"subtitle":"alt"}}`,
		},
		// Set cmd
		{
			in1: &Modifiers{cmd: &Modifier{subtitle: ps("cmd")}},
			in2: NewModifiers().Cmd(NewModifier().Subtitle("cmd")),
			out: `{"cmd":{"subtitle":"cmd"}}`,
		},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalModifiers", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in1, test.out)
			testMarshalJSON(t, i, test.in2, test.out)
		})
	}
}

func TestText_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in1 *Text
		in2 *Text
		out string
	}

	tests := []marshalJSONTest{
		// Set all
		{
			in1: &Text{copy: ps("copy"), largeType: ps("large")},
			in2: NewText().CopyText("copy").LargeText("large"),
			out: `{"copy":"copy","largetype":"large"}`,
		},
		// Set copy
		{in1: &Text{copy: ps("copy")}, in2: NewText().CopyText("copy"), out: `{"copy":"copy"}`},
		// Set largeType
		{in1: &Text{largeType: ps("large")}, in2: NewText().LargeText("large"), out: `{"largetype":"large"}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalText", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in1, test.out)
			testMarshalJSON(t, i, test.in2, test.out)
		})
	}
}

func TestItem_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in1 *Item
		in2 *Item
		out string
	}

	tests := []marshalJSONTest{
		// Minimal
		{in1: &Item{title: "title"}, in2: NewItem("title"), out: `{"title":"title"}`},
		// With uid
		{
			in1: &Item{title: "title", uid: ps("uid")},
			in2: NewItem("title").UID("uid"),
			out: `{"uid":"uid","title":"title"}`,
		},
		// With subtitle
		{
			in1: &Item{title: "title", subtitle: ps("sub")},
			in2: NewItem("title").Subtitle("sub"),
			out: `{"title":"title","subtitle":"sub"}`,
		},
		// With arg
		{
			in1: &Item{title: "title", arg: ps("arg")},
			in2: NewItem("title").Arg("arg"),
			out: `{"title":"title","arg":"arg"}`,
		},
		// With icon
		{
			in1: &Item{title: "title", icon: &Icon{path: "./icon.png"}},
			in2: NewItem("title").Icon(NewIcon("./icon.png")),
			out: `{"title":"title","icon":{"path":"./icon.png"}}`,
		},
		// With valid true
		{
			in1: &Item{title: "title", valid: pb(true)},
			in2: NewItem("title").Valid(true),
			out: `{"title":"title","valid":true}`,
		},
		// With valid false
		{
			in1: &Item{title: "title", valid: pb(false)},
			in2: NewItem("title").Valid(false),
			out: `{"title":"title","valid":false}`,
		},
		// With match
		{
			in1: &Item{title: "title", match: ps("match")},
			in2: NewItem("title").Match("match"),
			out: `{"title":"title","match":"match"}`,
		},
		// With autocomplete
		{
			in1: &Item{title: "title", autocomplete: ps("ac")},
			in2: NewItem("title").Autocomplete("ac"),
			out: `{"title":"title","autocomplete":"ac"}`,
		},
		// With typ default
		{
			in1: &Item{title: "title", typ: pItemType("default")},
			in2: NewItem("title").Type(ItemTypeDefault),
			out: `{"title":"title","type":"default"}`,
		},
		// With typ file
		{
			in1: &Item{title: "title", typ: pItemType("file")},
			in2: NewItem("title").Type(ItemTypeFile),
			out: `{"title":"title","type":"file"}`,
		},
		// With typ file:skipcheck
		{
			in1: &Item{title: "title", typ: pItemType("file:skipcheck")},
			in2: NewItem("title").Type(ItemTypeFileSkipCheck),
			out: `{"title":"title","type":"file:skipcheck"}`,
		},
		// With mods
		{
			in1: &Item{title: "title", mods: &Modifiers{shift: &Modifier{subtitle: ps("subtitle")}}},
			in2: NewItem("title").Mods(NewModifiers().Shift(NewModifier().Subtitle("subtitle"))),
			out: `{"title":"title","mods":{"shift":{"subtitle":"subtitle"}}}`,
		},
		// With text
		{
			in1: &Item{title: "title", text: &Text{copy: ps("copy")}},
			in2: NewItem("title").CopyText("copy"),
			out: `{"title":"title","text":{"copy":"copy"}}`,
		},
		// With quicklookURL
		{
			in1: &Item{title: "title", quicklookURL: ps("url")},
			in2: NewItem("title").QuicklookURL("url"),
			out: `{"title":"title","quicklookurl":"url"}`,
		},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalItem", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in1, test.out)
			testMarshalJSON(t, i, test.in2, test.out)
		})
	}
}
