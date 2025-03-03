package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/issajeanmarie/midai/types"
)

func HandleAddTodo(todos types.Todos) {

	for {
		fmt.Printf("|\n|\n🔺 ")
		UnderlinedTitle.Printf("ADD TODO No: %v\n", len(todos.Todos)+1)

		fmt.Print("|\tEnter title: ")
		titleScanner := bufio.NewScanner(os.Stdin)
		titleScanner.Scan()
		title := titleScanner.Text()

		// var description string
		fmt.Print("|\tEnter Description: ")
		descScanner := bufio.NewScanner(os.Stdin)
		descScanner.Scan()
		description := descScanner.Text()

		todo := types.Todo{
			ID:          len(todos.Todos) + 1,
			Title:       title,
			Description: description,
			Status:      types.IsPending,
		}

		todos.AddTodo(todo)

		var continueOption string
		fmt.Printf("|\n|\n✅ ")
		Title.Printf("TODO ADDED SUCCESFULLY!")
		fmt.Printf("\n|\tDo you want to continue? (0. No), (1. Yes): ")
		fmt.Scanln(&continueOption)

		res, err := strconv.Atoi(continueOption)
		if err != nil {
			if strings.ToLower(continueOption) == "yes" {
				continue
			}

			if strings.ToLower(continueOption) == "no" {
				exitTodoCreation(todos)
				break
			}

			fmt.Println("😭 We don't have that choice yet!")
			break
		}

		if types.AddTodoSuccessAction(res) == types.Exit {
			exitTodoCreation(todos)
			break
		}

	}

}

func exitTodoCreation(todos types.Todos) {
	fmt.Printf("\n|\n|👁️ ")
	UnderlinedTitle.Printf("HERE ARE YOUR TODOS\n")
	fmt.Println(todos)
}
