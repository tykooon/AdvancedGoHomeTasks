package main

import (
	"context"
	"net/http"
	"strconv"

	"example.com/tykoon/pkg/usersystem"
)

type userAuth struct {
	email    string
	password string
}

func (app *app) homeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		app.notFound(writer)
	} else {
		http.Redirect(writer, request, "/login", http.StatusPermanentRedirect)
	}
}

func (app *app) loginHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		app.templates.login.ExecuteTemplate(writer, "base.html", true)
		return
	} else if request.Method == http.MethodPost {
		userData := userAuth{}
		userData.email = request.PostFormValue("email")
		userData.password = request.PostFormValue("password")
		user, err := app.users.GetByEmail(userData.email)
		if err == nil && user.PassCheck(userData.password) {
			SignIn(user, writer)
			for k := range request.Form {
				request.Form.Del(k)
			}
			http.Redirect(writer, request, "info", http.StatusPermanentRedirect)
			return
		}
		app.templates.login.ExecuteTemplate(writer, "base.html", false)
		return
	}
	app.httpVerbError(writer)
}

func Auth(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		uId, err := r.Cookie("UserId")
		if err == nil {
			userId, err := strconv.ParseInt(uId.Value, 10, 16)
			if err == nil {
				ctxWithUser := context.WithValue(r.Context(), "userId", userId)
				rWithUser := r.WithContext(ctxWithUser)
				next.ServeHTTP(w, rWithUser)
				return
			}
		}
		anonymusTemplate().ExecuteTemplate(w, "base.html", 0)
	}
	return http.HandlerFunc(f)
}

func (app *app) infoHandler(writer http.ResponseWriter, request *http.Request) {
	valueId := request.Context().Value("userId")
	userId, ok := valueId.(int64)
	if ok {
		user, err := app.users.GetById(userId)
		if err == nil {
			app.templates.info.ExecuteTemplate(writer, "base.html", user)
			return
		}
	}
	app.notFound(writer)
}

func SignIn(user *usersystem.User, writer http.ResponseWriter) {
	http.SetCookie(writer, &http.Cookie{Name: "UserId",
		Value: strconv.FormatInt(user.Id, 16)})
}
