package main

import (
	"log"

	"github.com/czifro/go-templates/internal/templatetutorial"
)

func main() {
	t := templatetutorial.New("internal/templatetutorial/")
	if err := t.LoadDefaultTemplates(); err != nil {
		log.Fatalln(err)
	}
	if err := t.ExecuteAll(); err != nil {
		log.Fatalln(err)
	}
}
