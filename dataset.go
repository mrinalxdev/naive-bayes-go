package main

import (
	"os"
	"path/filepath"

)


type Document struct {
	Class string
	Text  string
}


func LoadDocumentsFromDirectory(root string) ([]Document, error) {
	var docs []Document

	subdirs, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, subdir := range subdirs {
		if !subdir.IsDir() {
			continue
		}
		class := subdir.Name()
		classPath := filepath.Join(root, class)

		files, err := os.ReadDir(classPath)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			filePath := filepath.Join(classPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}
			doc := Document{
				Class: class,
				Text:  string(content),
			}
			docs = append(docs, doc)
		}
	}
	return docs, nil
}