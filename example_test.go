package alfred_test

import (
	"github.com/youwkey/alfred-go"
)

func ExampleScriptFilter_Output() {
	item1 := alfred.NewItem("Title1").Subtitle("Sub1").Arg("Arg1")
	item2 := alfred.NewItem("Title2").Subtitle("Sub2").Arg("Arg2")

	sf := alfred.NewScriptFilter()
	sf.Items().Append(item1, item2)

	_ = sf.Output()
	// Output:
	// {"items":[{"title":"Title1","subtitle":"Sub1","arg":"Arg1"},{"title":"Title2","subtitle":"Sub2","arg":"Arg2"}]}
}

func ExampleScriptFilter_OutputIndent() {
	item1 := alfred.NewItem("Title1").Subtitle("Sub1").Arg("Arg1")
	item2 := alfred.NewItem("Title2").Subtitle("Sub2").Arg("Arg2")

	sf := alfred.NewScriptFilter()
	sf.Items().Append(item1, item2)
	sf.Variables().Put("key", "value")

	_ = sf.OutputIndent("", "  ")
	// Output:
	// {
	//   "items": [
	//     {
	//       "title": "Title1",
	//       "subtitle": "Sub1",
	//       "arg": "Arg1"
	//     },
	//     {
	//       "title": "Title2",
	//       "subtitle": "Sub2",
	//       "arg": "Arg2"
	//     }
	//   ],
	//   "variables": {
	//     "key": "value"
	//   }
	// }
}

func ExampleScriptFilter_OutputIndent_fullFields() {
	shift := alfred.NewModifier().
		Subtitle("Shift").
		Arg("ShiftArg").
		Icon(alfred.NewIcon("./shift.png").Type(alfred.IconTypeFileIcon)).
		Valid(true).
		Variables(map[string]string{"key": "shift"})
	fn := alfred.NewModifier().
		Subtitle("Fn").
		Arg("FnArg").
		Icon(alfred.NewIcon("./fn.png").Type(alfred.IconTypeFileIcon)).
		Valid(true).
		Variables(map[string]string{"key": "fn"})
	ctrl := alfred.NewModifier().
		Subtitle("Ctrl").
		Arg("CtrlArg").
		Icon(alfred.NewIcon("./ctrl.png").Type(alfred.IconTypeFileIcon)).
		Valid(true).
		Variables(map[string]string{"key": "ctrl"})
	alt := alfred.NewModifier().
		Subtitle("Alt").
		Arg("AltArg").
		Icon(alfred.NewIcon("./alt.png").Type(alfred.IconTypeFileIcon)).
		Valid(true).
		Variables(map[string]string{"key": "alt"})
	cmd := alfred.NewModifier().
		Subtitle("Cmd").
		Arg("CmdArg").
		Icon(alfred.NewIcon("./cmd.png").Type(alfred.IconTypeFileIcon)).
		Valid(true).
		Variables(map[string]string{"key": "cmd"})
	mods := alfred.NewModifiers().
		Shift(shift).
		Fn(fn).
		Ctrl(ctrl).
		Alt(alt).
		Cmd(cmd)

	item := alfred.NewItem("Title1").
		Uid("Uid").
		Subtitle("Sub").
		Arg("Arg").
		Icon(alfred.NewIcon("./icon.png").Type(alfred.IconTypeFileType)).
		Valid(true).
		Match("Match").
		Autocomplete("autocomplete").
		Type(alfred.ItemTypeDefault).
		Mods(mods).
		Text("Text").
		QuicklookURL("http://localhost")

	sf := alfred.NewScriptFilter()
	sf.Items().Append(item)
	sf.Variables().Put("key", "value")

	_ = sf.OutputIndent("", "  ")
	// Output:
	// {
	//   "items": [
	//     {
	//       "uid": "Uid",
	//       "title": "Title1",
	//       "subtitle": "Sub",
	//       "arg": "Arg",
	//       "icon": {
	//         "path": "./icon.png",
	//         "type": "filetype"
	//       },
	//       "valid": true,
	//       "match": "Match",
	//       "autocomplete": "autocomplete",
	//       "type": "default",
	//       "mods": {
	//         "shift": {
	//           "subtitle": "Shift",
	//           "arg": "ShiftArg",
	//           "icon": {
	//             "path": "./shift.png",
	//             "type": "fileicon"
	//           },
	//           "valid": true,
	//           "variables": {
	//             "key": "shift"
	//           }
	//         },
	//         "fn": {
	//           "subtitle": "Fn",
	//           "arg": "FnArg",
	//           "icon": {
	//             "path": "./fn.png",
	//             "type": "fileicon"
	//           },
	//           "valid": true,
	//           "variables": {
	//             "key": "fn"
	//           }
	//         },
	//         "ctrl": {
	//           "subtitle": "Ctrl",
	//           "arg": "CtrlArg",
	//           "icon": {
	//             "path": "./ctrl.png",
	//             "type": "fileicon"
	//           },
	//           "valid": true,
	//           "variables": {
	//             "key": "ctrl"
	//           }
	//         },
	//         "alt": {
	//           "subtitle": "Alt",
	//           "arg": "AltArg",
	//           "icon": {
	//             "path": "./alt.png",
	//             "type": "fileicon"
	//           },
	//           "valid": true,
	//           "variables": {
	//             "key": "alt"
	//           }
	//         },
	//         "cmd": {
	//           "subtitle": "Cmd",
	//           "arg": "CmdArg",
	//           "icon": {
	//             "path": "./cmd.png",
	//             "type": "fileicon"
	//           },
	//           "valid": true,
	//           "variables": {
	//             "key": "cmd"
	//           }
	//         }
	//       },
	//       "text": {
	//         "copy": "Text",
	//         "largetype": "Text"
	//       },
	//       "quicklookurl": "http://localhost"
	//     }
	//   ],
	//   "variables": {
	//     "key": "value"
	//   }
	// }
}
