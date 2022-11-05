package headertable

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewLabelHeaderCellMeta(t *testing.T) {
	to := TableOpts{}
	m := NewLabelHeaderCellMeta(&to)
	assert.NotNil(t, m.TableOpts())
}

func Test_labelHeaderCellMeta_NewHeader(t *testing.T) {
	to := TableOpts{ColAttrs: []ColAttr{
		{
			Name: "col1",
			Header:"header1",
			Alignment: fyne.TextAlignLeading,
			TextStyle: fyne.TextStyle{Bold: true},
			Wrapping: fyne.TextWrapOff,
			WidthPercent: 50,
		},
		{
			Name: "col2",
			Header:"header2",
			Alignment: fyne.TextAlignTrailing,
			TextStyle: fyne.TextStyle{Italic: true},
			Wrapping: fyne.TextWrapBreak,
			WidthPercent: 25,
		},
	}}
	m := NewLabelHeaderCellMeta(&to)
	h := m.NewHeader()

	rows, cols := h.Table.Length()
	assert.Equal(t, 1, rows, "Expecting 1 row, got %d", rows)
	assert.Equal(t, len(to.ColAttrs), cols, "Expecting %d cols, got %d", len(to.ColAttrs), cols)

	template := h.Table.CreateCell()
	assert.IsTypef(t, &widget.Label{}, template, "Expecting type %T, got %T", widget.Label{}, template)

	label:= template.(*widget.Label)
	for i := range to.ColAttrs{
		h.Table.UpdateCell(widget.TableCellID{Row: 0, Col: i}, template)
		assert.IsTypef(t, &widget.Label{}, template, "Expecting type %T, got %T", widget.Label{}, template)
		assert.Equal(t, to.ColAttrs[i].Header, label.Text)
		assert.Equal(t, to.ColAttrs[i].Alignment, label.Alignment)
		assert.Equal(t, to.ColAttrs[i].TextStyle, label.TextStyle)
		assert.Equal(t, to.ColAttrs[i].Wrapping, label.Wrapping)
	}
}

func Test_labelHeaderCellMeta_UpdateDataTable(t *testing.T) {
	m := NewLabelHeaderCellMeta(&TableOpts{})
	assert.Panics(t, m.UpdateDataTable)
}

func Test_labelHeaderCellMeta_SetDataTable(t *testing.T) {
	m := NewLabelHeaderCellMeta(&TableOpts{})
	table := widget.NewTable(nil, nil, nil)
	m.SetDataTable(table)
}

func Test_labelHeaderCellMeta_TableOpts(t *testing.T) {
	to := &TableOpts{ColAttrs: []ColAttr{}}
	m := NewLabelHeaderCellMeta(to)
	assert.Equal(t, to, m.TableOpts())
}
