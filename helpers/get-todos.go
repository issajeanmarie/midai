package helpers

import (
	"fmt"

	"github.com/issajeanmarie/midai/types"
)

func HandleGetTodos(todos types.Todos) {
	fmt.Println("TODOS: ", todos)
}
