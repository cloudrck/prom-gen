package datahandler

import (
	"html/template"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseYamltmpl(f string) (*template.Template, error) {
	templateFile := f // neet to switch for Windows
	templateData, err := os.ReadFile(templateFile)
	check(err)

	// Create a new template and parse the YAML template
	tmpl := template.Must(template.New("config").Parse(string(templateData)))
	return tmpl, nil
}

func GeneateYaml(tmpl *template.Template, yf string, s interface{}) {
	file, err := createF(yf)
	check(err)
	defer file.Close()

	err = tmpl.Execute(file, &s)
	check(err)
}
