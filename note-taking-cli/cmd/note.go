package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var appDataPath = os.Getenv("APPDATA")
var appFolder = filepath.Join(appDataPath, "Note")
var jsonfilePath = filepath.Join(appFolder, "note.json")

type Note struct {
	Content string `json:"content"`
}

type NoteWrapper struct {
	Notes []Note `json:"noteWrapper"`
}

var Notes = []string{}

func addNote(note string) (result string) {

	if checkFileExists(jsonfilePath) {

		if err := os.MkdirAll(appFolder, os.ModePerm); err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}

		myfile, err := os.Create(jsonfilePath)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}

		defer myfile.Close()
	}

	file, err := os.ReadFile(jsonfilePath)

	if err != nil {
		panic(err)
	}

	var savedNotes NoteWrapper
	if err := json.Unmarshal(file, &savedNotes); err != nil {
		fmt.Println("Error parsing json data", err)
	}

	newNote := Note{
		Content: note,
	}

	savedNotes.Notes = append(savedNotes.Notes, newNote)

	updatedData, err := json.MarshalIndent(savedNotes, "", " ")

	if err != nil {
		log.Fatalf("Error encoding JSON data: %v", err)
	}

	if err := os.WriteFile(jsonfilePath, updatedData, 0644); err != nil {
		return "failed to add the note"
	} else {
		return note
	}
}

func listNote() {

	if !checkFileExists(jsonfilePath) {
		file, err := os.ReadFile(jsonfilePath)

		if err != nil {
			panic(err)
		}

		var savedNotes NoteWrapper
		if err := json.Unmarshal(file, &savedNotes); err != nil {
			fmt.Println("Error parsing json data", err)
		}

		for i, notes := range savedNotes.Notes {
			fmt.Printf("%d: %s \n", i+1, notes.Content)
		}
	} else {
		fmt.Print("Notes does not exist")
	}

}

func deleteNote(id string) (result string) {

	if !checkFileExists(jsonfilePath) {

		file, err := os.ReadFile(jsonfilePath)

		if err != nil {
			panic(err)
		}

		var savedNotes NoteWrapper
		if err := json.Unmarshal(file, &savedNotes); err != nil {
			fmt.Println("Error parsing json data", err)
		}

		index, _ := strconv.Atoi(id)

		index -= 1

		if index >= 0 && index < len(savedNotes.Notes) {
			savedNotes.Notes = append(savedNotes.Notes[:index], savedNotes.Notes[index+1:]...)
		} else {
			return "Index out of range"
		}

		updatedData, err := json.MarshalIndent(savedNotes, "", " ")

		if err != nil {
			log.Fatalf("Error encoding JSON data: %v", err)
		}

		if err := os.WriteFile(jsonfilePath, updatedData, 0644); err != nil {
			return "failed to delete the note"
		} else {
			return "deleted the note"
		}
	} else {
		return "Notes does not exist"
	}
}

func checkFileExists(filepath string) bool {

	_, err := os.Stat(filepath)

	return errors.Is(err, os.ErrNotExist)
}
