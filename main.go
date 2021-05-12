package main

import (
	"fmt"

	"example.com/app/models"
)

func main() {
	fmt.Println(models.Db)

	// userを作成する
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)
	// {1 First Todo 2 2021-05-12 23:25:38.120923 +0900 +0900}

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// {1 First Todo 2 2021-05-12 23:25:38.120923 +0900 +0900}
	// {2 Second Todo 2 2021-05-12 23:36:02.418044 +0900 +0900}
	// user2, _ := models.GetUser(2)
	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// user_id = 2のユーザーに紐づくTodoのみを取得
	// {1 First Todo 2 2021-05-12 23:25:38.120923 +0900 +0900}
	// {2 Second Todo 2 2021-05-12 23:36:02.418044 +0900 +0900}

	t, _ := models.GetTodo(1)
	t.DeleteTodo()
	// {1|Update Todo|2|2021-05-12 23:25:38.120923+09:00}
}
