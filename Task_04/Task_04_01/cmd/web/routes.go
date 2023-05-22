package main

import "net/http"

func (app *app) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/login", http.HandlerFunc(app.loginHandler))
	mux.Handle("/info", http.Handler(Auth(http.HandlerFunc(app.infoHandler))))
	mux.Handle("/", http.HandlerFunc(app.homeHandler))
	return mux
}
