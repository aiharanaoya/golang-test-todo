package main

import (
	"fmt"

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

	u, _ := models.GetUser(1)
	fmt.Println("ID:1のuser", u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println("ID:1のuser（更新後）", u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
