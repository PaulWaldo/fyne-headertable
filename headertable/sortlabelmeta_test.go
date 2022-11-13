package headertable

// import (
// 	"fmt"
// 	"testing"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/data/binding"
// 	"fyne.io/fyne/v2/test"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/PaulWaldo/fyne-headertable/headertable/data"
// 	"github.com/stretchr/testify/assert"
// )

// // func TestNewSortLabelHeaderCellMeta(t *testing.T) {
// // 	to := TableOpts{}
// // 	m := NewSortLabelHeaderCellMeta(&to)
// // 	assert.NotNil(t, m.TableOpts())
// // }

// func Test_stringSort(t *testing.T) {
// 	type sample struct {
// 		Name string
// 		Num  int
// 	}
// 	data := []sample{
// 		{Name: "name2", Num: 0},
// 		{Name: "name1", Num: 1},
// 		{Name: "name0", Num: 2},
// 	}
// 	bindings := make([]binding.DataMap, len(data))
// 	colAttrs := []ColAttr{{Name: "Name"}, {Name: "Num"}}
// 	for i := 0; i < len(data); i++ {
// 		bindings[i] = binding.BindStruct(&data[i])
// 	}
// 	to := TableOpts{Bindings: bindings, ColAttrs: colAttrs}
// 	sortFn := stringSort(&to, 0)

// 	// Test Ascending
// 	sortFn(true)
// 	for i := 0; i < len(bindings); i++ {
// 		b, err := bindings[i].GetItem("Name")
// 		assert.NoError(t, err)
// 		val, err := b.(binding.String).Get()
// 		assert.NoError(t, err)
// 		assert.Equal(t, fmt.Sprintf("name%d", i), val)
// 	}

// 	// Test Descending
// 	sortFn(false)
// 	for i := len(bindings) - 1; i == 0; i-- {
// 		b, err := bindings[i].GetItem("Name")
// 		assert.NoError(t, err)
// 		val, err := b.(binding.String).Get()
// 		assert.NoError(t, err)
// 		assert.Equal(t, fmt.Sprintf("name%d", i), val)
// 	}
// }

// func Test_sortLabelHeaderCellMeta_NewHeader(t *testing.T) {
// 	to := TableOpts{ColAttrs: []ColAttr{
// 		{
// 			Name:         "col1",
// 			Header:       "header1",
// 			Alignment:    fyne.TextAlignLeading,
// 			TextStyle:    fyne.TextStyle{Bold: true},
// 			Wrapping:     fyne.TextWrapOff,
// 			WidthPercent: 50,
// 		},
// 		{
// 			Name:         "col2",
// 			Header:       "header2",
// 			Alignment:    fyne.TextAlignTrailing,
// 			TextStyle:    fyne.TextStyle{Italic: true},
// 			Wrapping:     fyne.TextWrapBreak,
// 			WidthPercent: 25,
// 		},
// 	}}
// 	h := NewHeaderTable(&to)

// 	test.NewApp()
// 	rows, cols := h.Header.Length()
// 	assert.Equal(t, 1, rows, "Expecting 1 row, got %d", rows)
// 	assert.Equal(t, len(to.ColAttrs), cols, "Expecting %d cols, got %d", len(to.ColAttrs), cols)

// 	template := h.Header.CreateCell()
// 	assert.IsTypef(t, &sortingLabel{}, template, "Expecting type %T, got %T", widget.Label{}, template)

// 	sl := template.(*sortingLabel)
// 	for i := range to.ColAttrs {
// 		h.Header.UpdateCell(widget.TableCellID{Row: 0, Col: i}, template)
// 		assert.IsTypef(t, &sortingLabel{}, template, "Expecting type %T, got %T", widget.Label{}, template)
// 		assert.Equal(t, to.ColAttrs[i].Header, sl.Label.Text)
// 		assert.Equal(t, to.ColAttrs[i].Alignment, sl.Label.Alignment)
// 		assert.Equal(t, to.ColAttrs[i].TextStyle, sl.Label.TextStyle)
// 		assert.Equal(t, to.ColAttrs[i].Wrapping, sl.Label.Wrapping)
// 	}
// }

// func TestSortingLabel_OnTapped_CyclesSortStates(t *testing.T) {
// 	sl := NewSortingLabel("some text")
// 	assert.Equal(t, SortUnsorted, sl.State)
// 	assert.Equal(t, data.IconSortSvg.Name(), sl.Button.Icon.Name())
// 	test.Tap(sl.Button)
// 	assert.Equal(t, SortAscending, sl.State)
// 	assert.Equal(t, data.IconSortDownSvg.Name(), sl.Button.Icon.Name())
// 	test.Tap(sl.Button)
// 	assert.Equal(t, SortDescending, sl.State)
// 	assert.Equal(t, data.IconSortUpSvg.Name(), sl.Button.Icon.Name())
// 	test.Tap(sl.Button)
// 	assert.Equal(t, SortAscending, sl.State)
// 	assert.Equal(t, data.IconSortDownSvg.Name(), sl.Button.Icon.Name())
// }

// func TestSortingLabel_OnTapped_CallsSorterAndUnsortsOthers(t *testing.T) {
// 	to := TableOpts{ColAttrs: []ColAttr{
// 		{Name: "col1", Header: "header1"},
// 		{Name: "col2", Header: "header2"},
// 	}}
// 	m := NewSortLabelHeaderCellMeta(&to)
// 	m.SetDataTable(&widget.Table{})
// 	h := m.NewHeader()
// 	numLabels := len(to.ColAttrs)

// 	// Setup all the labels in the table
// 	labels := make([]*sortingLabel, numLabels)
// 	sortFnCalled := make([]bool, numLabels)
// 	for i := range to.ColAttrs {
// 		labels[i] = h.CreateCell().(*sortingLabel)
// 		h.UpdateCell(widget.TableCellID{Row: 0, Col: i}, labels[i])
// 		i := i
// 		labels[i].Sorter = func(ascending bool) {
// 			sortFnCalled[i] = true
// 		}
// 	}

// 	// For each label, tap the sort button, making sure that all other labels are unsorted
// 	for i := range labels {
// 		l := labels[i]
// 		test.Tap(l.Button)
// 		assert.Equal(t, true, sortFnCalled[i])
// 		for j := range labels {
// 			if i == j {
// 				continue
// 			}
// 			assert.Equal(t, SortUnsorted, labels[j].State)
// 		}
// 	}
// }
