package main

import "net/http"

// AboutHandlerGET renders the main page.
func (app *app) AboutHandlerGET(w http.ResponseWriter, r *http.Request) {
	params := app.NewTemplateParams(w, r)
	app.renderer.HTML(w, http.StatusOK, "about", params)
}
