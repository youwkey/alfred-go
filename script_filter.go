package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	// TODO: unexported is better?
	// TODO: embedded is better?
	Items     []*Item           `json:"items"`
	// TODO: change type to interface{}?
	Variables map[string]string `json:"variables,omitempty"`
}

func (sf *ScriptFilter) Length() int {
	// TODO: Move to []Item method?
	return len(sf.Items)
}

func (sf *ScriptFilter) IsEmpty() bool {
	// TODO: Move to []Item method?
	return sf.Length() == 0
}

func (sf *ScriptFilter) Append(items ...*Item) {
	// TODO: Move to []Item method?
	sf.Items = append(sf.Items, items...)
}

func (sf *ScriptFilter) JsonMarshal() string {
	bytes, err := json.Marshal(sf)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (sf *ScriptFilter) Output() {
	// TODO: os.Stdout.Write is better?
	fmt.Println(sf.JsonMarshal())
}
