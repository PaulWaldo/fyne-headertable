package headertable

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

// func TestNewHeaderTable(t *testing.T) {
// }

// func TestHeaderTable_CreateRenderer(t *testing.T) {
// }

// func Test_headerTableRenderer_MinSize(t *testing.T) {
// }

// func Test_headerTableRenderer_Layout(t *testing.T) {
// }

// func Test_headerTableRenderer_Destroy(t *testing.T) {
// }

// func Test_headerTableRenderer_Refresh(t *testing.T) {
// }

// func Test_headerTableRenderer_Objects(t *testing.T) {
// }

func TestHeaderTable(t *testing.T) {
	type carInfo struct {
		Year, Make, Model string
	}
	cars := []carInfo{
		{Year: "1980", Make: "Toyota", Model: "Corolla"},
		{Year: "2020", Make: "Toyota", Model: "Corolla"},
		{Year: "2020", Make: "Ford", Model: "Mustang"},
		{Year: "2022", Make: "Tesla", Model: "Model X"},
	}
	bindings := make([]binding.Struct, len(cars))
	for i := range cars {
		bindings[i] = binding.BindStruct(&cars[i])
	}
	opts := TableOpts{
		ColAttrs: []ColAttr{
			{Name: "Year", Header: "The Year", WidthPercent: 25, HeaderStyle: CellStyle{TextStyle: fyne.TextStyle{Bold: true}}},
			{Name: "Make", Header: "The Make", WidthPercent: 50},
			{Name: "Model", Header: "The Model", WidthPercent: 75, HeaderStyle: CellStyle{Wrapping: fyne.TextTruncate}},
		},
		RefWidth: "I am prototypical",
		Bindings: bindings,
	}
	test.NewApp()
	ht := NewHeaderTable(&opts)

	rows, cols := ht.Header.Length()
	assert.Equal(t, 3, cols, "Expecting %d cols, got %d", 3, cols)
	assert.Equal(t, 1, rows, "Expecting %d rows, got %d", 1, rows)

	// Test that the headers are as expected
	// Create a template and see what the Header's UpdateCell callback transforms it to
	for i := range opts.ColAttrs {
		template := ht.Header.CreateCell().(*widget.Label)
		ht.Header.UpdateCell(widget.TableCellID{Row: 0, Col: i}, template)
		assert.Equal(t, opts.ColAttrs[i].Header, template.Text)
		assert.Equal(t, opts.ColAttrs[i].HeaderStyle.Alignment, template.Alignment)
		assert.Equal(t, opts.ColAttrs[i].HeaderStyle.TextStyle, template.TextStyle)
		assert.Equal(t, opts.ColAttrs[i].HeaderStyle.Wrapping, template.Wrapping)
	}

	// Test that the data table contains expected values
	for i := range cars {
		template := ht.Data.CreateCell().(*widget.Label)
		ht.Data.UpdateCell(widget.TableCellID{Row: i, Col: 0}, template)
		assert.Equal(t, cars[i].Year, template.Text)
		ht.Data.UpdateCell(widget.TableCellID{Row: i, Col: 1}, template)
		assert.Equal(t, cars[i].Make, template.Text)
		ht.Data.UpdateCell(widget.TableCellID{Row: i, Col: 2}, template)
		assert.Equal(t, cars[i].Model, template.Text)
	}
}
