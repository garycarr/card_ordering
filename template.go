package main

import "net/http"

// TemplateParams is the struct that should be passed into all template calls that
// use the layout template.
type TemplateParams struct {
	// Layout is used by the layout template to render the stardard pieces of the
	// interface. NewTemplateParams is responsible for filling this out properly.
	Layout LayoutParams

	// Page is used to pass page specific data to templates for rendering.
	Page interface{}
}

// LayoutParams is the data needed by the layout template.
type LayoutParams struct {
	// Path is the path of the current request.
	Path string

	UpToCard int

	Attempts int

	AttemptsString string

	// Infos is a list of all info level statements to be rendered at the top of the content of the page.
	Results map[int]int

	// Infos is a list of all info level statements to be rendered at the top of the content of the page.
	Infos []string

	// Errors is a list of all error level statements to be rendered at the top of the content of the page.
	Errors []string
}

// NewTemplateParams returns a new TemplateParams object with the Layout field filled in
// properly for rendering content embedded inside a the layout template.
func (a *app) NewTemplateParams(w http.ResponseWriter, r *http.Request) *TemplateParams {
	params := &TemplateParams{
		Layout: LayoutParams{
			Path: r.RequestURI,
		},
	}
	return params
}
