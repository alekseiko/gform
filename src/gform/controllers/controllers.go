package controllers

import (
    "net/http"
    "net/html"
    "net/html/atom"
	"github.com/corneldamian/httpway"
    "gform/models"
    "github.com/yhat/scrape"
)


func Index(w http.ResponseWriter, r *http.Request) {
    ctx := httpway.GetContext(r)
    ctx.Set("tmpl", "index.tmpl")
    ctx.Set("vars", "")
    ctx.Set("status", 200)
}

func StoreForm(s *models.Storage) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
        // get google form url from user
        gformUrl:=r.FormValue("url")

        // get google form
        resp, err:=http.Get(gformUrl)
        if err != nil {
            // TODO: error need to render template
            http.Error(w, "Couldn't download google form", http.BadRequest)
            return
        }

        // parse html
        gformPage, err:=html.Parse(resp.Body)
        if err != nil {
            // TODO: error need to render template
            http.Error(w, "Couldn't parse html of google form", http.BadRequest)
            return
        }

        // try to find <form> node
        gformMatcher:=func(n *html.Node) bool {
            return n != nil && n.DataAtom == atom.Form
        }
        gform, ok:=scrape.Find(gformPage, formMatcher)
        if !ok {
            // TODO: error need to render template
            http.Error(w, "Couldn't parse html of google form", http.BadRequest)
            return
        }

        form:=models.Form{}
        form.PageUrl = gformUrl
        form.Url = scrape.Attr(gform, "action")
        form.method = scrape.Attr(gform, "method")

        gformInputMatcher:= func (n *html.Node) bool {
            return n != nil && n.DataAtom == atom.Input
        }
        gformInputs,ok:=scrape.FindAll(gform)
        if !ok {
            // TODO: error need to render template
            http.(w, "Couldn't parse html of google form", http.BadRequest)
            return
        }
        for i, gformInput := range gformInputs {
            field:=models.Field{}
            field.FieldId=scrape.Attr(gformInput, "name")
            field.Name=scrape.Attr(gformInput, "label")
            field.MappedName=field.Name
            append(form.Fields, field)
        }
        s.SaveForm(&form)

        w.WriteHeader(200)
    }
}

func FormHandler(s *models.Storage) httpway.Handler {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}
