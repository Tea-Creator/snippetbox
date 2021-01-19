package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tea-Creator/snippetbox/pkg/models"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	snippets, err := a.snippets.Latest()

	if err != nil {
		a.internalError(w, err)
		return
	}

	for _, s := range snippets {
		fmt.Fprintf(w, "%v\n", s)
	}
}

func (a *app) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		a.notFound(w)
		return
	}

	s, err := a.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
			return
		}

		a.internalError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", s)
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
