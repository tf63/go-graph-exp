package entity

type Todo struct {
	Id   uint
	Text string
	Done bool
}

type NewTodo struct {
	Text string
}
