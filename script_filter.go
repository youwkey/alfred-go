package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	Items     []*Item           `json:"items"`
	Variables map[string]string `json:"variables,omitempty"`
}

func (sf *ScriptFilter) Length() int {
	return len(sf.Items)
}

func (sf *ScriptFilter) IsEmpty() bool {
	return sf.Length() == 0
}

func (sf *ScriptFilter) Append(items ...*Item) {
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
	fmt.Println(sf.JsonMarshal())
}
