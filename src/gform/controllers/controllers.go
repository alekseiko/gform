package controllers

import (
    "net/http"
	"github.com/corneldamian/httpway"
    "gopkg.in/mgo.v2"
)


func Index(w http.ResponseWriter, r *http.Request) {
    ctx := httpway.GetContext(r)
    ctx.Set("tmpl", "index.tmpl")
    ctx.Set("vars", "")
    ctx.Set("status", 200)
}

func StoreForm(session *mgo.Session) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}

func FormHandler(session *mgo.Session) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}
