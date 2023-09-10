package repository

import (
	"database/sql"
	"time"

	"github.com/tf63/go-graph-exp/internal/entity"
)

type TodoRepository interface {
	CreateTodo(input entity.NewTodo) (TodoId int, err error)
	ReadTodos() (Todos []entity.Todo, err error)
}

type todoRepository struct {
	db *sql.DB // repositoryはdbへの接続を担う
}

// インターフェースを実装しているかチェック
var _ TodoRepository = (*todoRepository)(nil)

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) CreateTodo(input entity.NewTodo) (TodoId int, err error) {

	// 入力
	text := input.Text

	// トランザクションを開始
	tx, err := tr.db.Begin()
	if err != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	// エラーが生じた場合はロールバック
	defer tx.Rollback()

	// 実行するSQL (プレースホルダを使う)
	query := `
	INSERT INTO Todos (text, done, created_at, updated_at)
	VALUES (?, ?, ?, ?)
	`

	// 　SQL文の入力
	args := []interface{}{
		text,
		false,
		time.Now(),
		time.Now(),
	}

	// レコードの作成
	_, err = tr.db.Exec(query, args...)
	if err != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	// 戻り値を取得
	err = tr.db.QueryRow(`SELECT id FROM Todos ORDER BY id DESC LIMIT 1`).Scan(&TodoId)
	if err != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
	}

	return
}

func (tr *todoRepository) ReadTodos() (todos []entity.Todo, err error) {

	// レコードをlimit件取得
	record := []entity.Todo{}

	// 実行するSQL (プレースホルダを使う)
	query := `SELECT id, text, done, created_at, updated_at FROM Todos`

	// レコードを割り当てる
	err = tr.db.QueryRow(query).Scan(&record)

	if err != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	// 戻り値
	todos = record
	return
}
