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
		in  *Icon
		out string
	}

	tests := []marshalJSONTest{
		// Minimal
		{in: &Icon{Path: "./icon.png"}, out: `{"path":"./icon.png"}`},
		// With type fileicon
		{in: &Icon{Path: "./icon.png", Type: pIconType("fileicon")}, out: `{"path":"./icon.png","type":"fileicon"}`},
		// With type filetype
		{in: &Icon{Path: "./icon.png", Type: pIconType("filetype")}, out: `{"path":"./icon.png","type":"filetype"}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalIcon", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in, test.out)
		})
	}
}

func TestModifier_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in  *Modifier
		out string
	}

	tests := []marshalJSONTest{
		// Set subtitle
		{in: &Modifier{Subtitle: ps("subtitle")}, out: `{"subtitle":"subtitle"}`},
		// Set arg
		{in: &Modifier{Arg: ps("arg")}, out: `{"arg":"arg"}`},
		// Set icon
		{in: &Modifier{Icon: &Icon{Path: "./icon.png"}}, out: `{"icon":{"path":"./icon.png"}}`},
		// Set valid true
		{in: &Modifier{Valid: pb(true)}, out: `{"valid":true}`},
		// Set valid false
		{in: &Modifier{Valid: pb(false)}, out: `{"valid":false}`},
		// Set variables
		{in: &Modifier{Variables: map[string]string{"key": "value"}}, out: `{"variables":{"key":"value"}}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalModifier", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in, test.out)
		})
	}
}

func TestModifiers_MarshalJSON(t *testing.T) {
	t.Parallel()

	// TODO: 3way test is better? or example test?
	type marshalJSONTest struct {
		in  *Modifiers
		out string
	}

	tests := []marshalJSONTest{
		// Set shift
		{in: &Modifiers{Shift: &Modifier{Subtitle: ps("shift")}}, out: `{"shift":{"subtitle":"shift"}}`},
		// Set fn
		{in: &Modifiers{Fn: &Modifier{Subtitle: ps("fn")}}, out: `{"fn":{"subtitle":"fn"}}`},
		// Set ctrl
		{in: &Modifiers{Ctrl: &Modifier{Subtitle: ps("ctrl")}}, out: `{"ctrl":{"subtitle":"ctrl"}}`},
		// Set alt
		{in: &Modifiers{Alt: &Modifier{Subtitle: ps("alt")}}, out: `{"alt":{"subtitle":"alt"}}`},
		// Set cmd
		{in: &Modifiers{Cmd: &Modifier{Subtitle: ps("cmd")}}, out: `{"cmd":{"subtitle":"cmd"}}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalModifiers", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in, test.out)
		})
	}
}

func TestText_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in  *Text
		out string
	}

	tests := []marshalJSONTest{
		// Set all
		{in: &Text{Copy: ps("copy"), LargeType: ps("large")}, out: `{"copy":"copy","largetype":"large"}`},
		// Set copy
		{in: &Text{Copy: ps("copy")}, out: `{"copy":"copy"}`},
		// Set largeType
		{in: &Text{LargeType: ps("large")}, out: `{"largetype":"large"}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalText", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in, test.out)
		})
	}
}

func TestItem_MarshalJSON(t *testing.T) {
	t.Parallel()

	type marshalJSONTest struct {
		in  *Item
		out string
	}
	tests := []marshalJSONTest{
		// Minimal
		{in: &Item{title: "title"}, out: `{"title":"title"}`},
		// With uid
		{in: &Item{title: "title", uid: ps("uid")}, out: `{"uid":"uid","title":"title"}`},
		// With subtitle
		{in: &Item{title: "title", subtitle: ps("sub")}, out: `{"title":"title","subtitle":"sub"}`},
		// With arg
		{in: &Item{title: "title", arg: ps("arg")}, out: `{"title":"title","arg":"arg"}`},
		// With icon
		{in: &Item{title: "title", icon: &Icon{Path: "./icon.png"}}, out: `{"title":"title","icon":{"path":"./icon.png"}}`},
		// With valid true
		{in: &Item{title: "title", valid: pb(true)}, out: `{"title":"title","valid":true}`},
		// With valid false
		{in: &Item{title: "title", valid: pb(false)}, out: `{"title":"title","valid":false}`},
		// With match
		{in: &Item{title: "title", match: ps("match")}, out: `{"title":"title","match":"match"}`},
		// With autocomplete
		{in: &Item{title: "title", match: ps("ac")}, out: `{"title":"title","match":"ac"}`},
		// With typ default
		{in: &Item{title: "title", typ: pItemType("default")}, out: `{"title":"title","type":"default"}`},
		// With typ file
		{in: &Item{title: "title", typ: pItemType("file")}, out: `{"title":"title","type":"file"}`},
		// With typ file:skipcheck
		{in: &Item{title: "title", typ: pItemType("file:skipcheck")}, out: `{"title":"title","type":"file:skipcheck"}`},
		// With mods
		{in: &Item{title: "title", mods: &Modifiers{Shift: &Modifier{Subtitle: ps("subtitle")}}},
			out: `{"title":"title","mods":{"shift":{"subtitle":"subtitle"}}}`},
		// With text
		{in: &Item{title: "title", text: &Text{Copy: ps("copy")}},
			out: `{"title":"title","text":{"copy":"copy"}}`},
		// With quicklookURL
		{in: &Item{title: "title", quicklookURL: ps("url")}, out: `{"title":"title","quicklookurl":"url"}`},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalItem", i), func(t *testing.T) {
			t.Parallel()
			testMarshalJSON(t, i, test.in, test.out)
		})
	}
}
