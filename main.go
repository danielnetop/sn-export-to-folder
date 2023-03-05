package main

import (
	"fmt"
)

func main() {
	exportDataFile, err := getExportDataFile()
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if pathExists(exportDir) {
		fmt.Println(errExportFolderExists)

		return
	}

	exportData, err := readInputFile(exportDataFile)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if len(exportData.Items) == 0 {
		fmt.Println(errEmptyFile)

		return
	}

	tags, notes, noteTags := parseExportData(exportData)

	err = createTagFolders(tags)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	err = createNotes(notes, tags, noteTags)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
}
