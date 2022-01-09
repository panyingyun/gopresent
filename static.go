package main

import "net/http"

func intiStaticFile() {
	//names := AssetNames()
	//for _, name := range names {
	//	fmt.Println("AssetDir name = ", name)
	//}
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/favicon.ico")
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(buffer)
	})
	http.HandleFunc("/static/styles.css", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/styles.css")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(buffer)
	})
	http.HandleFunc("/static/dir.css", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/dir.css")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(buffer)
	})
	http.HandleFunc("/static/jquery-ui.js", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/jquery-ui.js")
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(buffer)
	})
	http.HandleFunc("/static/notes.css", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/notes.css")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(buffer)
	})
	http.HandleFunc("/static/notes.js", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/notes.js")
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(buffer)
	})
	http.HandleFunc("/static/slides.js", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/slides.js")
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(buffer)
	})
	http.HandleFunc("/static/article.css", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/article.css")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(buffer)
	})
	http.HandleFunc("/static/dir.js", func(w http.ResponseWriter, r *http.Request) {
		buffer, _ := Asset("tpl/static/dir.js")
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(buffer)

	})
}
