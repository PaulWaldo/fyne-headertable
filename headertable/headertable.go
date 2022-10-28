package headertable

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.WidgetRenderer = headerTableRenderer{}

type ColAttr struct {
	Name         string
	Header       string
	Alignment    fyne.TextAlign
	Wrapping     fyne.TextWrap
	TextStyle    fyne.TextStyle
	WidthPercent int
}

type TableOpts struct {
	ColAttrs         []ColAttr
	RefWidth         string
	Bindings         []binding.DataMap
	OnDataCellSelect func(cellID widget.TableCellID)
}

type HeaderTable struct {
	widget.BaseWidget
	TableOpts    *TableOpts
	Header       *Header
	Data         *widget.Table
}

func NewHeaderTable(tableOpts *TableOpts) *HeaderTable {
	t := &HeaderTable{
		TableOpts:    tableOpts,
		Header: NewHeader(tableOpts),
		Data: widget.NewTable(
			// Dimensions (rows, cols)
			func() (int, int) {
				return len(tableOpts.Bindings), len(tableOpts.ColAttrs)
			},

			// Default value
			func() fyne.CanvasObject {
				return widget.NewLabel("wide content")
			},

			// Cell values
			func(cellID widget.TableCellID, cnvObj fyne.CanvasObject) {
				// str,_:=
				b := tableOpts.Bindings[cellID.Row]
				d, _ := b.GetItem(tableOpts.ColAttrs[cellID.Col].Name)
				str, _ := d.(binding.String).Get()
				l := cnvObj.(*widget.Label)
				l.SetText(str)
			},
		),
	}
	t.ExtendBaseWidget(t)

	// Set Column widths
	refWidth := widget.NewLabel(tableOpts.RefWidth).MinSize().Width
	for i, colAttr := range tableOpts.ColAttrs {
		t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
		t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*refWidth)
	}

	return t
}

//*******************************************************************************

type headerTableRenderer struct {
	headerTable *HeaderTable
	container   *fyne.Container
}

func (h *HeaderTable) CreateRenderer() fyne.WidgetRenderer {
	return headerTableRenderer{
		headerTable: h,
		container:   container.NewBorder(h.Header, nil, nil, nil, h.Data),
	}
}

func (r headerTableRenderer) MinSize() fyne.Size {
	return fyne.NewSize(
		float32(math.Max(float64(r.headerTable.Data.MinSize().Width), float64(r.headerTable.Header.MinSize().Width))),
		r.headerTable.Data.MinSize().Height+r.headerTable.Header.MinSize().Height)
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
