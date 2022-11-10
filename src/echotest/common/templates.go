package common

import "html/template"
import "io"
import "errors"
import "github.com/labstack/echo/v4"

type TemplateRenderer struct {
	Templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
	  err := errors.New("Template not found -> " + name)
	  return err
	}
	return tmpl.ExecuteTemplate(w, name, data)
}

//
// Подготовка шаблонов
// 
func GetTemplates() (map[string]*template.Template) {
	templates := make(map[string]*template.Template)
    templates["home.html"] = template.Must(template.ParseFiles("view/static/home.html", "view/layouts/layout.html"))

	// login
    templates["login.html"] = template.Must(template.ParseFiles("view/static/login.html", "view/layouts/layout.html"))

	// errors
    templates["error401.html"] = template.Must(template.ParseFiles("view/errors/error401.html", "view/layouts/layout.html"))
    templates["error.html"] = template.Must(template.ParseFiles("view/errors/error.html", "view/layouts/layout.html"))
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	
    return templates;
}
  
