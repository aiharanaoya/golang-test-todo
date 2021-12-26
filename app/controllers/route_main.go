package controllers

import (
	"log"
	"net/http"

	"github.com/aiharanaoya/golang-test-todo/app/models"
)

// トップ
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "hello world!", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

	// t, _ := template.ParseFiles("app/views/templates/top.html")
	// t.Execute(w, "hello world!")
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

// TODO登録画面表示
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// TODO保存
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

// TODO編集画面表示
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}

		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

// TODO更新
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserId: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

// TODO削除
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}

		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
