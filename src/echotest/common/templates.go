package common

import "html/template"
import "io"
import "fmt"
import "errors"
import "bytes"


import "github.com/labstack/echo/v4"

type TemplateRenderer struct {
	Templates map[string]*template.Template
}

var templatesCache map[string]*template.Template

func init() {
 fmt.Println("Templates.init");
 templatesCache = make(map[string]*template.Template)
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
	  err := errors.New("Template not found -> " + name)
	  return err
	}

	// here add gloab user data?

	return tmpl.ExecuteTemplate(w, name, data)
}

//
// Расширяем шаблонизатор методом для получения данных пользователя
//

//
// Подготовка шаблонов
// 
func GetTemplates() (map[string]*template.Template) {
	templates := make(map[string]*template.Template)

	// index
    templates["home.html"] = template.Must(template.ParseFiles("view/static/home.html", "view/layouts/layout.html"))

	// login
    templates["login.html"] = template.Must(template.ParseFiles("view/static/login.html", "view/layouts/layout.html"))

	// profile
    templates["profile.html"] = template.Must(template.ParseFiles("view/profile/profile.html", "view/layouts/layout.html"))

	// errors
    templates["error401.html"] = template.Must(template.ParseFiles("view/errors/error401.html", "view/layouts/layout.html"))
    templates["error.html"] = template.Must(template.ParseFiles("view/errors/error.html", "view/layouts/layout.html"))
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	fmt.Printf("%v\n", templates);
	
    return templates;
}

type UserData struct {
	Name string `aaaaAAA`
}


  
func LazyRender(templateName string, templateFile string, layout string, data map[string]interface{}) (string, error) {
	
	// fmt.Print("Templates.LazyRender", templateName, templateFile, layout);

	tmpl, exists := templatesCache[templateName]

	if !exists {
		tmpl := template.Must(template.ParseFiles(templateFile, layout))
		// if err != nil {
		// 	return "", err
		// }

		templatesCache[templateName] = tmpl
	}

	// add global data
	data["User"] = UserData{Name: "asdasd"}
	fmt.Printf("data: %+v\n", data)

	// priint 
	buf := &bytes.Buffer{}

	// tmpl, err := template.ParseFiles("view/profile/profile.html", "view/layouts/layout.html")
	err := tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	// fmt.Printf("error2222: %+v\n", error)

	return buf.String(), nil
}