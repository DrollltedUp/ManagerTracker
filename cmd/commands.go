package cmd

import (
	"fmt"
	"main/crud"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var crudInstance = &crud.Amounts{}

func Main() {
	var rootCMD = &cobra.Command{
		Use:   "expenses",
		Short: "Work with finance",
	}

	// add
	var addCMD = &cobra.Command{
		Use:   "add [Title] [description] amount",
		Short: "Add new Finance",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			description := args[1]
			amount, err := strconv.Atoi(args[2])

			if err != nil {
				fmt.Println("Некорректное значение amount")
				return
			}

			if err := crudInstance.Add(title, description, amount); err != nil {
				fmt.Println("Ошибка при добавлении:", err)
			} else {
				fmt.Println("Запись добавлена.")
			}
		},
	}

	var deleteCMD = &cobra.Command{
		Use:   "delete [index]",
		Short: "Delete for index",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Print("Incorrect Index", err)
				return
			}

			if err := crudInstance.Delete(index); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Finance is deleted")
			}
		},
	}

	var updateCMD = &cobra.Command{
		Use:   "update [index] [new title]",
		Short: "Update title for finance",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Print("Incorrect index ", err)
				return
			}
			newTitle := args[1]

			if err := crudInstance.Update(index-1, newTitle); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Title is Updates")
			}
		},
	}

	var readCMD = &cobra.Command{
		Use:   "read",
		Short: "Read all finance",
		Run: func(cmd *cobra.Command, args []string) {
			if err := crudInstance.Read(); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("All Readest")
			}
		},
	}

	rootCMD.AddCommand(addCMD, deleteCMD, updateCMD, readCMD)

	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
