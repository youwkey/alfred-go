// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

import (
	"encoding/json"
	"fmt"
	"os"
)

// Variables describes a set of variables.
type Variables map[string]interface{}

// Put sets the variable.
func (v *Variables) Put(key string, value interface{}) {
	(*v)[key] = value
}

// ScriptFilter represents the output for Alfred Script Filter.
type ScriptFilter struct {
	items     Items
	variables Variables
}

// NewScriptFilter returns a new initialized alfred.ScriptFilter.
func NewScriptFilter() *ScriptFilter {
	return &ScriptFilter{
		items:     make(Items, 0),
		variables: make(map[string]interface{}),
	}
}

// Items returns Items bound to this ScriptFilter.
func (sf *ScriptFilter) Items() *Items {
	return &sf.items
}

// Variables return Variables bound to this ScriptFilter.
func (sf *ScriptFilter) Variables() *Variables {
	return &sf.variables
}

// MarshalJSON implements the json.Marshaler interface.
func (sf *ScriptFilter) MarshalJSON() ([]byte, error) {
	v := &struct {
		Items     Items                  `json:"items"`
		Variables map[string]interface{} `json:"variables,omitempty"`
	}{
		Items:     sf.items,
		Variables: sf.variables,
	}

	return json.Marshal(v)
}

// Output prints the Alfred Script Filter results to os.Stdout.
func (sf *ScriptFilter) Output() error {
	bytes, err := json.Marshal(sf)
	if err != nil {
		return err
	}

	return sf.output(bytes)
}

// OutputIndent is like Output but applies Indent to format the output.
func (sf *ScriptFilter) OutputIndent(prefix, indent string) error {
	bytes, err := json.MarshalIndent(sf, prefix, indent)
	if err != nil {
		return err
	}

	return sf.output(bytes)
}

func (sf *ScriptFilter) output(bytes []byte) error {
	if _, err := os.Stdout.Write(bytes); err != nil {
		return fmt.Errorf("os stdout write: %w", err)
	}

	return nil
}
