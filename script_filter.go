package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	items     Items
	variables map[string]string
}

func NewScriptFilter() *ScriptFilter {
	return &ScriptFilter{
		items:     make(Items, 0),
		variables: make(map[string]string),
	}
}

func (sf *ScriptFilter) MarshalJSON() ([]byte, error) {
	v := &struct {
		Items     Items             `json:"items"`
		Variables map[string]string `json:"variables,omitempty"`
	}{
		Items:     sf.items,
		Variables: sf.variables,
	}
	return json.Marshal(v)
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
