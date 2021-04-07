package alfred

import "encoding/json"

type ScriptFilter struct {
	Items     []Item            `json:"items"`
	Variables map[string]string `json:"variables,omitempty"`
}

type Item struct {
	Uid          string         `json:"uid,omitempty"`
	Title        string         `json:"title"`
	Subtitle     string         `json:"subtitle,omitempty"`
	Arg          string         `json:"arg,omitempty"`
	Icon         *Icon          `json:"icon,omitempty"`
	Valid        bool           `json:"valid,omitempty"`
	Match        string         `json:"match,omitempty"`
	Autocomplete string         `json:"autocomplete,omitempty"`
	Type         ItemType       `json:"type,omitempty"`
	Mods         map[ModKey]Mod `json:"mods,omitempty"`
	Text         *Text          `json:"text,omitempty"`
	QuicklookURL string         `json:"quicklookurl,omitempty"`
}

type Icon struct {
	Type IconType `json:"type"`
	Path string   `json:"path"`
}

type IconType string

const (
	FileIcon IconType = "fileicon"
	FileType IconType = "filetype"
)

type ItemType string

const (
	Default       ItemType = "default"
	File          ItemType = "file"
	FileSkipCheck ItemType = "file:skipcheck"
)

type ModKey string

const (
	Shift ModKey = "shift"
	Fn    ModKey = "fn"
	Ctrl  ModKey = "ctrl"
	Alt   ModKey = "alt"
	Cmd   ModKey = "cmd"
)

type Mod struct {
	Valid     bool              `json:"valid,omitempty"`
	Arg       string            `json:"arg,omitempty"`
	Subtitle  string            `json:"subtitle,omitempty"`
	Icon      *Icon             `json:"icon,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type Text struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

func (sf *ScriptFilter) Length() int {
	return len(sf.Items)
}

func (sf *ScriptFilter) Append(items ...Item) {
	sf.Items = append(sf.Items, items...)
}

func (sf *ScriptFilter) ToJson(emptyItem Item) string {
	if sf.Length() == 0 {
		sf.Append(emptyItem)
	}
	bytes, err := json.Marshal(sf)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
