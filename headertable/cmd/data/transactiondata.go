package data

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"github.com/PaulWaldo/fyne-headertable/headertable"
)

var TableOpts = headertable.TableOpts{
	RefWidth: "reference width",
	ColAttrs: []headertable.ColAttr{
		{
			Name:   "Payee",
			Header: "Payee",
			HeaderStyle: headertable.CellStyle{
				Alignment: fyne.TextAlignCenter,
				TextStyle: fyne.TextStyle{Bold: true},
			},
			WidthPercent: 120,
		},
		{
			Name:   "Amount",
			Header: "Amount",
			HeaderStyle: headertable.CellStyle{
				Alignment: fyne.TextAlignCenter,
				TextStyle: fyne.TextStyle{Bold: true},
			},
			DataStyle: headertable.CellStyle{
				Alignment: fyne.TextAlignTrailing,
			},
			WidthPercent: 90,
			Converter:    func(i interface{}) string { return fmt.Sprintf("%0.2f", i) },
		},
		{
			Name:   "Memo",
			Header: "Memo",
			HeaderStyle: headertable.CellStyle{
				Alignment: fyne.TextAlignCenter,
				TextStyle: fyne.TextStyle{Bold: true},
			},
			DataStyle: headertable.CellStyle{
				Wrapping: fyne.TextTruncate,
			},
			WidthPercent: 120,
		},
		{
			Name:   "Date",
			Header: "Date",
			HeaderStyle: headertable.CellStyle{
				Alignment: fyne.TextAlignCenter,
				TextStyle: fyne.TextStyle{Bold: true},
			},
			DataStyle: headertable.CellStyle{
				TextStyle: fyne.TextStyle{Italic: true},
			},
			WidthPercent: 70,
			Converter: func(i interface{}) string {
				t := i.(time.Time)
				return t.Format("2006-01-02")
			},
		},
	},
}

type Transaction struct {
	Payee  string
	Amount float64
	Memo   string
	Date   time.Time
}

var Transactions = []Transaction{
	{Payee: "Grocery Store", Amount: -154.96, Memo: "Food for party that I hosted", Date: time.Now().Add(-5 * time.Hour * 24)},
	{Payee: "Drug store", Amount: -4.36, Memo: "Toothpaste", Date: time.Now().Add(-5 * time.Hour * 24)},
	{Payee: "My company", Amount: 2000, Memo: "Payday!", Date: time.Now().Add(-4 * time.Hour * 24)},
	{Payee: "Gasco", Amount: -63.78, Memo: "Gasoline", Date: time.Now().Add(-3 * time.Hour * 24)},
}
