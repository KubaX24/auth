package main

import (
	"auth/utils"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var (
	key = []byte("lqjrKcITgK2U9eeqOWCpU3xUiwlJH2pjucU+g9hmN0I2Wxh/Sr0inWVqfAD+oAd2")
)

func Server() {
	router := mux.NewRouter()

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(static)

	router.HandleFunc("/", homePage)

	router.HandleFunc("/login", loginPage)
	router.HandleFunc("/login/{siteCode}", loginForSitePage)

	router.HandleFunc("/logout", logoutPage)
	router.HandleFunc("/account", accountPage)

	router.HandleFunc("/time/{timeToken}", timePage)

	fmt.Println(utils.Green + "Server started on port " + port + utils.Reset)

	_ = http.ListenAndServe(":"+port, router)
}

func loginForSitePage(w http.ResponseWriter, r *http.Request) {
	cookieGet, cookieError := r.Cookie("chyt.TOKEN")
	vars := mux.Vars(r)
	tmpl := template.Must(template.ParseFiles("templates/login_other.html"))

	var token string

	site, exitPoint := getWebsiteByCode(vars["siteCode"])

	if site == "" {
		fmt.Fprintf(w, "Error!")
		return
	}

	if r.Method == http.MethodPost {
		loginSuccess, loginErrorOrToken := checkLogin(
			UserDetails{
				r.FormValue("username"),
				r.FormValue("password"),
			},
			r.FormValue("cf-turnstile-response"))

		if !loginSuccess {
			fmt.Println("err", loginErrorOrToken)
			tmpl.Execute(w, struct{ WebsiteURL string }{site})
			return
		}

		cookie := http.Cookie{
			Name:   "chyt.TOKEN",
			Value:  loginErrorOrToken,
			Secure: true,
		}

		token = loginErrorOrToken

		http.SetCookie(w, &cookie)
	} else {
		if cookieError != nil {
			tmpl.Execute(w, struct{ WebsiteURL string }{site})
			return
		}

		if !isTokenValid(cookieGet.Value) {
			tmpl.Execute(w, struct{ WebsiteURL string }{site})
			return
		}

		token = cookieGet.Value
	}

	timeToken := utils.RandStringRunes(12) // TODO: CHANGE TO 512

	go ManageTimeTokens(timeToken, token)

	http.Redirect(w, r, "http://"+site+"/"+exitPoint+"?t="+timeToken, http.StatusSeeOther)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	cookieGet, cookieError := r.Cookie("chyt.TOKEN")

	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	var token string

	if r.Method == http.MethodPost {
		loginSuccess, loginErrorOrToken := checkLogin(
			UserDetails{
				r.FormValue("username"),
				r.FormValue("password"),
			},
			r.FormValue("cf-turnstile-response"))

		if !loginSuccess {
			fmt.Println("err", loginErrorOrToken)
			tmpl.Execute(w, nil)
			return
		}

		cookie := http.Cookie{
			Name:   "chyt.TOKEN",
			Value:  loginErrorOrToken,
			Secure: true,
		}

		token = loginErrorOrToken

		http.SetCookie(w, &cookie)
	} else {
		if cookieError != nil {
			tmpl.Execute(w, nil)
			return
		}

		if !isTokenValid(cookieGet.Value) {
			tmpl.Execute(w, nil)
			return
		}

		token = cookieGet.Value
	}

	fmt.Print("[auth] Token (interact): ")
	fmt.Println(token)

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}

func accountPage(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("chyt.TOKEN")

	if cookieError != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !isTokenValid(cookie.Value) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/account.html"))

	tmpl.Execute(w, struct{ Username interface{} }{parseToken(cookie.Value)["username"]})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "chyt.TOKEN",
		Value:  "",
		Secure: true,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func timePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	timeToken := vars["timeToken"]

	token, status := getTokenFromTime(timeToken)
	removeToken(timeToken)

	w.WriteHeader(status)
	_, _ = fmt.Fprintf(w, token)
}
