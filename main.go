package main

import (
	"fmt"

	"github.com/aiharanaoya/golang-test-todo/app/controllers"
	"github.com/aiharanaoya/golang-test-todo/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Panicln("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println("ID:1のuser", u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println("ID:1のuser（更新後）", u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("user2 Todo2")

	// t, _ := models.GetTodo(5)
	// fmt.Println("ID:1のTODO", t)

	// u, _ := models.GetUser(2)
	// todos, _ := u.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println("TODO：", v)
	// }

	// t.Content = "更新後のTODO1"
	// t.UpdateTodo()

	// fmt.Println("ID:1のTODO", t)

	// t.DeleteTodo()

	// user, _ := models.GetUserByEmail("test@test.com")
	// fmt.Println("ユーザー", user)

	// session, err := user.CreateSession()

	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println("セッション", session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

	controllers.StartMainServer()
}
