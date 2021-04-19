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
	sf.Append(NewItem("TestTitle"))
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
	sf.Append(NewItem("TestTitle").
		Uid("uid01").
		Subtitle("TestSubtitle").
		Arg("OutputArg").
		Icon(NewIconWithType("~/icon.png", IconTypeFileIcon)).
		Valid(true).
		Match("TestMatchTitle").
		Autocomplete("ac").
		Type(ItemTypeDefault).
		ModShift(&Modifier{
			valid:     pb(true),
			arg:       ps("ModOutputArg"),
			subtitle:  ps("ModSubtitle"),
			icon:      NewIconWithType("public.png", IconTypeFileType),
			variables: map[string]string{"key": "value"},
		}).
		Text("Text").
		QuicklookURL("http://localhost"),
	)
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
					"path": "~/icon.png",
					"type": "fileicon"
				},
				"valid": true,
				"match": "TestMatchTitle",
				"autocomplete": "ac",
				"type": "default",
				"mods": {
					"shift": {
						"subtitle": "ModSubtitle",
						"arg": "ModOutputArg",
						"icon": {
							"path": "public.png",
							"type": "filetype"
						},
						"valid": true,
						"variables": {
							"key": "value"
						}
					}
				},
				"text": {
					"copy": "Text",
					"largetype": "Text"
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
