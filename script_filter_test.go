package alfred

import (
	"bytes"
	"encoding/json"
	"testing"
)

func jsonCompact(s string) string {
	var expected bytes.Buffer
	_ = json.Compact(&expected, []byte(s))
	return expected.String()
}

func TestScriptFilter_ToJson_TitleOnly(t *testing.T) {
	sf := ScriptFilter{}
	sf.Append(Item{Title: "TestTitle"})
	got := sf.JsonMarshal()
	want := jsonCompact(`
	{
		"items": [
			{ "title": "TestTitle" }
		]
	}`)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestScriptFilter_ToJson_All(t *testing.T) {
	sf := ScriptFilter{}
	sf.Append(Item{
		Uid:          "uid01",
		Title:        "TestTitle",
		Subtitle:     "TestSubtitle",
		Arg:          "OutputArg",
		Icon:         &Icon{Type: FileIcon, Path: "~/icon.png"},
		Valid:        true,
		Match:        "TestMatchTitle",
		Autocomplete: "ac",
		Type:         Default,
		Mods: map[ModKey]Mod{Shift: {
			Valid:     true,
			Arg:       "ModOutputArg",
			Subtitle:  "ModSubtitle",
			Icon:      &Icon{Type: FileType, Path: "public.png"},
			Variables: map[string]string{"key": "value"},
		}},
		Text: &Text{
			Copy:      "CopyText",
			LargeType: "LargeText",
		},
		QuicklookURL: "http://localhost",
	})
	got := sf.JsonMarshal()
	want := jsonCompact(`
	{
		"items": [
			{
				"uid": "uid01",
				"title": "TestTitle",
				"subtitle": "TestSubtitle",
				"arg": "OutputArg",
				"icon": {
					"type": "fileicon",
					"path": "~/icon.png"
				},
				"valid": true,
				"match": "TestMatchTitle",
				"autocomplete": "ac",
				"type": "default",
				"mods": {
					"shift": {
						"valid": true,
						"arg": "ModOutputArg",
						"subtitle": "ModSubtitle",
						"icon": {
							"type": "filetype",
							"path": "public.png"
						},
						"variables": {
							"key": "value"
						}
					}
				},
				"text": {
					"copy": "CopyText",
					"largetype": "LargeText"
				},
				"quicklookurl": "http://localhost"
			}
		]
	}
	`)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
