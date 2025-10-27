package cmd

import (
	"fmt"
	"main/crud"
	"main/storage"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	crudInstance   = &crud.Amounts{}
	storageManager = storage.NewStorage[crud.Amounts]("data.json")
)

// loadData загружает данные из хранилища
func loadData() error {
	return storageManager.Load(crudInstance)
}

// saveData сохраняет данные в хранилище
func saveData() error {
	return storageManager.Save(*crudInstance)
}

func init() {
	// Загружаем данные при инициализации пакета
	if err := loadData(); err != nil {
		if os.IsNotExist(err) {
			// Файл не существует - создаем пустой список
			*crudInstance = crud.Amounts{}
			fmt.Println("No existing data found. Starting with empty list.")
		} else {
			fmt.Printf("Warning: Error loading data: %v\n", err)
			// Продолжаем с пустым списком
			*crudInstance = crud.Amounts{}
		}
	} else {
		fmt.Printf("Loaded %d existing records\n", len(*crudInstance))
	}
}

func Main() {
	var rootCMD = &cobra.Command{
		Use:   "expenses",
		Short: "Personal finance management tool",
		Long:  "A CLI application for managing personal expenses and financial records",
	}

	// ADD command
	var addCMD = &cobra.Command{
		Use:     "add <title> <description> <amount>",
		Short:   "Add new expense",
		Args:    cobra.ExactArgs(3),
		Example: "expenses add \"Groceries\" \"Weekly shopping\" 1500",
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			description := args[1]
			amount, err := strconv.Atoi(args[2])

			if err != nil || amount < 0 {
				fmt.Println("Error: Amount must be a positive number")
				return
			}

			// Показываем текущее количество записей перед добавлением
			currentCount := len(*crudInstance)
			fmt.Printf("Current records: %d\n", currentCount)

			if err := crudInstance.Add(title, description, amount); err != nil {
				fmt.Printf("Error adding record: %v\n", err)
				return
			}

			if err := saveData(); err != nil {
				fmt.Printf("Error saving data: %v\n", err)
				return
			}

			fmt.Printf("✅ Expense added successfully. Total records: %d\n", len(*crudInstance))
		},
	}

	// DELETE command
	var deleteCMD = &cobra.Command{
		Use:     "delete <index>",
		Short:   "Delete expense by index",
		Args:    cobra.ExactArgs(1),
		Example: "expenses delete 1",
		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error: Index must be a number")
				return
			}

			// Преобразуем в 0-based индекс
			actualIndex := index - 1

			if err := crudInstance.Delete(actualIndex); err != nil {
				fmt.Printf("Error deleting record: %v\n", err)
				return
			}

			if err := saveData(); err != nil {
				fmt.Printf("Error saving data: %v\n", err)
				return
			}

			fmt.Printf("✅ Expense #%d deleted successfully. Remaining records: %d\n", index, len(*crudInstance))
		},
	}

	// UPDATE command
	var updateCMD = &cobra.Command{
		Use:     "update <index> <title> <description> <amount>",
		Short:   "Update expense details",
		Args:    cobra.ExactArgs(4),
		Example: "expenses update 1 \"Supermarket\" \"Monthly groceries\" 2000",
		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Error: Index must be a number")
				return
			}

			title := args[1]
			description := args[2]
			amount, err := strconv.Atoi(args[3])
			if err != nil || amount < 0 {
				fmt.Println("Error: Amount must be a positive number")
				return
			}

			// Преобразуем в 0-based индекс
			actualIndex := index - 1

			if err := crudInstance.Update(actualIndex, title, description, amount); err != nil {
				fmt.Printf("Error updating record: %v\n", err)
				return
			}

			if err := saveData(); err != nil {
				fmt.Printf("Error saving data: %v\n", err)
				return
			}

			fmt.Printf("✅ Expense #%d updated successfully\n", index)
		},
	}

	// READ command
	var readCMD = &cobra.Command{
		Use:     "list",
		Short:   "Display all expenses",
		Aliases: []string{"read", "ls", "show"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := crudInstance.Read(); err != nil {
				fmt.Printf("Error reading records: %v\n", err)
			}
		},
	}

	// STATS command - новая команда для статистики
	var statsCMD = &cobra.Command{
		Use:     "stats",
		Short:   "Show statistics",
		Aliases: []string{"stat", "info"},
		Run: func(cmd *cobra.Command, args []string) {
			total := 0
			for _, amount := range *crudInstance {
				total += amount.Amount
			}

			fmt.Printf("📊 Statistics:\n")
			fmt.Printf("   Total records: %d\n", len(*crudInstance))
			fmt.Printf("   Total amount: %d\n", total)
			fmt.Printf("   Average amount: %.2f\n", float64(total)/float64(len(*crudInstance)))
		},
	}

	rootCMD.AddCommand(addCMD, deleteCMD, updateCMD, readCMD, statsCMD)

	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
