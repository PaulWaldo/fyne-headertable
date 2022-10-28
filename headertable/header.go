package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Header struct {
	widget.Table
	// tableOpts *TableOpts
}

func NewHeader(tableOpts *TableOpts) *Header {
	h := &Header{
		// tableOpts: tableOpts,
	}
	h.ExtendBaseWidget(h)

	h.Length = func() (int, int) {
		return 1, len(tableOpts.ColAttrs)
	}

	h.CreateCell = func() fyne.CanvasObject {
		return widget.NewLabel("the content")
	}

	h.UpdateCell = func(cellID widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*widget.Label)
		opts := tableOpts.ColAttrs[cellID.Col]
		l.SetText(opts.Header)
		l.TextStyle = opts.TextStyle
		l.Alignment = opts.Alignment
		l.Wrapping = opts.Wrapping
	}
	return h
}
