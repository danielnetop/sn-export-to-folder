package main

import (
	"errors"
	"fmt"
)

var errNoFileProvided = errors.New("path for `Standard Notes Backup File` is required")

var errExportFolderExists = fmt.Errorf("folder `%s` already exists", exportDir)

var errReadingInputFile = errors.New("error reading input file")

var errUnmarshallingInputFile = errors.New("error extracting data from input file")

var errEmptyFile = errors.New("input file is empty")

var errCreatingFolder = errors.New("error creating folder")

var errUpdatingTimes = errors.New("error updating times")

var errCreatingNote = errors.New("error creating note")

var errWritingNote = errors.New("error writing the content of note")

var errSavingNote = errors.New("error saving content note")
