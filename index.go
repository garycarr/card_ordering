package main

import (
	"net/http"

	"github.com/divan/num2words"
)

func (a *app) IndexHandlerGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templateParams := a.NewTemplateParams(w, r)
	templateParams.Layout.Attempts = a.attempts
	templateParams.Layout.AttemptsString = num2words.Convert(a.attempts)
	templateParams.Layout.UpToCard = a.upToCard
	templateParams.Layout.Results = a.results
	a.renderer.HTML(w, http.StatusOK, "index", templateParams)
}
