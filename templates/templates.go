package templates

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// LoginTemplate is template for login page for wiki
var LoginTemplate *template.Template

func init() {
	var allFiles []string
	templatesDir := "../templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates, err := template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	LoginTemplate = templates.Lookup("login.html")
}
