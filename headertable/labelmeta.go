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

func (m *labelHeaderCellMeta) NewHeader(
	length func() (int, int),
	create func() fyne.CanvasObject,
	update func(widget.TableCellID, fyne.CanvasObject),
) *Header {
	h := &Header{Table: widget.Table{Length: length, CreateCell: create, UpdateCell: update}}
	h.ExtendBaseWidget(h)
	return h
}

func (m *labelHeaderCellMeta) UpdateDataTable() {
	panic("Not implimented")
}

func (m *labelHeaderCellMeta) SetDataTable(t *widget.Table) {
	m.DataTable = t
}

func (m labelHeaderCellMeta) LengthFn() func() (int, int) {
	return func() (int, int) { return 1, len(m.tableOpts.ColAttrs) }
}

func (m labelHeaderCellMeta) CreateCellFn() func() fyne.CanvasObject {
	return func() fyne.CanvasObject { return widget.NewLabel("the content") }
}

func (m labelHeaderCellMeta) UpdateCellFn() func(cellID widget.TableCellID, o fyne.CanvasObject) {
	return func(cellID widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*widget.Label)
		opts := m.tableOpts.ColAttrs[cellID.Col]
		l.SetText(opts.Header)
		l.TextStyle = opts.TextStyle
		l.Alignment = opts.Alignment
		l.Wrapping = opts.Wrapping
		l.Refresh()
	}
}

func (m *labelHeaderCellMeta) TableOpts() *TableOpts {
	return m.tableOpts
}
