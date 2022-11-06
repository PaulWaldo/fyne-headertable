package headertable

import (
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ HeaderCellMeta = (*sortLabelHeaderCellMeta)(nil)

type sortLabelHeaderCellMeta struct {
	tableOpts  *TableOpts
	DataTable  *widget.Table
	sortLabels []*SortingLabel
}

func NewSortLabelHeaderCellMeta(tableOpts *TableOpts) HeaderCellMeta {
	return &sortLabelHeaderCellMeta{tableOpts: tableOpts, sortLabels: make([]*SortingLabel, len(tableOpts.ColAttrs))}
}

func stringSort(tableOpts *TableOpts, col int) SortFn {
	return func(ascending bool) {
		log.Printf("Request to sort column %d ascending: %t\n", col, ascending)
		bindings := tableOpts.Bindings
		sort.Slice(bindings, func(i int, j int) bool {
			b1 := bindings[i]
			b2 := bindings[j]
			d1, err := b1.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				panic(err)
			}
			d2, err := b2.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				panic(err)
			}
			str1, err := d1.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			str2, err := d2.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			if ascending {
				return str1 < str2
			} else {
				return str1 > str2
			}
		})
	}
}

func (m *sortLabelHeaderCellMeta) NewHeader() *Header {
	h := &Header{Table: widget.Table{
		Length:     func() (int, int) { return 1, len(m.tableOpts.ColAttrs) },
		CreateCell: func() fyne.CanvasObject { return NewSortingLabel("the content", func() {}) },
		UpdateCell: func(cellID widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*SortingLabel)
			m.sortLabels[cellID.Col] = l
			col := cellID.Col
			opts := m.tableOpts.ColAttrs[col]
			l.Sorter = stringSort(m.TableOpts(), col)
			l.OnAfterSort = func() {
				m.UpdateDataTable()
				// Set all but this column to unsorted
				for i, sl := range m.sortLabels {
					if i != cellID.Col {
						sl.SetState(SortUnsorted)
					}
				}
			}
			l.Col = col
			l.Label.SetText(opts.Header)
			l.Label.TextStyle = opts.TextStyle
			l.Label.Alignment = opts.Alignment
			l.Label.Wrapping = opts.Wrapping
			l.Refresh()
		},
	}}
	h.ExtendBaseWidget(h)
	return h
}

func (m *sortLabelHeaderCellMeta) SetDataTable(t *widget.Table) {
	m.DataTable = t
}

func (m *sortLabelHeaderCellMeta) UpdateDataTable() {
	m.DataTable.Refresh()
}

func (m *sortLabelHeaderCellMeta) TableOpts() *TableOpts {
	return m.tableOpts
}
