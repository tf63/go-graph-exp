package repository

import (
	"database/sql"
	"log"

	"github.com/tf63/go-graph-exp/internal/entity"
)

type TodoRepository interface {
	CreateTodo(input entity.NewTodo) (todoId int, err error)
	ReadTodos() (todos []entity.Todo, err error)
	ReadTodo(todoId int) (todo entity.Todo, err error)
}

type todoRepository struct {
	Db *sql.DB // repositoryはdbへの接続を担う
}

// インターフェースを実装しているかチェック
var _ TodoRepository = (*todoRepository)(nil)

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

// Todoの作成
func (tr *todoRepository) CreateTodo(input entity.NewTodo) (todoId int, err error) {

	// 入力
	text := input.Text

	println("start transaction")
	// トランザクションを開始
	tx, err := tr.Db.Begin()
	if err != nil {
		log.Fatal(err)
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return 0, err
	}

	// エラーが生じた場合はロールバック
	defer tx.Rollback()

	// 実行するSQL (プレースホルダを使う)
	query := `
	INSERT INTO todos (text, done)
		VALUES ($1, $2)
		RETURNING id
	`

	// 　SQL文の入力
	args := []interface{}{
		text,
		false,
	}

	// レコードの作成
	err = tr.Db.QueryRow(query, args...).Scan(&todoId)
	if err != nil {
		log.Fatal(err)
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return 0, err
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return 0, err
	}

	return todoId, nil
}

// Todosの取得
func (tr *todoRepository) ReadTodos() (todos []entity.Todo, err error) {

	// // 取得するレコード
	// record := []entity.Todo{}

	// 実行するSQL (プレースホルダを使う)
	query := `SELECT id, text, done FROM todos LIMIT 100`

	// レコードを割り当てる
	rows, err := tr.Db.Query(query)
	if err != nil {
		log.Fatal(err)
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := entity.Todo{}
		err := rows.Scan(&todo.Id, &todo.Text, &todo.Done)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Todosの取得

// エラー時に初期値を返す
func (tr *todoRepository) ReadTodo(todoId int) (todo entity.Todo, err error) {

	// 実行するSQL
	query := `SELECT id, text, done FROM todos WHERE id = $1`

	// レコードを割り当てる
	err = tr.Db.QueryRow(query, todoId).Scan(&todo.Id, &todo.Text, &todo.Done)

	if err != nil {
		log.Fatal(err)

		// 初期値を返す
		return entity.Todo{}, err
	}

	return todo, nil
}
