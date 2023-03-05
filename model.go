package main

import "time"

type ContentType string

const (
	ContentTypeNote            = "Note"
	ContentTypeTag             = "Tag"
	ContentTypeUserPreferences = "SN|UserPreferences"
	ContentTypeTheme           = "SN|Theme"
)

type ReferenceType string

const (
	ReferenceTypeTagToParentTag = "TagToParentTag"
)

type Reference struct {
	UUID          string        `json:"uuid"`
	ContentType   ContentType   `json:"content_type"`
	ReferenceType ReferenceType `json:"reference_type"`
}

type Content struct {
	References       []Reference `json:"references"`
	Text             string      `json:"text,omitempty"`
	Title            string      `json:"title,omitempty"`
	NoteType         string      `json:"noteType"`
	EditorIdentifier string      `json:"editorIdentifier"`
	Trashed          bool        `json:"trashed"`
}

type Item struct {
	ContentType        ContentType `json:"content_type"`
	Content            Content     `json:"content"`
	CreatedAtTimestamp int         `json:"created_at_timestamp"`
	CreatedAt          time.Time   `json:"created_at"`
	Deleted            bool        `json:"deleted"`
	UpdatedAtTimestamp int         `json:"updated_at_timestamp"`
	UpdatedAt          time.Time   `json:"updated_at"`
	UUID               string      `json:"uuid"`
}

type ExportData struct {
	Version string `json:"version"`
	Items   []Item `json:"items"`
}

type Tag struct {
	Name   string
	Parent string
}

type Note struct {
	UUID      string
	Title     string
	Content   string
	UpdatedAt time.Time
}
