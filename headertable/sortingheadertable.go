package headertable

import (
	"fmt"
	"log"
	"math"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &HeaderTable{}

type SortingHeaderTable struct {
	widget.BaseWidget
	TableOpts  *TableOpts
	Header     *widget.Table
	Data       *widget.Table
	sortLabels []*SortingLabel
}

func NewSortingHeaderTable(tableOpts *TableOpts) *SortingHeaderTable {
	sortLabels := make([]*SortingLabel, len(tableOpts.ColAttrs))
	dataTable := widget.NewTable(dataTableLengthFunc(tableOpts), dataTableCreateFunc, dataTableUpdateFunc(tableOpts))
	headerTable := widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) { return 1, len(tableOpts.ColAttrs) },
		// Default value
		func() fyne.CanvasObject { return NewSortingLabel("the content") },
		// Cell values
		func(cellID widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*SortingLabel)
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
			l.Label.TextStyle = opts.HeaderStyle.TextStyle
			l.Label.Alignment = opts.HeaderStyle.Alignment
			l.Label.Wrapping = opts.HeaderStyle.Wrapping
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
			itemKey := tableOpts.ColAttrs[col].Name
			v1, err := b1.GetValue(itemKey)
			if err != nil {
				fyne.LogError("Error getting value for key:", err)
				return true
			}
			v2, err := b2.GetValue(itemKey)
			if err != nil {
				fyne.LogError("Error getting value for key:", err)
				return true
			}
			convert := tableOpts.ColAttrs[col].Converter
			if convert == nil {
				convert = func(i interface{}) string { return fmt.Sprintf("%s", i) }
			}
			str1 := convert(v1)
			str2 := convert(v2)

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

// ****************** Renderer *******************************

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
