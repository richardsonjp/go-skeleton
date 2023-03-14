package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go-skeleton/sample/apis"
)

var (
	rootCMD = &cobra.Command{
		Short: "sample",
	}

	GetTodoListCMD = &cobra.Command{
		Use:     "get_todo_list",
		Aliases: []string{"gtl"},
		Short:   "call get todo list",
		Run:     func(*cobra.Command, []string) { apis.GetTodoList() },
	}

	CreateTodoCMD = &cobra.Command{
		Use:     "create_todo_item",
		Aliases: []string{"cti"},
		Short:   "call create todo item",
		Run:     func(*cobra.Command, []string) { apis.CreateTodoItem() },
	}
)

func main() {
	rootCMD.AddCommand(GetTodoListCMD)
	rootCMD.AddCommand(CreateTodoCMD)
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
