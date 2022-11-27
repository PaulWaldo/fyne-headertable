package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/PaulWaldo/fyne-headertable/headertable"
	"github.com/PaulWaldo/fyne-headertable/headertable/cmd/data"
)

func main() {
	// Create a binding for each transaction
	bindings := make([]binding.Struct, len(data.Transactions))
	for i := 0; i < len(data.Transactions); i++ {
		bindings[i] = binding.BindStruct(&data.Transactions[i])
	}
	data.TableOpts.Bindings = bindings
	a := app.New()
	w := a.NewWindow("Header Table Test")
	w.Resize(fyne.NewSize(600, 300))
	ht := headertable.NewHeaderTable(&data.TableOpts)
	w.SetContent(container.NewMax(ht))
	w.CenterOnScreen()
	w.ShowAndRun()
}
