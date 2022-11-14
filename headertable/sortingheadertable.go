package headertable

import (
	"log"
	"math"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &HeaderTable{}

type SortingHeaderTable struct {
	widget.BaseWidget
	TableOpts  *TableOpts
	Header     *widget.Table
	Data       *widget.Table
	sortLabels []*sortingLabel
}

func NewSortingHeaderTable(tableOpts *TableOpts) *SortingHeaderTable {
	sortLabels := make([]*sortingLabel, len(tableOpts.ColAttrs))
	dataTable := widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) { return len(tableOpts.Bindings), len(tableOpts.ColAttrs) },

		// Default value
		func() fyne.CanvasObject { return widget.NewLabel("wide content") },

		// Cell values
		func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
			b := tableOpts.Bindings[cellID.Row]
			itemKey := tableOpts.ColAttrs[cellID.Col].Name
			d, err := b.GetItem(itemKey)
			if err != nil {
				log.Fatalf("Data table Update Cell callback, GetItem(%s): %s", itemKey, err)
			}
			str, err := d.(binding.String).Get()
			if err != nil {
				log.Fatalf("Data table Update Cell callback, Get: %s", err)
			}
			l := cnvObj.(*widget.Label)
			l.SetText(str)
		},
	)
	headerTable := widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) { return 1, len(tableOpts.ColAttrs) },
		// Default value
		func() fyne.CanvasObject { return NewSortingLabel("the content") },
		// Cell values
		func(cellID widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*sortingLabel)
			sortLabels[cellID.Col] = l
			col := cellID.Col
			opts := tableOpts.ColAttrs[col]
			l.Sorter = stringSort(tableOpts, col)
			l.OnAfterSort = func() {
				dataTable.Refresh()
				// Set all but this column to unsorted
				for i, sl := range sortLabels {
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
	)
	t := &SortingHeaderTable{
		sortLabels: sortLabels,
		TableOpts:  tableOpts,
		Header:     headerTable,
		Data:       dataTable,
	}
	t.ExtendBaseWidget(t)

	// Set Column widths
	refWidth := widget.NewLabel(t.TableOpts.RefWidth).MinSize().Width
	for i, colAttr := range t.TableOpts.ColAttrs {
		t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
		t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
	}

	return t
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

func (h *SortingHeaderTable) CreateRenderer() fyne.WidgetRenderer {
	return sortingHeaderTableRenderer{
		headerTable: h,
		container:   container.NewBorder(h.Header, nil, nil, nil, h.Data),
	}
}

//*******************************************************************************

var _ fyne.WidgetRenderer = sortingHeaderTableRenderer{}

type sortingHeaderTableRenderer struct {
	headerTable *SortingHeaderTable
	container   *fyne.Container
}

func (r sortingHeaderTableRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		float32(math.Max(float64(r.headerTable.Data.MinSize().Width), float64(r.headerTable.Header.MinSize().Width))),
		r.headerTable.Data.MinSize().Height+r.headerTable.Header.MinSize().Height)
}

func (r sortingHeaderTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r sortingHeaderTableRenderer) Destroy() {
}

func (r sortingHeaderTableRenderer) Refresh() {
	r.container.Refresh()
}

func (r sortingHeaderTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}
