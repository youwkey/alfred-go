package alfred

import (
	"encoding/json"
)

type IconType string

const (
	IconTypeFileIcon IconType = "fileicon"
	IconTypeFileType IconType = "filetype"
)

type Icon struct {
	path string
	typ  *IconType
}

func NewIcon(path string) *Icon {
	return &Icon{
		path: path,
	}
}

func NewIconWithType(path string, typ IconType) *Icon {
	return &Icon{
		path: path,
		typ:  &typ,
	}
}

func (i *Icon) Path(path string) *Icon {
	i.path = path
	return i
}

func (i *Icon) Type(typ IconType) *Icon {
	i.typ = &typ
	return i
}

func (i *Icon) MarshalJSON() ([]byte, error) {
	v := &struct {
		Path string    `json:"path"`
		Type *IconType `json:"type,omitempty"`
	}{
		Path: i.path,
		Type: i.typ,
	}
	return json.Marshal(v)
}

type ItemType string

const (
	ItemTypeDefault       ItemType = "default"
	ItemTypeFile                   = "file"
	ItemTypeFileSkipCheck          = "file:skipcheck"
)

type Modifier struct {
	subtitle *string
	arg      *string
	icon     *Icon
	valid    *bool
	// TODO: change type to interface{}?
	variables map[string]string
}

func NewModifier() *Modifier {
	return new(Modifier)
}

func (m *Modifier) Subtitle(subtitle string) *Modifier {
	m.subtitle = &subtitle
	return m
}

func (m *Modifier) Arg(arg string) *Modifier {
	m.arg = &arg
	return m
}

func (m *Modifier) Icon(icon *Icon) *Modifier {
	m.icon = icon
	return m
}

func (m *Modifier) Valid(valid bool) *Modifier {
	m.valid = &valid
	return m
}

func (m *Modifier) Variables(variables map[string]string) *Modifier {
	m.variables = variables
	return m
}

func (m *Modifier) MarshalJSON() ([]byte, error) {
	v := &struct {
		Subtitle  *string           `json:"subtitle,omitempty"`
		Arg       *string           `json:"arg,omitempty"`
		Icon      *Icon             `json:"icon,omitempty"`
		Valid     *bool             `json:"valid,omitempty"`
		Variables map[string]string `json:"variables,omitempty"`
	}{
		Subtitle:  m.subtitle,
		Arg:       m.arg,
		Icon:      m.icon,
		Valid:     m.valid,
		Variables: m.variables,
	}
	return json.Marshal(v)
}

type Modifiers struct {
	shift *Modifier
	fn    *Modifier
	ctrl  *Modifier
	alt   *Modifier
	cmd   *Modifier
}

func NewModifiers() *Modifiers {
	return new(Modifiers)
}

func (m *Modifiers) Shift(modifier *Modifier) *Modifiers {
	m.shift = modifier
	return m
}

func (m *Modifiers) Fn(modifier *Modifier) *Modifiers {
	m.fn = modifier
	return m
}

func (m *Modifiers) Ctrl(modifier *Modifier) *Modifiers {
	m.ctrl = modifier
	return m
}

func (m *Modifiers) Alt(modifier *Modifier) *Modifiers {
	m.alt = modifier
	return m
}

func (m *Modifiers) Cmd(modifier *Modifier) *Modifiers {
	m.cmd = modifier
	return m
}

func (m *Modifiers) MarshalJSON() ([]byte, error) {
	v := &struct {
		Shift *Modifier `json:"shift,omitempty"`
		Fn    *Modifier `json:"fn,omitempty"`
		Ctrl  *Modifier `json:"ctrl,omitempty"`
		Alt   *Modifier `json:"alt,omitempty"`
		Cmd   *Modifier `json:"cmd,omitempty"`
	}{
		Shift: m.shift,
		Fn:    m.fn,
		Ctrl:  m.ctrl,
		Alt:   m.alt,
		Cmd:   m.cmd,
	}
	return json.Marshal(v)
}

type Text struct {
	copy      *string
	largeType *string
}

func NewText() *Text {
	return new(Text)
}

func (t *Text) CopyText(text string) *Text {
	t.copy = &text
	return t
}

func (t *Text) LargeText(text string) *Text {
	t.largeType = &text
	return t
}

func (t *Text) MarshalJSON() ([]byte, error) {
	v := struct {
		Copy      *string `json:"copy,omitempty"`
		LargeType *string `json:"largetype,omitempty"`
	}{
		Copy:      t.copy,
		LargeType: t.largeType,
	}
	return json.Marshal(v)
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

func (i *Item) Mods(mods *Modifiers) *Item {
	i.mods = mods
	return i
}

func (i *Item) ModShift(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Shift(mod)
	return i
}

func (i *Item) ModFn(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Fn(mod)
	return i
}

func (i *Item) ModCtrl(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Ctrl(mod)
	return i
}

func (i *Item) ModAlt(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Alt(mod)
	return i
}

func (i *Item) ModCmd(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Cmd(mod)
	return i
}

func (i *Item) CopyText(text string) *Item {
	if i.text == nil {
		i.text = new(Text)
	}
	i.text.CopyText(text)
	return i
}

func (i *Item) LargeText(text string) *Item {
	if i.text == nil {
		i.text = new(Text)
	}
	i.text.LargeText(text)
	return i
}

func (i *Item) Text(text string) *Item {
	return i.CopyText(text).LargeText(text)
}

func (i *Item) QuicklookURL(url string) *Item {
	i.quicklookURL = &url
	return i
}

func (i *Item) MarshalJSON() ([]byte, error) {
	v := &struct {
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
	}{
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
