package alfred

import (
	"bytes"
	"fmt"
	"testing"
)

func TestScriptFilter_MarshalJSON(t *testing.T) {
	t.Parallel()

	type Test struct {
		in1 *ScriptFilter
		in2 *ScriptFilter
		out []byte
	}

	tests := []Test {
		// Minimal
		{ in1: &ScriptFilter{items: Items{}}, in2: NewScriptFilter(), out: []byte(`{"items":[]}`) },
		// With item
		{ in1: &ScriptFilter{items: Items{&Item{title: "title"}}},
			in2: NewScriptFilter().AppendItem(NewItem("title")),
			out: []byte(`{"items":[{"title":"title"}]}`) },
		// With variables
		{ in1: &ScriptFilter{items: Items{}, variables: map[string]string{"key": "value"}},
			in2: NewScriptFilter().Variables(map[string]string{"key": "value"}),
			out: []byte(`{"items":[],"variables":{"key":"value"}}`)},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(fmt.Sprintf("#%d:MarshalJSON", i), func(t *testing.T) {
			t.Parallel()
			in1, err := test.in1.MarshalJSON()
			if err != nil {
				t.Errorf("#%d: marshal error: %v", i, err)
			}
			if !bytes.Equal(in1, test.out) {
				t.Errorf("#%d: got: %s want: %s", i,  string(in1), string(test.out))
			}
			in2, err := test.in2.MarshalJSON()
			if err != nil {
				t.Errorf("#%d: marshal error: %v", i, err)
			}
			if !bytes.Equal(in2, test.out) {
				t.Errorf("#%d: got: %s want: %s", i, string(in2), string(test.out))
			}
		})
	}
}
