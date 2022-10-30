package headertable

import (
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/widget"
)

// type HeaderCellMeta interface {
// 	LengthFn() func() (int, int)
// 	CreateCellFn() func() fyne.CanvasObject
// 	UpdateCellFn() func(cellID widget.TableCellID, o fyne.CanvasObject)
// 	TableOpts() *TableOpts
// }

// type Header struct {
// 	widget.Table
// 	// tableOpts *TableOpts
// }

// func NewHeader(length func() (int, int), create func() fyne.CanvasObject, update func(widget.TableCellID, fyne.CanvasObject)) *Header{
// 	return &Header{}
// }

// func NewHeader(meta HeaderCellMeta) *Header {
// 	length := func() (int, int) {
// 		return 1, len(meta.TableOpts().ColAttrs)
// 	}

// 	createCell := func() fyne.CanvasObject {
// 		return widget.NewLabel("the content")
// 	}

// 	updateCell := func(cellID widget.TableCellID, o fyne.CanvasObject) {
// 		l := o.(*widget.Label)
// 		opts := tableOpts.ColAttrs[cellID.Col]
// 		l.SetText(opts.Header)
// 		l.TextStyle = opts.TextStyle
// 		l.Alignment = opts.Alignment
// 		l.Wrapping = opts.Wrapping
// 		l.Refresh()
// 	}
// 	h := &Header{
// 		// tableOpts: tableOpts,
// 		Table:     widget.Table{Length: length, CreateCell: createCell, UpdateCell: updateCell},
// 	}
// 	h.ExtendBaseWidget(h)

// 	return h
// }
