package main

import (
	"log"
	"os"
	"text/template"
)

func saveToFile(bid float64) error {
	t, tErr := template.ParseFiles(TemplateName)
	if tErr != nil {
		log.Printf("error reading the template %v", tErr)
		return tErr
	}

	outputFile, createErr := os.Create(FileNameCotacao)
	if createErr != nil {
		return createErr
	}
	defer func(outputFile *os.File) {
		closeErr := outputFile.Close()
		if closeErr != nil {
			return
		}
	}(outputFile)

	tmplErr := t.Execute(outputFile, bid)
	if tmplErr != nil {
		log.Printf("error parsing the template %v", tmplErr)
		return tmplErr
	}

	log.Print("file created!")
	return nil
}
