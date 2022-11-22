package headertable

import (
	"fmt"
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestSortingLabel_OnTapped_CallsSorterAndUnsortsOthers(t *testing.T) {
	to := TableOpts{ColAttrs: []ColAttr{
		{Name: "col1", Header: "header1"},
		{Name: "col2", Header: "header2"},
	}}
	m := NewSortingHeaderTable(&to)
	m.Refresh()
	numLabels := len(to.ColAttrs)

	// Setup all the labels in the table
	labels := m.sortLabels
	sortFnCalled := make([]bool, numLabels)
	for i := range to.ColAttrs {
		m.Header.CreateCell()
		m.Header.UpdateCell(widget.TableCellID{Row: 0, Col: i}, labels[i])
		// labels[i] = m.CreateCell().(*sortingLabel)
		// h.UpdateCell(widget.TableCellID{Row: 0, Col: i}, labels[i])
		i := i
		labels[i].Sorter = func(ascending bool) {
			sortFnCalled[i] = true
		}
	}

	// For each label, tap the sort button, making sure that all other labels are unsorted
	for i := range m.sortLabels {
		test.Tap(labels[i].Button)
		assert.Equal(t, true, sortFnCalled[i])
		for j := range labels {
			if i == j {
				continue
			}
			assert.Equal(t, SortUnsorted, labels[j].State)
		}
	}
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
	bindings := make([]binding.Struct, len(data))
	colAttrs := []ColAttr{{Name: "Name"}, {Name: "Num"}}
	for i := 0; i < len(data); i++ {
		bindings[i] = binding.BindStruct(&data[i])
	}
	to := TableOpts{Bindings: bindings, ColAttrs: colAttrs}
	sortFn := stringSort(&to, 0)

	// Test Ascending
	sortFn(true)
	for i := 0; i < len(bindings); i++ {
		b, err := bindings[i].GetItem("Name")
		assert.NoError(t, err)
		val, err := b.(binding.String).Get()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("name%d", i), val)
	}

	// Test Descending
	sortFn(false)
	for i := len(bindings) - 1; i == 0; i-- {
		b, err := bindings[i].GetItem("Name")
		assert.NoError(t, err)
		val, err := b.(binding.String).Get()
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("name%d", i), val)
	}
}
