package alfred

import (
	"encoding/json"
	"os"
)

type Variables map[string]string

func (v *Variables) Put(key, value string) {
	(*v)[key] = value
}

type ScriptFilter struct {
	items     Items
	variables Variables
}

func NewScriptFilter() *ScriptFilter {
	return &ScriptFilter{
		items:     make(Items, 0),
		variables: make(map[string]string),
	}
}

func (sf *ScriptFilter) Items() *Items {
	return &sf.items
}

func (sf *ScriptFilter) Variables() *Variables {
	return &sf.variables
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
	return sf.output(bytes)
}

func (sf *ScriptFilter) OutputIndent(prefix, indent string) error {
	bytes, err := json.MarshalIndent(sf, prefix, indent)
	if err != nil {
		return err
	}
	return sf.output(bytes)
}

func (sf *ScriptFilter) output(bytes []byte) error {
	_, err := os.Stdout.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
