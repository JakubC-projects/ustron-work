package pages

import (
	"embed"
	"html/template"
)

//go:embed templates
var templatesFS embed.FS

var templates = template.Must(template.New("").ParseFS(templatesFS, "templates/*"))
