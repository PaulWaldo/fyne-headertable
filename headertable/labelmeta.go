package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ HeaderCellMeta = (*labelHeaderCellMeta)(nil)

type labelHeaderCellMeta struct {
	tableOpts *TableOpts
	DataTable *widget.Table
}

func NewLabelHeaderCellMeta(tableOpts *TableOpts) HeaderCellMeta {
	return &labelHeaderCellMeta{tableOpts: tableOpts}
}

func (m *labelHeaderCellMeta) NewHeader() *Header {
	h := &Header{Table: widget.Table{
		Length:     func() (int, int) { return 1, len(m.tableOpts.ColAttrs) },
		CreateCell: func() fyne.CanvasObject { return widget.NewLabel("the content") },
		UpdateCell: func(cellID widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*widget.Label)
			opts := m.tableOpts.ColAttrs[cellID.Col]
			l.SetText(opts.Header)
			l.TextStyle = opts.TextStyle
			l.Alignment = opts.Alignment
			l.Wrapping = opts.Wrapping
			l.Refresh()
		},
	}}
	h.ExtendBaseWidget(h)
	return h
}

func (m *labelHeaderCellMeta) UpdateDataTable() {
	panic("Not implemented")
}

func (m *labelHeaderCellMeta) SetDataTable(t *widget.Table) {
	m.DataTable = t
}

func (m *labelHeaderCellMeta) TableOpts() *TableOpts {
	return m.tableOpts
}
