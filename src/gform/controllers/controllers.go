package controllers

import (
    "net/http"
	"github.com/corneldamian/httpway"
    "gform/models"
)


func Index(w http.ResponseWriter, r *http.Request) {
    ctx := httpway.GetContext(r)
    ctx.Set("tmpl", "index.tmpl")
    ctx.Set("vars", "")
    ctx.Set("status", 200)
}

func StoreForm(s models.Storage) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}

func FormHandler(s models.Storage) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}
