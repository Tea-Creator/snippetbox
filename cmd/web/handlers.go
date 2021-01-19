package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	ts, err := template.ParseFiles([]string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}...)

	if err != nil {
		a.internalError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		a.internalError(w, err)
	}
}

func (a *app) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "Specific snippet with ID %d", id)
}

func (a *app) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := a.snippets.Insert("Carl", "Carl lived long time ago", "10")

	if err != nil {
		a.internalError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
