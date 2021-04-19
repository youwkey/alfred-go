package alfred

import "encoding/json"

type IconType string

const (
	IconTypeFileIcon IconType = "fileicon"
	IconTypeFileType IconType = "filetype"
)

type Icon struct {
	Path string    `json:"path"`
	Type *IconType `json:"type,omitempty"`
}

func NewIcon(path string, typ IconType) *Icon {
	return &Icon{
		Path: path,
		Type: &typ,
	}
}

type ItemType string

const (
	ItemTypeDefault       ItemType = "default"
	ItemTypeFile                   = "file"
	ItemTypeFileSkipCheck          = "file:skipcheck"
)

type Modifier struct {
	// TODO: require setter?
	Subtitle  *string           `json:"subtitle,omitempty"`
	Arg       *string           `json:"arg,omitempty"`
	Icon      *Icon             `json:"icon,omitempty"`
	Valid     *bool             `json:"valid,omitempty"`
	// TODO: change type to interface{}?
	Variables map[string]string `json:"variables,omitempty"`
}

type Modifiers struct {
	Shift *Modifier `json:"shift,omitempty"`
	Fn    *Modifier `json:"fn,omitempty"`
	Ctrl  *Modifier `json:"ctrl,omitempty"`
	Alt   *Modifier `json:"alt,omitempty"`
	Cmd   *Modifier `json:"cmd,omitempty"`
}

type Text struct {
	Copy      *string `json:"copy,omitempty"`
	LargeType *string `json:"largetype,omitempty"`
}

type Item struct {
	uid          *string
	title        string
	subtitle     *string
	arg          *string
	icon         *Icon
	valid        *bool
	match        *string
	autocomplete *string
	typ          *ItemType
	mods         *Modifiers
	text         *Text
	quicklookURL *string
}

func NewItem(title string) *Item {
	return &Item{
		title: title,
	}
}

func NewInvalidItem(title string) *Item {
	return &Item{
		title: title,
		valid: new(bool),
	}
}

func (i *Item) Uid(uid string) *Item {
	i.uid = &uid
	return i
}

func (i *Item) Title(title string) *Item {
	i.title = title
	return i
}

func (i *Item) Subtitle(subtitle string) *Item {
	i.subtitle = &subtitle
	return i
}

func (i *Item) Arg(arg string) *Item {
	i.arg = &arg
	return i
}

func (i *Item) Icon(icon *Icon) *Item {
	i.icon = icon
	return i
}

func (i *Item) Valid(valid bool) *Item {
	i.valid = &valid
	return i
}

func (i *Item) Match(match string) *Item {
	i.match = &match
	return i
}

func (i *Item) Autocomplete(autocomplete string) *Item {
	i.autocomplete = &autocomplete
	return i
}

func (i *Item) Type(typ ItemType) *Item {
	i.typ = &typ
	return i
}

func (i *Item) Mods(mods Modifiers) *Item {
	i.mods = &mods
	return i
}

func (i *Item) ModShift(mod Modifier) *Item {
	if i.mods == nil {
		i.mods = &Modifiers{Shift: &mod}
	} else {
		i.mods.Shift = &mod
	}
	return i
}

func (i *Item) ModFn(mod Modifier) *Item {
	if i.mods == nil {
		i.mods = &Modifiers{Fn: &mod}
	} else {
		i.mods.Shift = &mod
	}
	return i
}

func (i *Item) ModCtrl(mod Modifier) *Item {
	if i.mods == nil {
		i.mods = &Modifiers{Ctrl: &mod}
	} else {
		i.mods.Shift = &mod
	}
	return i
}

func (i *Item) ModAlt(mod Modifier) *Item {
	if i.mods == nil {
		i.mods = &Modifiers{Alt: &mod}
	} else {
		i.mods.Shift = &mod
	}
	return i
}

func (i *Item) ModCmd(mod Modifier) *Item {
	if i.mods == nil {
		i.mods = &Modifiers{Cmd: &mod}
	} else {
		i.mods.Shift = &mod
	}
	return i
}

func (i *Item) CopyText(text string) *Item {
	if i.text == nil {
		i.text = &Text{
			Copy: &text,
		}
	} else {
		i.text.Copy = &text
	}
	return i
}

func (i *Item) LargeText(text string) *Item {
	if i.text == nil {
		i.text = &Text{
			LargeType: &text,
		}
	} else {
		i.text.LargeType = &text
	}
	return i
}

func (i *Item) Text(text string) *Item {
	return i.CopyText(text).LargeText(text)
}

func (i *Item) QuicklookURL(url string) *Item {
	i.quicklookURL = &url
	return i
}

type ItemJSON struct {
	UID          *string    `json:"uid,omitempty"`
	Title        string     `json:"title"`
	Subtitle     *string    `json:"subtitle,omitempty"`
	Arg          *string    `json:"arg,omitempty"`
	Icon         *Icon      `json:"icon,omitempty"`
	Valid        *bool      `json:"valid,omitempty"`
	Match        *string    `json:"match,omitempty"`
	Autocomplete *string    `json:"autocomplete,omitempty"`
	Type         *ItemType  `json:"type,omitempty"`
	Mods         *Modifiers `json:"mods,omitempty"`
	Text         *Text      `json:"text,omitempty"`
	QuicklookURL *string    `json:"quicklookurl,omitempty"`
}

func (i *Item) MarshalJSON() ([]byte, error) {
	v := &ItemJSON{
		UID:          i.uid,
		Title:        i.title,
		Subtitle:     i.subtitle,
		Arg:          i.arg,
		Icon:         i.icon,
		Valid:        i.valid,
		Match:        i.match,
		Autocomplete: i.autocomplete,
		Type:         i.typ,
		Mods:         i.mods,
		Text:         i.text,
		QuicklookURL: i.quicklookURL,
	}
	return json.Marshal(v)
}
