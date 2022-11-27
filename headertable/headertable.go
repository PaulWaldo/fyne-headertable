package headertable

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = &HeaderTable{}

type HeaderTable struct {
	widget.BaseWidget
	TableOpts *TableOpts
	Header    *widget.Table
	Data      *widget.Table
}

type BindingConverter func(interface{}) string

func NewHeaderTable(tableOpts *TableOpts) *HeaderTable {
	t := &HeaderTable{
		TableOpts: tableOpts,
		Header: widget.NewTable(
			// Dimensions (rows, cols)
			func() (int, int) { return 1, len(tableOpts.ColAttrs) },
			// Default value
			func() fyne.CanvasObject { return widget.NewLabel("the content") },
			// Cell values
			func(cellID widget.TableCellID, o fyne.CanvasObject) {
				l := o.(*widget.Label)
				opts := tableOpts.ColAttrs[cellID.Col]
				l.TextStyle = opts.HeaderStyle.TextStyle
				l.Alignment = opts.HeaderStyle.Alignment
				l.Wrapping = opts.HeaderStyle.Wrapping
				l.SetText(opts.Header)
			},
		),
		Data: widget.NewTable(dataTableLengthFunc(tableOpts), dataTableCreateFunc, dataTableUpdateFunc(tableOpts)),
	}
	t.ExtendBaseWidget(t)

	// Set Column widths
	refWidth := widget.NewLabel(t.TableOpts.RefWidth).MinSize().Width
	for i, colAttr := range t.TableOpts.ColAttrs {
		if t.Data != nil {
			t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
		}
		if t.Header != nil {
			t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
		}
	}

	return t
}

// ****************** Renderer *******************************

var _ fyne.WidgetRenderer = headerTableRenderer{}

type headerTableRenderer struct {
	headerTable *HeaderTable
	container   *fyne.Container
}

func (h *HeaderTable) CreateRenderer() fyne.WidgetRenderer {
	return headerTableRenderer{
		headerTable: h,
		container:   container.NewBorder(h.Header, nil, nil, nil, h.Data),
		// container:   container.NewVBox(h.Header, h.Data),
	}
}

func (r headerTableRenderer) MinSize() fyne.Size {
	dataMinSize := fyne.NewSize(0, 0)
	if r.headerTable.Data != nil {
		dataMinSize = r.headerTable.Data.MinSize()
	}
	headerMinSize := fyne.NewSize(0, 0)
	if r.headerTable.Header != nil {
		dataMinSize = r.headerTable.Header.MinSize()
	}
	return fyne.NewSize(
		float32(math.Max(float64(dataMinSize.Width), float64(headerMinSize.Width))),
		dataMinSize.Height+headerMinSize.Height)
}

func (r headerTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r headerTableRenderer) Destroy() {
}

func (r headerTableRenderer) Refresh() {
	r.container.Refresh()
}

func (r headerTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}
