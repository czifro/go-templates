// templatetutorial is based on https://github.com/Bamimore-Tomi/go-templates-guide/blob/example-01/main.go
package templatetutorial

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

type TemplateTutorial struct {
	rootDir  string
	examples map[string]ExampleTemplate
}

type ExampleTemplate struct {
	template *template.Template
	data     []interface{}
}

func New(root string) *TemplateTutorial {
	return &TemplateTutorial{
		rootDir:  root,
		examples: make(map[string]ExampleTemplate),
	}
}

func (t *TemplateTutorial) LoadTemplate(templateName string, data []interface{}) error {
	parse := func(templateFile string) (*template.Template, error) {
		if strings.HasSuffix(string(templateFile), "*") {
			return template.ParseGlob(string(templateFile))
		}
		return template.ParseFiles(string(templateFile))
	}
	tf := t.rootDir + "tmpls/" + templateName
	tmp, err := parse(tf)
	if err != nil {
		return err
	}
	t.examples[templateName] = ExampleTemplate{template: tmp, data: data}
	return nil
}

func (t *TemplateTutorial) LoadDefaultTemplates() error {
	data := []interface{}{nil}
	if err := t.LoadTemplate("ex01.tpl", data); err != nil {
		return err
	}
	data = []interface{}{"Frodo", "Sam"}
	if err := t.LoadTemplate("ex02/*", data); err != nil {
		return err
	}
	return nil
}

func (t *TemplateTutorial) ExecuteExample(ex string, output io.Writer) error {
	example, ok := t.examples[ex]
	if !ok {
		return fmt.Errorf("Could not find example %v", ex)
	}
	tmpls := example.template.Templates()
	sort.SliceStable(tmpls, func(i, j int) bool {
		return tmpls[i].Name() < tmpls[j].Name()
	})
	for i, tmpl := range tmpls {
		var data interface{}
		if i < len(example.data) {
			data = example.data[i]
		}
		log.Println(tmpl.ParseName)
		if err := tmpl.Execute(output, data); err != nil {
			return err
		}
	}
	return nil
}

func (t *TemplateTutorial) ExecuteAll() error {
	for k := range t.examples {
		if err := t.ExecuteExample(k, os.Stdout); err != nil {
			return err
		}
	}
	return nil
}
