package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserId    int
	CreatedAt time.Time
}

// TODO作成
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values ($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// TODO取得
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at
		from todos where id = $1`

	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserId,
		&todo.CreatedAt,
	)

	return todo, err
}

// TODO全件取得
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`

	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// ユーザーに紐づくTODOを全件取得
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from
		todos where user_id = $1`

	rows, err := Db.Query(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// TODO更新
func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set content = $1, user_id = $2 where id = $3`
	_, err = Db.Exec(cmd, t.Content, t.UserId, t.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// TODO削除
func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = $1`
	_, err = Db.Exec(cmd, t.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
