package headertable

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
)

func TestNewHeaderTable(t *testing.T) {
	// type args struct {
	// 	meta HeaderCellMeta
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want *HeaderTable
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := NewHeaderTable(tt.args.meta); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("NewHeaderTable() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestHeaderTable_CreateRenderer(t *testing.T) {
	tests := []struct {
		name string
		h    *HeaderTable
		want fyne.WidgetRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.CreateRenderer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeaderTable.CreateRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerTableRenderer_MinSize(t *testing.T) {
	tests := []struct {
		name string
		r    headerTableRenderer
		want fyne.Size
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MinSize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headerTableRenderer.MinSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerTableRenderer_Layout(t *testing.T) {
	type args struct {
		s fyne.Size
	}
	tests := []struct {
		name string
		r    headerTableRenderer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Layout(tt.args.s)
		})
	}
}

func Test_headerTableRenderer_Destroy(t *testing.T) {
	tests := []struct {
		name string
		r    headerTableRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Destroy()
		})
	}
}

func Test_headerTableRenderer_Refresh(t *testing.T) {
	tests := []struct {
		name string
		r    headerTableRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Refresh()
		})
	}
}

func Test_headerTableRenderer_Objects(t *testing.T) {
	tests := []struct {
		name string
		r    headerTableRenderer
		want []fyne.CanvasObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Objects(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headerTableRenderer.Objects() = %v, want %v", got, tt.want)
			}
		})
	}
}
