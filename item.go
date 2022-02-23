// Copyright 2021 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package alfred

import (
	"encoding/json"
)

type IconType string

const (
	IconTypeFileIcon IconType = "fileicon"
	IconTypeFileType IconType = "filetype"
)

// Icon represents the icon for Item.
type Icon struct {
	path string
	typ  *IconType
}

// NewIcon returns an Icon with the given path.
func NewIcon(path string) *Icon {
	return &Icon{
		path: path,
	}
}

// NewIconWithType returns an Icon with the given path and name.
func NewIconWithType(path string, typ IconType) *Icon {
	return &Icon{
		path: path,
		typ:  &typ,
	}
}

// Path sets the path of the Item for Alfred results.
func (i *Icon) Path(path string) *Icon {
	i.path = path
	return i
}

// Type sets the type of the Item for Alfred results.
func (i *Icon) Type(typ IconType) *Icon {
	i.typ = &typ
	return i
}

// MarshalJSON implements the json.Marshaler interface.
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
	subtitle  *string
	arg       *string
	icon      *Icon
	valid     *bool
	variables map[string]string
}

// NewModifier returns a Modifier.
func NewModifier() *Modifier {
	return new(Modifier)
}

// Subtitle sets the subtitle of the Item for Alfred results.
func (m *Modifier) Subtitle(subtitle string) *Modifier {
	m.subtitle = &subtitle
	return m
}

// Arg sets the arg of the Item for Alfred results.
func (m *Modifier) Arg(arg string) *Modifier {
	m.arg = &arg
	return m
}

// Icon sets the icon of the Item for Alfred results.
func (m *Modifier) Icon(icon *Icon) *Modifier {
	m.icon = icon
	return m
}

// Valid sets valid of the Item for Alfred results.
func (m *Modifier) Valid(valid bool) *Modifier {
	m.valid = &valid
	return m
}

// Variables sets the variables of the Item for Alfred results.
func (m *Modifier) Variables(variables map[string]string) *Modifier {
	m.variables = variables
	return m
}

// MarshalJSON implements the json.Marshaler interface.
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

// Modifiers represents the modifier keys.
type Modifiers struct {
	shift *Modifier
	fn    *Modifier
	ctrl  *Modifier
	alt   *Modifier
	cmd   *Modifier
}

// NewModifiers returns a new initialized Modifiers
func NewModifiers() *Modifiers {
	return new(Modifiers)
}

// Shift sets Modifier when shift key is pressed.
func (m *Modifiers) Shift(modifier *Modifier) *Modifiers {
	m.shift = modifier
	return m
}

// Fn sets Modifier when fn key is pressed.
func (m *Modifiers) Fn(modifier *Modifier) *Modifiers {
	m.fn = modifier
	return m
}

// Ctrl sets Modifier when ctrl key is pressed.
func (m *Modifiers) Ctrl(modifier *Modifier) *Modifiers {
	m.ctrl = modifier
	return m
}

// Alt sets Modifier when alt key is pressed.
func (m *Modifiers) Alt(modifier *Modifier) *Modifiers {
	m.alt = modifier
	return m
}

// Cmd sets Modifier when cmd key is press.
func (m *Modifiers) Cmd(modifier *Modifier) *Modifiers {
	m.cmd = modifier
	return m
}

// MarshalJSON implements the json.Marshaler interface.
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

// Text represents the result text.
type Text struct {
	copy      *string
	largeType *string
}

// NewText returns a new initialized Text.
func NewText() *Text {
	return new(Text)
}

// CopyText sets the text.
func (t *Text) CopyText(text string) *Text {
	t.copy = &text
	return t
}

// LargeText sets the large text.
func (t *Text) LargeText(text string) *Text {
	t.largeType = &text
	return t
}

// MarshalJSON implements the json.Marshaler interface.
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

// Item represents the result item.
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

// NewItem returns a new initialized Item.
func NewItem(title string) *Item {
	return &Item{
		title: title,
	}
}

// NewInvalidItem returns a new initialized invalid Item.
func NewInvalidItem(title string) *Item {
	return &Item{
		title: title,
		valid: new(bool),
	}
}

// Uid sets the uid.
func (i *Item) Uid(uid string) *Item {
	i.uid = &uid
	return i
}

// Title sets the title.
func (i *Item) Title(title string) *Item {
	i.title = title
	return i
}

// Subtitle sets the subtitle.
func (i *Item) Subtitle(subtitle string) *Item {
	i.subtitle = &subtitle
	return i
}

// Arg sets the arg.
func (i *Item) Arg(arg string) *Item {
	i.arg = &arg
	return i
}

// Icon sets the icon.
func (i *Item) Icon(icon *Icon) *Item {
	i.icon = icon
	return i
}

// Valid sets the valid.
func (i *Item) Valid(valid bool) *Item {
	i.valid = &valid
	return i
}

// Match sets the match text.
func (i *Item) Match(match string) *Item {
	i.match = &match
	return i
}

// Autocomplete sets the autocomplete text.
func (i *Item) Autocomplete(autocomplete string) *Item {
	i.autocomplete = &autocomplete
	return i
}

// Type sets the ItemType.
func (i *Item) Type(typ ItemType) *Item {
	i.typ = &typ
	return i
}

// Mods sets the Modifiers.
func (i *Item) Mods(mods *Modifiers) *Item {
	i.mods = mods
	return i
}

// ModShift sets the Modifier when shift key is pressed.
func (i *Item) ModShift(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Shift(mod)
	return i
}

// ModFn sets the Modifier when fn key is pressed.
func (i *Item) ModFn(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Fn(mod)
	return i
}

// ModCtrl sets the Modifier when ctrl key is pressed.
func (i *Item) ModCtrl(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Ctrl(mod)
	return i
}

// ModAlt sets the Modifier when alt key is pressed.
func (i *Item) ModAlt(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Alt(mod)
	return i
}

// ModCmd sets the Modifier when cmd key is pressed.
func (i *Item) ModCmd(mod *Modifier) *Item {
	if i.mods == nil {
		i.mods = new(Modifiers)
	}
	i.mods.Cmd(mod)
	return i
}

// CopyText sets the copy text.
func (i *Item) CopyText(text string) *Item {
	if i.text == nil {
		i.text = new(Text)
	}
	i.text.CopyText(text)
	return i
}

// LargeText sets the large text.
func (i *Item) LargeText(text string) *Item {
	if i.text == nil {
		i.text = new(Text)
	}
	i.text.LargeText(text)
	return i
}

// Text sets the text.
func (i *Item) Text(text string) *Item {
	return i.CopyText(text).LargeText(text)
}

// QuicklookURL sets the quick look url.
func (i *Item) QuicklookURL(url string) *Item {
	i.quicklookURL = &url
	return i
}

// MarshalJSON implements the json.Marshaler interface.
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
