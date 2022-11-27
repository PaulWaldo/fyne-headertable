package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type CellStyle struct {
	Alignment fyne.TextAlign
	TextStyle fyne.TextStyle
	Wrapping  fyne.TextWrap
}

type ColAttr struct {
	Converter    BindingConverter
	DataStyle    CellStyle
	Header       string
	HeaderStyle  CellStyle
	Name         string
	WidthPercent int
}

type TableOpts struct {
	Bindings         []binding.Struct
	ColAttrs         []ColAttr
	OnDataCellSelect func(cellID widget.TableCellID)
	RefWidth         string
}

type Header struct {
	widget.Table
}
