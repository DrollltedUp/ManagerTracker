package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type CRUD interface {
	Add(title, description string, amount int) error
	Delete(index int) error
	Update(index int, title string) error
	Read() error
}

type Amount struct {
	Title       string
	Description string
	Amount      int
	DateBuy     time.Time
}

type Amounts []Amount

//ADD

func (amounts *Amounts) Add(title, description string, amount int) error {

	a := Amount{
		Title:       title,
		Description: description,
		Amount:      amount,
		DateBuy:     time.Now(),
	}

	*amounts = append(*amounts, a)
	return nil
}

//DELETE

func (amounts *Amounts) Delete(index int) error {
	a := *amounts

	if err := a.isValidate(index); err != nil {
		return err
	}

	*amounts = append(a[:index], a[index+1:]...)
	return nil
}

func (amounts *Amounts) isValidate(index int) error {
	if index < 0 || index >= len(*amounts) {
		fmt.Println("Invalid Index")
	}

	return nil
}

//Update

func (amounts *Amounts) Update(index int, title string) error {
	a := *amounts

	if err := a.isValidate(index); err != nil {
		return err
	}
	a[index].Title = title

	return nil
}

//Read

func (amounts *Amounts) Read() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Title", "Description", "Amount", "Date"})
	for i, amount := range *amounts {
		createdAt := amount.DateBuy.Format("2006-01-02 15:04")
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			amount.Title,
			amount.Description,
			string(amount.Amount),
			string(createdAt),
		})
	}
	table.Render()

	return nil
}
