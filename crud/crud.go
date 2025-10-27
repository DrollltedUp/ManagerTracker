package crud

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type CRUD interface {
	Add(title, description string, amount int) error
	Delete(index int) error
	Update(index int, title, description string, amount int) error
	Read() error
	GetCount() int
}

type Amount struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	DateBuy     time.Time `json:"dateBuy"`
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

func (amounts *Amounts) Update(index int, title, description string, amount int) error {
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
			string(rune(amount.Amount)),
			string(createdAt),
		})
	}
	table.Render()

	return nil
}

// Get Count

func (amounts *Amounts) GetCount() int {

	return len(*amounts)
}
