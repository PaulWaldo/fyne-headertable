package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Header struct {
	widget.Table
	tableOpts *TableOpts
}

func NewHeader(tableOpts *TableOpts) *Header {
	length := func() (int, int) {
		return 1, len(tableOpts.ColAttrs)
	}

	createCell := func() fyne.CanvasObject {
		return widget.NewLabel("the content")
	}

	updateCell := func(cellID widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*widget.Label)
		opts := tableOpts.ColAttrs[cellID.Col]
		l.SetText(opts.Header)
		l.TextStyle = opts.TextStyle
		l.Alignment = opts.Alignment
		l.Wrapping = opts.Wrapping
		l.Refresh()
	}
	h := &Header{
		tableOpts: tableOpts,
		Table:     widget.Table{Length: length, CreateCell: createCell, UpdateCell: updateCell},
	}
	h.ExtendBaseWidget(h)

	return h
}
