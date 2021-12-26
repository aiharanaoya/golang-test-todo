package controllers

import (
	"log"
	"net/http"

	"github.com/aiharanaoya/golang-test-todo/app/models"
)

// 新規登録
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}

		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Panicln(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

// ログイン
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// ログイン処理
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	// 入力されたメールアドレスからDBに登録されているユーザーを取得
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Panicln(err)
		http.Redirect(w, r, "/login", 302)
	}

	// 入力されたパスワードが取得したユーザーのパスワードと一致するか
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Panicln(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// 登録されたセッションのUUIDをクッキーに保存
		// その後トップ画面リダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		// 一致しない場合ログイン画面にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
}

// ログアウト
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(writer, request, "/login", 302)
}
