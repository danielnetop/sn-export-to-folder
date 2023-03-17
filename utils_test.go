package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	standardNotesBackupAndImportFile = ExportData{
		Version: "004",
		Items: []Item{
			{
				ContentType: ContentTypeUserPreferences,
				Content: Content{
					References: []Reference{},
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 3, 4, 12, 30, 22, 426000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 3, 4, 13, 53, 40, 936000000, time.UTC),
				UUID:               "6afacecc-6dea-4625-aa66-111229fa0e67",
			},
			{
				ContentType: ContentTypeTheme,
				Content: Content{
					References: []Reference{},
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 3, 4, 12, 30, 22, 397000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				UUID:               "b75e80a4-a129-441c-b59b-a1b966f64eb9",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "This is the content of said tag",
					Title:            "New note on weird tag name",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 2, 25, 14, 42, 52, 736000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 3, 4, 11, 35, 4, 297000000, time.UTC),
				UUID:               "b81cc35d-bf5a-411e-a710-127f62585215",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{
						{
							UUID:        "7a7d1f9d-b392-4473-a79f-1d85cd1b28d5",
							ContentType: ContentTypeNote,
						},
						{
							UUID:        "bd61c08e-825c-4430-861e-026406695d6f",
							ContentType: ContentTypeNote,
						},
					},
					Title: "Tag 2",
				},
				CreatedAt: time.Date(2023, 3, 1, 14, 17, 34, 700000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 38, 44, 396000000, time.UTC),
				UUID:      "063d9f12-868f-4ae8-8103-32f23698edfe",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "Note without a title",
					Title:            "",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 2, 21, 1, 18, 12, 272000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 2, 21, 1, 18, 21, 512000000, time.UTC),
				UUID:               "23f8a1e5-3bcd-4431-9b78-ffb156b149ce",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "just another new note",
					Title:            "Sunday, Feb 19, 2023 at 9:25 PM",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 2, 19, 21, 25, 54, 753000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 2, 19, 21, 25, 58, 484000000, time.UTC),
				UUID:               "7a7d1f9d-b392-4473-a79f-1d85cd1b28d5",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{
						{
							UUID:          "50be64a8-cf6d-4bec-aa17-b2094a8c7b3b",
							ContentType:   ContentTypeTag,
							ReferenceType: ReferenceTypeTagToParentTag,
						},
						{
							UUID:        "d577853e-0c84-470e-9d49-4d518ad37488",
							ContentType: ContentTypeNote,
						},
					},
					Title: "SubTag of Tag 1",
				},
				CreatedAt: time.Date(2023, 3, 1, 14, 17, 36, 614000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 53, 42, 958000000, time.UTC),
				UUID:      "8842ba70-428b-42d8-b523-5e0b6983ceec",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{
						{
							UUID:        "b81cc35d-bf5a-411e-a710-127f62585215",
							ContentType: ContentTypeNote,
						},
					},
					Title: "<>:\"/\\|?* weird title",
				},
				CreatedAt: time.Date(2023, 2, 25, 14, 42, 50, 752000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 34, 48, 120000000, time.UTC),
				UUID:      "9306fc4e-91fe-467f-b4b2-bbacd7bdef31",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{
						{
							UUID:        "ed09f0be-cb1f-4ef2-8c5f-5f996f8907e0",
							ContentType: ContentTypeNote,
						},
						{
							UUID:        "23f8a1e5-3bcd-4431-9b78-ffb156b149ce",
							ContentType: ContentTypeNote,
						},
					},
					Title: "Emoji tag 😺",
				},
				CreatedAt: time.Date(2023, 3, 1, 14, 19, 11, 338000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 37, 19, 81000000, time.UTC),
				UUID:      "ccb70cb5-7acf-453e-9dd4-9a4e6412862d",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "new note for testing purposes",
					Title:            "Sunday, Feb 19, 2023 at 9:25 PM",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 02, 19, 21, 25, 46, 958000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 02, 19, 21, 25, 52, 906000000, time.UTC),
				UUID:               "e48301a9-156e-4094-8244-de1718921435",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "# asdasd\n",
					Title:            "Note with random text",
					NoteType:         "",
					EditorIdentifier: "",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 3, 3, 17, 39, 38, 206000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 3, 4, 11, 36, 55, 2000000, time.UTC),
				UUID:               "ed09f0be-cb1f-4ef2-8c5f-5f996f8907e0",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{},
					Title:      "Tag without notes",
				},
				CreatedAt: time.Date(2023, 3, 4, 11, 36, 5, 309000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 36, 10, 589000000, time.UTC),
				UUID:      "1ac36088-9a5f-4248-9fb3-28237fef3bd9",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{},
					Title:      "Deleted Tag",
				},
				CreatedAt: time.Date(2023, 3, 4, 11, 39, 18, 830000000, time.UTC),
				Deleted:   true,
				UpdatedAt: time.Date(2023, 3, 4, 11, 39, 20, 579000000, time.UTC),
				UUID:      "0ed5588e-53d9-4b61-b73c-238cfd98b0a5",
			},
			{
				ContentType: ContentTypeTag,
				Content: Content{
					References: []Reference{
						{
							UUID:        "bd61c08e-825c-4430-861e-026406695d6f",
							ContentType: ContentTypeNote,
						},
						{
							UUID:        "e48301a9-156e-4094-8244-de1718921435",
							ContentType: ContentTypeNote,
						},
					},
					Title: "Tag 1",
				},
				CreatedAt: time.Date(2023, 3, 4, 11, 36, 05, 309000000, time.UTC),
				UpdatedAt: time.Date(2023, 3, 4, 11, 53, 22, 241000000, time.UTC),
				UUID:      "50be64a8-cf6d-4bec-aa17-b2094a8c7b3b",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "",
					Title:            "Note that belongs to Tag 1 and Tag 2",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 3, 4, 11, 36, 57, 592000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 3, 4, 11, 37, 06, 899000000, time.UTC),
				UUID:               "bd61c08e-825c-4430-861e-026406695d6f",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "A note of the children of tag `1 SubTag of Tag 1`",
					Title:            "Note of subtag",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          false,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 03, 04, 11, 53, 42, 958000000, time.UTC),
				Deleted:            false,
				UpdatedAtTimestamp: 0,
				UpdatedAt:          time.Date(2023, 03, 04, 11, 54, 13, 357000000, time.UTC),
				UUID:               "d577853e-0c84-470e-9d49-4d518ad37488",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "A deleted note",
					Title:            "A deleted note",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
					Trashed:          true,
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 03, 04, 14, 54, 44, 935000000, time.UTC),
				Deleted:            false,
				UpdatedAt:          time.Date(2023, 03, 04, 14, 54, 49, 958000000, time.UTC),
				UUID:               "af9f8e2f-a3cb-48f0-befa-6388a4a34c7b",
			},
			{
				ContentType: ContentTypeNote,
				Content: Content{
					References:       []Reference{},
					Text:             "A note without tags",
					Title:            "A note without tags",
					NoteType:         "plain-text",
					EditorIdentifier: "com.standardnotes.plain-text",
				},
				CreatedAtTimestamp: 0,
				CreatedAt:          time.Date(2023, 03, 04, 14, 54, 44, 935000000, time.UTC),
				Deleted:            false,
				UpdatedAt:          time.Date(2023, 03, 04, 14, 54, 49, 958000000, time.UTC),
				UUID:               "2e9f7cfb-3001-4e56-bdc4-82656f6f747c",
			},
		},
	}
	standardNotesBackupAndImportTags = map[string]Tag{
		"063d9f12-868f-4ae8-8103-32f23698edfe": {Name: "Tag 2", Parent: ""},
		"8842ba70-428b-42d8-b523-5e0b6983ceec": {Name: "SubTag of Tag 1", Parent: "50be64a8-cf6d-4bec-aa17-b2094a8c7b3b"},
		"9306fc4e-91fe-467f-b4b2-bbacd7bdef31": {Name: "<>:\"/\\|?* weird title", Parent: ""},
		"ccb70cb5-7acf-453e-9dd4-9a4e6412862d": {Name: "Emoji tag 😺", Parent: ""},
		"1ac36088-9a5f-4248-9fb3-28237fef3bd9": {Name: "Tag without notes", Parent: ""},
		"50be64a8-cf6d-4bec-aa17-b2094a8c7b3b": {Name: "Tag 1", Parent: ""},
	}
	standardNotesBackupAndImportNotes = map[string]Note{
		"b81cc35d-bf5a-411e-a710-127f62585215": {
			UUID:      "b81cc35d-bf5a-411e-a710-127f62585215",
			Title:     "New note on weird tag name",
			Content:   "This is the content of said tag",
			UpdatedAt: time.Date(2023, 3, 4, 11, 35, 4, 297000000, time.UTC),
		},
		"23f8a1e5-3bcd-4431-9b78-ffb156b149ce": {
			UUID:      "23f8a1e5-3bcd-4431-9b78-ffb156b149ce",
			Title:     "23f8a1e5-3bcd-4431-9b78-ffb156b149ce",
			Content:   "Note without a title",
			UpdatedAt: time.Date(2023, 2, 21, 1, 18, 21, 512000000, time.UTC),
		},
		"7a7d1f9d-b392-4473-a79f-1d85cd1b28d5": {
			UUID:      "7a7d1f9d-b392-4473-a79f-1d85cd1b28d5",
			Title:     "Sunday, Feb 19, 2023 at 9-25 PM",
			Content:   "just another new note",
			UpdatedAt: time.Date(2023, 2, 19, 21, 25, 58, 484000000, time.UTC),
		},
		"e48301a9-156e-4094-8244-de1718921435": {
			UUID:      "e48301a9-156e-4094-8244-de1718921435",
			Title:     "Sunday, Feb 19, 2023 at 9-25 PM",
			Content:   "new note for testing purposes",
			UpdatedAt: time.Date(2023, 02, 19, 21, 25, 52, 906000000, time.UTC),
		},
		"ed09f0be-cb1f-4ef2-8c5f-5f996f8907e0": {
			UUID:      "ed09f0be-cb1f-4ef2-8c5f-5f996f8907e0",
			Title:     "Note with random text",
			Content:   "# asdasd\n",
			UpdatedAt: time.Date(2023, 3, 4, 11, 36, 55, 2000000, time.UTC),
		},
		"bd61c08e-825c-4430-861e-026406695d6f": {
			UUID:      "bd61c08e-825c-4430-861e-026406695d6f",
			Title:     "Note that belongs to Tag 1 and Tag 2",
			Content:   "",
			UpdatedAt: time.Date(2023, 3, 4, 11, 37, 06, 899000000, time.UTC),
		},
		"d577853e-0c84-470e-9d49-4d518ad37488": {
			UUID:      "d577853e-0c84-470e-9d49-4d518ad37488",
			Title:     "Note of subtag",
			Content:   "A note of the children of tag `1 SubTag of Tag 1`",
			UpdatedAt: time.Date(2023, 03, 04, 11, 54, 13, 357000000, time.UTC),
		},
		"2e9f7cfb-3001-4e56-bdc4-82656f6f747c": {
			UUID:      "2e9f7cfb-3001-4e56-bdc4-82656f6f747c",
			Title:     "A note without tags",
			Content:   "A note without tags",
			UpdatedAt: time.Date(2023, 03, 04, 14, 54, 49, 958000000, time.UTC),
		},
	}
	standardNotesBackupAndImportNoteTags = map[string][]string{
		"23f8a1e5-3bcd-4431-9b78-ffb156b149ce": {
			"ccb70cb5-7acf-453e-9dd4-9a4e6412862d",
		},
		"7a7d1f9d-b392-4473-a79f-1d85cd1b28d5": {
			"063d9f12-868f-4ae8-8103-32f23698edfe",
		},
		"bd61c08e-825c-4430-861e-026406695d6f": {
			"063d9f12-868f-4ae8-8103-32f23698edfe",
			"50be64a8-cf6d-4bec-aa17-b2094a8c7b3b",
		},
		"d577853e-0c84-470e-9d49-4d518ad37488": {
			"8842ba70-428b-42d8-b523-5e0b6983ceec",
		},
		"b81cc35d-bf5a-411e-a710-127f62585215": {
			"9306fc4e-91fe-467f-b4b2-bbacd7bdef31",
		},
		"ed09f0be-cb1f-4ef2-8c5f-5f996f8907e0": {
			"ccb70cb5-7acf-453e-9dd4-9a4e6412862d",
		},
		"e48301a9-156e-4094-8244-de1718921435": {
			"50be64a8-cf6d-4bec-aa17-b2094a8c7b3b",
		},
	}
)

func Test_getExportDataFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tests := []struct {
		name    string
		osArgs  []string
		want    string
		wantErr error
	}{
		{
			name:    "no file provided on input",
			osArgs:  []string{"cmd"},
			want:    "",
			wantErr: errNoFileProvided,
		},
		{
			name:    "file provided on input",
			osArgs:  []string{"cmd", "StandardNotesImport.txt"},
			want:    "StandardNotesImport.txt",
			wantErr: nil,
		},
		{
			name:    "no path provided on input",
			osArgs:  []string{"cmd", ""},
			want:    "",
			wantErr: errNoFileProvided,
		},
	}
	for _, tt := range tests {
		os.Args = tt.osArgs
		t.Run(tt.name, func(t *testing.T) {
			got, err := getExportDataFile()
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_readInputFile(t *testing.T) {
	tests := []struct {
		name                string
		exportDataInputFile string
		want                ExportData
		wantErr             error
	}{
		{
			name:                "file doesn't exist",
			exportDataInputFile: "testFiles/nonexistent.txt",
			want:                ExportData{},
			wantErr:             errReadingInputFile,
		},
		{
			name:                "file doesn't have right format",
			exportDataInputFile: "testFiles/wrongFormat.txt",
			want:                ExportData{},
			wantErr:             errUnmarshallingInputFile,
		},
		{
			name:                "empty file",
			exportDataInputFile: "testFiles/emptyFile.txt",
			want:                ExportData{},
			wantErr:             nil,
		},
		{
			name:                "real valid file",
			exportDataInputFile: "testFiles/Standard Notes Backup and Import File.txt",
			want:                standardNotesBackupAndImportFile,
			wantErr:             nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readInputFile(tt.exportDataInputFile)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseExportData(t *testing.T) {
	tests := []struct {
		name         string
		exportData   ExportData
		wantTags     map[string]Tag
		wantNotes    map[string]Note
		wantNoteTags map[string][]string
	}{
		{
			name:         "Empty export data",
			exportData:   ExportData{},
			wantTags:     make(map[string]Tag),
			wantNotes:    make(map[string]Note),
			wantNoteTags: make(map[string][]string),
		},
		{
			name:         "Standard Notes Backup And Import File export data",
			exportData:   standardNotesBackupAndImportFile,
			wantTags:     standardNotesBackupAndImportTags,
			wantNotes:    standardNotesBackupAndImportNotes,
			wantNoteTags: standardNotesBackupAndImportNoteTags,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, gotNote, gotNoteTags := parseExportData(tt.exportData)
			assert.Equal(t, tt.wantTags, gotTag)
			assert.Equal(t, tt.wantNotes, gotNote)
			assert.Equal(t, tt.wantNoteTags, gotNoteTags)
		})
	}
}

func Test_getFilePath(t *testing.T) {
	tests := []struct {
		name string
		tags map[string]Tag
		tag  Tag
		want string
	}{
		{
			name: "Return filepath of tag",
			tags: standardNotesBackupAndImportTags,
			tag:  standardNotesBackupAndImportTags["50be64a8-cf6d-4bec-aa17-b2094a8c7b3b"],
			want: "Tag 1",
		},
		{
			name: "Return filepath of tag with child",
			tags: standardNotesBackupAndImportTags,
			tag:  standardNotesBackupAndImportTags["8842ba70-428b-42d8-b523-5e0b6983ceec"],
			want: filepath.Join("Tag 1", "SubTag of Tag 1"),
		},
		{
			name: "Return filepath of tag with special characters",
			tags: standardNotesBackupAndImportTags,
			tag:  standardNotesBackupAndImportTags["9306fc4e-91fe-467f-b4b2-bbacd7bdef31"],
			want: "--------- weird title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getFilePath(tt.tags, tt.tag))
		})
	}
}

func tagPointer(tag Tag) *Tag {
	return &tag
}

func Test_getExportedFilePath(t *testing.T) {
	tests := []struct {
		name string
		tags map[string]Tag
		tag  *Tag
		want string
	}{
		{
			name: "Return filepath of tag",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["50be64a8-cf6d-4bec-aa17-b2094a8c7b3b"]),
			want: filepath.Join(exportDir, "Tag 1"),
		},
		{
			name: "Return filepath of tag with child",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["8842ba70-428b-42d8-b523-5e0b6983ceec"]),
			want: filepath.Join(exportDir, "Tag 1", "SubTag of Tag 1"),
		},
		{
			name: "Return filepath of tag with special characters",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["9306fc4e-91fe-467f-b4b2-bbacd7bdef31"]),
			want: filepath.Join(exportDir, "--------- weird title"),
		},
		{
			name: "Return only exportDir if tag is nil",
			tags: standardNotesBackupAndImportTags,
			tag:  nil,
			want: exportDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getExportedFilePath(tt.tags, tt.tag))
		})
	}
}

func Test_replaceFirstRune(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		replacement string
		want        string
	}{
		{
			name:        "replace first dot with dash",
			str:         "..",
			replacement: "-",
			want:        "-.",
		},
		{
			name:        "replace first dot with letter A",
			str:         "..",
			replacement: "A",
			want:        "A.",
		},
		{
			name:        "replace first space with dash",
			str:         " .",
			replacement: "-",
			want:        "-.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, replaceFirstRune(tt.str, tt.replacement))
		})
	}
}

func Test_sanitizeName(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "Replace special characters with a dash",
			filename: "<>:\"/\\|?*",
			want:     "---------",
		},
		{
			name:     "remove leading and trailing whitespaces",
			filename: "   :::   ",
			want:     "---",
		},
		{
			name:     "remove leading dot",
			filename: ".   Dot   ",
			want:     "-   Dot",
		},
		{
			name:     "remove leading dot if leading whitespaces",
			filename: "   .Dot   ",
			want:     "-Dot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sanitizeName(tt.filename))
		})
	}
}

func Test_updateTimes(t *testing.T) {
	var (
		note = Note{
			UUID:      "37997200-4d33-43e0-90c7-8b1552b0a32b",
			Title:     "Note created for testing purposes",
			Content:   "",
			UpdatedAt: time.Now(),
		}
		notePath = "testing_purposes.md"
	)

	require.NoError(t, createNote(note, notePath))

	tests := []struct {
		name      string
		path      string
		updatedAt time.Time
		err       error
	}{
		{
			name:      "No file exists",
			path:      "invalid_path_just_for_testing_purposes.go",
			updatedAt: time.Now(),
			err:       errUpdatingTimes,
		},
		{
			name:      "Update time of file",
			path:      notePath,
			updatedAt: time.Now().Add(10 * time.Second),
			err:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				fileInfoBeforeUpdate os.FileInfo
				err                  error
			)
			if tt.err == nil {
				fileInfoBeforeUpdate, err = os.Stat(tt.path)
				require.NoError(t, err)
			}

			assert.ErrorIs(t, tt.err, updateTimes(tt.path, tt.updatedAt))

			if tt.err == nil {
				fileInfo, err := os.Stat(tt.path)
				require.NoError(t, err)

				assert.True(t, tt.updatedAt.Equal(fileInfo.ModTime()))
				assert.NotEqual(t, fileInfoBeforeUpdate, fileInfo.ModTime())
			}
		})
	}

	require.NoError(t, os.RemoveAll(notePath))
}

func Test_nameOrUUID(t *testing.T) {
	tests := []struct {
		name  string
		title string
		uuid  string
		want  string
	}{
		{
			name:  "replace all whitespaces name with uuid",
			title: "          ",
			uuid:  "50877d67-92ea-4254-9e74-3856221d1a9e",
			want:  "50877d67-92ea-4254-9e74-3856221d1a9e",
		},
		{
			name:  "replace empty name with uuid",
			title: "",
			uuid:  "614328f3-167b-4ffb-a38a-8950303034a1",
			want:  "614328f3-167b-4ffb-a38a-8950303034a1",
		},
		{
			name:  "don't replace name with uuid",
			title: "valid not empty name",
			uuid:  "2f82f8a5-680d-48da-b181-6169f0d57053",
			want:  "valid not empty name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nameOrUUID(tt.title, tt.uuid))
		})
	}
}

func Test_pathExists(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "path exists",
			path: "utils_test.go",
			want: true,
		},
		{
			name: "path doesn't exist",
			path: "invalid_path_just_for_testing_purposes.go",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pathExists(tt.path))
		})
	}
}

func Test_checkIfPathExistsAndRename(t *testing.T) {
	tests := []struct {
		name      string
		filepath  string
		extraPath string
		want      string
	}{
		{
			name:      "path doesn't exist",
			filepath:  "invalid_path_just_for_testing_purposes.go",
			extraPath: "extra_path",
			want:      "invalid_path_just_for_testing_purposes.go",
		},
		{
			name:      "path exists",
			filepath:  "utils_test.go",
			extraPath: "extra_path",
			want:      "utils_test.go-extra_path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, checkIfPathExistsAndRename(tt.filepath, tt.extraPath))
		})
	}
}

func Test_createTagFolders(t *testing.T) {
	tests := []struct {
		name         string
		tags         map[string]Tag
		err          error
		validatePath bool
	}{
		{
			name:         "Successfully create folders for tags",
			tags:         standardNotesBackupAndImportTags,
			err:          nil,
			validatePath: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.False(t, pathExists(exportDir))

			assert.ErrorIs(t, tt.err, createTagFolders(tt.tags))
			if tt.validatePath {
				for _, tag := range tt.tags {
					tag := tag
					assert.True(t, pathExists(getExportedFilePath(tt.tags, &tag)))
				}
			}

			require.NoError(t, os.RemoveAll(exportDir))
		})
	}
}

func Test_createFolders(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Folder created",
			path:    "asd",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, createFolders(tt.path))
			assert.True(t, pathExists(tt.path))
			require.NoError(t, os.RemoveAll(tt.path))
		})
	}
}

func Test_createNotes(t *testing.T) {
	tests := []struct {
		name     string
		notes    map[string]Note
		tags     map[string]Tag
		noteTags map[string][]string
		err      error
	}{
		{
			name:     "Creating all notes",
			notes:    standardNotesBackupAndImportNotes,
			tags:     standardNotesBackupAndImportTags,
			noteTags: standardNotesBackupAndImportNoteTags,
			err:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.False(t, pathExists(exportDir))
			require.NoError(t, createTagFolders(tt.tags))
			assert.ErrorIs(t, tt.err, createNotes(tt.notes, tt.tags, tt.noteTags))

			require.NoError(t, os.RemoveAll(exportDir))
		})
	}
}

func Test_getFinalNotePath(t *testing.T) {
	testNote := Note{
		UUID:  "fcea99a9-ac85-4f3b-8b83-fa4e70440479",
		Title: "Title of Note",
	}
	tests := []struct {
		name string
		tags map[string]Tag
		tag  *Tag
		note Note
		want string
	}{
		{
			name: "Return full path of tagged note with extension",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["50be64a8-cf6d-4bec-aa17-b2094a8c7b3b"]),
			note: testNote,
			want: filepath.Join(exportDir, "Tag 1", testNote.Title+fileExtension),
		},
		{
			name: "Return full path of note with a tag and subtag with extension",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["8842ba70-428b-42d8-b523-5e0b6983ceec"]),
			note: testNote,
			want: filepath.Join(exportDir, "Tag 1", "SubTag of Tag 1", testNote.Title+fileExtension),
		},
		{
			name: "Return filepath of note with extension tagged with a tag that has special characters",
			tags: standardNotesBackupAndImportTags,
			tag:  tagPointer(standardNotesBackupAndImportTags["9306fc4e-91fe-467f-b4b2-bbacd7bdef31"]),
			note: testNote,
			want: filepath.Join(exportDir, "--------- weird title", testNote.Title+fileExtension),
		},
		{
			name: "Return filepath of note in exportDir",
			tags: standardNotesBackupAndImportTags,
			tag:  nil,
			note: testNote,
			want: filepath.Join(exportDir, testNote.Title+fileExtension),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getFinalNotePath(tt.tags, tt.tag, tt.note))
		})
	}
}
