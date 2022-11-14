package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ColAttr struct {
	Alignment    fyne.TextAlign
	Header       string
	Name         string
	TextStyle    fyne.TextStyle
	WidthPercent int
	Wrapping     fyne.TextWrap
}

type TableOpts struct {
	Bindings         []binding.DataMap
	ColAttrs         []ColAttr
	OnDataCellSelect func(cellID widget.TableCellID)
	RefWidth         string
}

type Header struct {
	widget.Table
}
