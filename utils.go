package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const exportDir = "exported"
const fileExtension = ".md"

func getExportDataFile() (string, error) {
	// Validate arguments
	if len(os.Args) < 2 {
		return "", errNoFileProvided
	}

	flag.Parse()
	file := flag.Arg(0)

	if file == "" {
		return "", errNoFileProvided
	}

	return file, nil
}

func readInputFile(exportDataInputFile string) (ExportData, error) {
	exportDataFile, err := os.ReadFile(exportDataInputFile)
	if err != nil {
		return ExportData{}, errReadingInputFile
	}

	var exportData ExportData

	err = json.Unmarshal(exportDataFile, &exportData)
	if err != nil {
		return ExportData{}, errUnmarshallingInputFile
	}

	return exportData, nil
}

func parseExportData(exportData ExportData) (map[string]Tag, map[string]Note, map[string][]string) {
	var (
		tags     = make(map[string]Tag)
		notes    = make(map[string]Note)
		noteTags = make(map[string][]string)
	)

	for _, item := range exportData.Items {
		if item.ContentType == ContentTypeTag {
			parseTags(item, tags, noteTags)
		}

		if item.ContentType == ContentTypeNote {
			parseNotes(item, notes)
		}
	}

	return tags, notes, noteTags
}

func parseTags(item Item, tags map[string]Tag, noteTags map[string][]string) {
	if item.Deleted {
		fmt.Printf("Tag \"%s\":%s is deleted, will not be exported\n", item.Content.Title, item.UUID)

		return
	}

	var parent string

	for _, reference := range item.Content.References {
		if reference.ReferenceType == ReferenceTypeTagToParentTag {
			parent = reference.UUID
		}

		if reference.ContentType == ContentTypeNote {
			noteTags[reference.UUID] = append(noteTags[reference.UUID], item.UUID)
		}
	}

	tags[item.UUID] = Tag{
		Name:   nameOrUUID(item.Content.Title, item.UUID),
		Parent: parent,
	}
}

func parseNotes(item Item, notes map[string]Note) {
	if item.Content.Trashed {
		fmt.Printf("Note \"%s\":%s is deleted, will not be exported\n", item.Content.Title, item.UUID)

		return
	}

	notes[item.UUID] = Note{
		UUID:      item.UUID,
		Title:     sanitizeName(nameOrUUID(item.Content.Title, item.UUID)),
		Content:   item.Content.Text,
		UpdatedAt: item.UpdatedAt,
		CreatedAt: item.CreatedAt,
	}
}

func getFilePath(tags map[string]Tag, tag Tag) string {
	if tag.Parent == "" {
		return sanitizeName(tag.Name)
	}

	return filepath.Join(getFilePath(tags, tags[tag.Parent]), sanitizeName(tag.Name))
}

func getExportedFilePath(tags map[string]Tag, tag *Tag) string {
	if tag == nil {
		return exportDir
	}

	return filepath.Join(exportDir, getFilePath(tags, *tag))
}

func replaceFirstRune(str, replacement string) string {
	var sb strings.Builder

	sb.WriteString(string([]rune(str)[:0]))
	sb.WriteString(replacement)
	sb.WriteString(string([]rune(str)[1:]))

	return sb.String()
}

func sanitizeName(filename string) string {
	filename = strings.TrimSpace(filename)

	// If the following condition is true then both Name and item UUID are empty
	// I don't think that should happen on a valid export file
	// might need to update the nameOrUUID
	if len(filename) == 0 {
		return filename
	}

	filename = strings.ReplaceAll(filename, "<", "-")
	filename = strings.ReplaceAll(filename, ">", "-")
	filename = strings.ReplaceAll(filename, ":", "-")
	filename = strings.ReplaceAll(filename, "\"", "-")
	filename = strings.ReplaceAll(filename, "/", "-")
	filename = strings.ReplaceAll(filename, "\\", "-")
	filename = strings.ReplaceAll(filename, "|", "-")
	filename = strings.ReplaceAll(filename, "?", "-")
	filename = strings.ReplaceAll(filename, "*", "-")

	if string(filename[0]) == "." {
		filename = replaceFirstRune(filename, "-")
	}

	return filename
}

func updateTimes(path string, updatedAt time.Time, createdAt time.Time) error {
	err := os.Chtimes(path, updatedAt, createdAt)
	if err != nil {
		return errUpdatingTimes
	}

	return nil
}

func nameOrUUID(title, uuid string) string {
	if len(strings.TrimSpace(title)) == 0 {
		return uuid
	}

	return title
}

func pathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func checkIfPathExistsAndRename(filepath string, extraPath string) string {
	if pathExists(filepath) {
		return checkIfPathExistsAndRename(filepath+"-"+extraPath, strconv.FormatInt(time.Now().Unix(), 10))
	}

	return filepath
}

func createTagFolders(tags map[string]Tag) error {
	if len(tags) == 0 {
		err := createFolders(getExportedFilePath(nil, nil))
		if err != nil {
			return fmt.Errorf("%w \"%s\"", errCreatingFolder, exportDir)
		}

		return nil
	}

	for _, tag := range tags {
		tag := tag
		path := getExportedFilePath(tags, &tag)

		err := createFolders(path)
		if err != nil {
			return fmt.Errorf("%w \"%s\"", errCreatingFolder, path)
		}
	}

	return nil
}

func createFolders(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errCreatingFolder
	}

	return nil
}

func createNotes(notes map[string]Note, tags map[string]Tag, noteTags map[string][]string) error {
	for _, note := range notes {
		noteTags := noteTags[note.UUID]

		printDuplicatesHeader(noteTags, note)

		for _, noteTag := range noteTags {
			tag := tags[noteTag]
			notePath := getFinalNotePath(tags, &tag, note)

			printDuplicatePaths(noteTags, notePath)

			err := createNoteAndUpdateTimes(note, notePath)
			if err != nil {
				return err
			}
		}

		if len(noteTags) == 0 {
			err := createNoteAndUpdateTimes(note, getFinalNotePath(tags, nil, note))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getFinalNotePath(tags map[string]Tag, tag *Tag, note Note) string {
	notePath := filepath.Join(getExportedFilePath(tags, tag), note.Title)
	return checkIfPathExistsAndRename(notePath, note.UUID) + fileExtension
}

func printDuplicatesHeader(noteTags []string, note Note) {
	if len(noteTags) > 1 {
		fmt.Printf("Note \"%s\":%s is duplicated\n", note.Title, note.UUID)
		fmt.Printf("\tDuplicate paths:\n")
	}
}

func printDuplicatePaths(noteTags []string, notePath string) {
	if len(noteTags) > 1 {
		fmt.Printf("\t\t%s\n", notePath)
	}
}

func createNote(note Note, notePath string) error {
	f, err := os.Create(notePath)
	if err != nil {
		return errCreatingNote
	}

	_, err = f.Write([]byte(note.Content))
	if err != nil {
		return errWritingNote
	}

	err = f.Close()
	if err != nil {
		return errSavingNote
	}

	return nil
}

func createNoteAndUpdateTimes(note Note, notePath string) error {
	err := createNote(note, notePath)
	if err != nil {
		return fmt.Errorf("%w - \"%s\"", err, note.Title)
	}

	err = updateTimes(notePath, note.UpdatedAt, note.CreatedAt)
	if err != nil {
		return fmt.Errorf("%w - \"%s\"", err, notePath)
	}

	return nil
}
