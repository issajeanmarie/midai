package types

type AddTodoSuccessAction int

const (
	Exit AddTodoSuccessAction = iota
	Continue
)

type TodoStatus int

const (
	IsPending TodoStatus = iota + 1
	IsComplete
	IsInprogress
)

type UserAction int

const (
	AddTodo UserAction = iota + 1
	GetTodos
	DeleteTodo
	FindOneTodo
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Status      TodoStatus
}

type Todos struct {
	Todos []Todo
}

func (todos *Todos) AddTodo(todo Todo) {
	todos.Todos = append(todos.Todos, todo)
}

func (todos Todos) GetTodos() Todos {
	return todos
}
