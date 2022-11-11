package common

import "html/template"
import "io"
import "fmt"
import "errors"
import "bytes"

// echo
import "github.com/labstack/echo/v4"

type TemplateRenderer struct {
	Templates map[string]*template.Template
}

//
// Prepare teplates cashe
//
var templatesCache map[string]*template.Template

type templateItem struct {
	file string
	layout string
}

//
// Map of templates 
// name -> template data
//
var templatesMap map[string]templateItem


func init() {
	// 
 	fmt.Println("Templates.init");
 	templatesCache = make(map[string]*template.Template)


	templatesMap = make(map[string]templateItem)

	// index page
	templatesMap["index"] = templateItem{file: "view/static/index.html",}

	// profile inde page
	templatesMap["profile.index"] = templateItem{file: "view/profile/profile.html", layout: "view/layouts/profileLayout.html",}
}

//
// echo template renderer
//
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
    templates["error404.html"] = template.Must(template.ParseFiles("view/errors/error404.html", "view/layouts/layout.html"))
    templates["error.html"] = template.Must(template.ParseFiles("view/errors/error.html", "view/layouts/layout.html"))
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	
    // templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	fmt.Printf("%v\n", templates);
	
    return templates;
}


//
// Lazy rabder template profile template
// extend data with user data
//
func LazyProfileRender(c echo.Context, templateName string, data map[string]interface{}) (string, error) {	
	// user data
	userData, _ := GetAuthUser(c)
	data["User"] = userData

	return LazyRender(c, templateName, data)
}



//
// Lazy rabder template
//
func LazyRender(c echo.Context, templateName string, data map[string]interface{}) (string, error) {	
	return render(templateName, data)
}


//
// render template with cache
//
func render (templateName string, data map[string]interface{}) (string, error) {
	// check template cache exists
	tmpl, exists := templatesCache[templateName]

	if !exists {

		// template data
		tmplData, exists := templatesMap[templateName]
		if !exists {
			return "", errors.New("Template "+templateName+" noe exists")
		}

		if tmplData.layout != "" {
			tmpl = template.Must(template.ParseFiles(tmplData.file, tmplData.layout))
		} else {
			tmpl = template.Must(template.ParseFiles(tmplData.file))
		}
		// if err != nil {
		// 	return "", err
		// }

		templatesCache[templateName] = tmpl
	}

	// priint 
	buf := &bytes.Buffer{}

	// tmpl, err := template.ParseFiles("view/profile/profile.html", "view/layouts/layout.html")
	err := tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}