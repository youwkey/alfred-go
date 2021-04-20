package alfred

import (
	"encoding/json"
	"os"
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

func (sf *ScriptFilter) Output() error {
	bytes, err := json.Marshal(sf)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
