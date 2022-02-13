// templatetutorial is based on https://github.com/Bamimore-Tomi/go-templates-guide/blob/example-01/main.go
package templatetutorial

import (
	"io"
	"log"
	"os"
	"text/template"
)

type TemplateFile string

const (
	ex01 TemplateFile = "template-impls/ex01.tpl"
	//ex02 TemplateFile = "template-impls/ex02.tpl"
	//ex03 TemplateFile = "template-impls/ex03.tpl"
)

func loadAndRender(templateFile TemplateFile, data interface{}, output io.Writer) error {
	tmp, err := template.ParseFiles(string(templateFile))
	if err != nil {
		return err
	}
	return tmp.Execute(output, data)
}

// Example01 runs the first example from the tutorial
func Example01() {
	err := loadAndRender(ex01, nil, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
