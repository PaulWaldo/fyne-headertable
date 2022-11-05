package headertable

import (
	"fmt"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewSortLabelHeaderCellMeta(t *testing.T) {
	to := TableOpts{}
	m := NewSortLabelHeaderCellMeta(&to)
	assert.NotNil(t, m.TableOpts())
}

func Test_stringSort(t *testing.T) {
	type sample struct {
		Name string
		Num  int
	}
	data := []sample{
		{Name: "name2", Num: 0},
		{Name: "name1", Num: 1},
		{Name: "name0", Num: 2},
	}
	bindings := make([]binding.DataMap, len(data))
	colAttrs := []ColAttr{{Name: "Name"},{Name: "Num"}}
	for i := 0; i < len(data); i++ {
		bindings[i] = binding.BindStruct(&data[i])
	}
	to := TableOpts{Bindings: bindings, ColAttrs: colAttrs}
	sortFn := stringSort(&to, 0)

	sortFn(true)

	for i := 0; i < len(bindings); i++ {
		b, err := bindings[i].GetItem("Name")
		assert.NoError(t, err)
		val, err := b.(binding.String).Get()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("name%d", i), val)
	}
}

func Test_sortLabelHeaderCellMeta_NewHeader(t *testing.T) {
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
	m := NewSortLabelHeaderCellMeta(&to)
	h := m.NewHeader()

	rows, cols := h.Table.Length()
	assert.Equal(t, 1, rows, "Expecting 1 row, got %d", rows)
	assert.Equal(t, len(to.ColAttrs), cols, "Expecting %d cols, got %d", len(to.ColAttrs), cols)

	template := h.Table.CreateCell()
	assert.IsTypef(t, &SortingLabel{}, template, "Expecting type %T, got %T", widget.Label{}, template)

	sl:= template.(*SortingLabel)
	for i := range to.ColAttrs{
		h.Table.UpdateCell(widget.TableCellID{Row: 0, Col: i}, template)
		assert.IsTypef(t, &SortingLabel{}, template, "Expecting type %T, got %T", widget.Label{}, template)
		assert.Equal(t, to.ColAttrs[i].Header, sl.Label.Text)
		assert.Equal(t, to.ColAttrs[i].Alignment, sl.Label.Alignment)
		assert.Equal(t, to.ColAttrs[i].TextStyle, sl.Label.TextStyle)
		assert.Equal(t, to.ColAttrs[i].Wrapping, sl.Label.Wrapping)
	}
}

func Test_sortLabelHeaderCellMeta_SetDataTable(t *testing.T) {
}

func Test_sortLabelHeaderCellMeta_UpdateDataTable(t *testing.T) {
}

func Test_sortLabelHeaderCellMeta_TableOpts(t *testing.T) {
}
