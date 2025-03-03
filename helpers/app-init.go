package helpers

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/issajeanmarie/midai/types"
)

func InitiateApp(todos types.Todos) {

	for {
		fmt.Printf("🏁 ")
		UnderlinedTitle.Printf("WELCOME TO ISSA's TODO APP!\n|\n")
		fmt.Printf("|\tPRESS: \n")
		fmt.Printf("|\t1. To add todo\n|\t2. To list your todos \n|\t3. To delete a todo \n|\t4. To find one todo \n⚆\tEnter your choice: ")
		var userAction string
		fmt.Scanln(&userAction)

		res, err := strconv.Atoi(userAction)
		if err != nil {
			fmt.Print("Invalid choice")
			continue
		}

		if types.UserAction(res) == types.GetTodos {
			HandleGetTodos(todos)
			break
		}

		reflect.TypeOf("Hello").Kind()

		if types.UserAction(res) == types.AddTodo {
			HandleAddTodo(todos)
			break
		}

		fmt.Println("😭 We don't have that choice yet, let's try again!")
	}
}
