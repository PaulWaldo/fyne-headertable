package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/PaulWaldo/fyne-headertable/headertable"
)

var myTableOpts = headertable.TableOpts{
	RefWidth: "reference width",
	ColAttrs: []headertable.ColAttr{
		{
			Name:         "Name",
			Header:       "Name",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "Weight",
			Header:       "Weight",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
		{
			Name:         "Type",
			Header:       "Type",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 80,
		},
		{
			Name:         "Color",
			Header:       "Color",
			Alignment:    fyne.TextAlignCenter,
			TextStyle:    fyne.TextStyle{Bold: true},
			WidthPercent: 100,
		},
	},
}

type Animal struct {
	Name, Type, Color, Weight string
}

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}

func main() {
	var AnimalBindings []binding.DataMap

	// Create a binding for each animal data
	for i := 0; i < len(animals); i++ {
		AnimalBindings = append(AnimalBindings, binding.BindStruct(&animals[i]))
	}
	myTableOpts.Bindings = AnimalBindings
	a := app.New()
	w := a.NewWindow("Header Table Test")
	w.Resize(fyne.NewSize(624, 428))
	ht := headertable.NewSortingHeaderTable(&myTableOpts)
	w.SetContent(container.NewMax(ht))
	w.CenterOnScreen()
	w.ShowAndRun()
}
