package handlers

import (
	"html/template"
	"net/http"
	"path"
)

type IndexHandler struct{}

func (handler IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]string{
		"title":   "this is my first title with GO template HTML",
		"content": "lorem ipsum dolor siamet",
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, getMessage("", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, getMessage("", err), http.StatusInternalServerError)
		return
	}
}

func (handler IndexHandler) Product(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title": "GET Product",
	}

	tmpl, err := template.ParseFiles(path.Join("views", "get-product.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, getMessage("", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, getMessage("", err), http.StatusInternalServerError)
		return
	}
}

func (handler IndexHandler) ProcessProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, getMessage("", err), http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"title": "POST Product",
			"id":    r.Form.Get("id"),
			"name":  r.Form.Get("name"),
		}

		tmpl, err := template.ParseFiles(path.Join("views", "post-product.html"), path.Join("views", "layout.html"))

		if err != nil {
			http.Error(w, getMessage("", err), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, getMessage("", err), http.StatusInternalServerError)
			return
		}
	} else {
		handler.Product(w, r)
	}
}

func getMessage(msg string, err error) string {
	if err != nil || msg == "" {
		return err.Error()
	}
	return msg
}
