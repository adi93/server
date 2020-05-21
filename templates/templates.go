package templates

import (
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"server/config"
)

// LoginTemplate is template for login page for wiki
var LoginTemplate *template.Template

func init() {
	var allFiles []string
	templatesDir := config.TemplateDir()
	if templatesDir == "" {
		// supply a default one
		templatesDir = "../templates/"
		return
	}

	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Fatalf("Could not open template dir: %s", err) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	templates, err := template.ParseFiles(allFiles...)
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err) // No point in running app if templates aren't read
	}
	LoginTemplate = templates.Lookup("login.html")
}
