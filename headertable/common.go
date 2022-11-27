package headertable

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Dimensions (rows, cols)
func dataTableLengthFunc(tableOpts *TableOpts) func() (int, int) {
	return func() (int, int) {
		return len(tableOpts.Bindings), len(tableOpts.ColAttrs)
	}
}

// Default value
var dataTableCreateFunc = func() fyne.CanvasObject { return widget.NewLabel("wide content") }

// Cell values
func dataTableUpdateFunc(tableOpts *TableOpts) func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
	return func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
		b := tableOpts.Bindings[cellID.Row]
		itemKey := tableOpts.ColAttrs[cellID.Col].Name
		v, err := b.GetValue(itemKey)
		if err != nil {
			fyne.LogError("Error getting value for key:", err)
			return
		}
		convert := tableOpts.ColAttrs[cellID.Col].Converter
		if convert == nil {
			convert = func(i interface{}) string { return fmt.Sprintf("%s", i) }
		}
		l := cnvObj.(*widget.Label)
		opts := tableOpts.ColAttrs[cellID.Col]
		l.TextStyle = opts.DataStyle.TextStyle
		l.Alignment = opts.DataStyle.Alignment
		l.Wrapping = opts.DataStyle.Wrapping
		l.SetText(convert(v))
	}
}
