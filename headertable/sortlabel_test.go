package headertable

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"github.com/PaulWaldo/fyne-headertable/headertable/data"
	"github.com/stretchr/testify/assert"
)

func TestNewSortingLabel(t *testing.T) {
	type args struct {
		text     string
		sortFunc func()
	}
	tests := []struct {
		name string
		args args
		want *SortingLabel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSortingLabel(tt.args.text, tt.args.sortFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSortingLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortingLabel_SetState(t *testing.T) {
	type args struct {
		state            SortState
		expectedIconName string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "ascending", args: args{
			state:            SortAscending,
			expectedIconName: data.IconSortDownSvg.Name()},
		},
		{name: "descending", args: args{
			state:            SortDescending,
			expectedIconName: data.IconSortUpSvg.Name()},
		},
		{name: "unsorted", args: args{
			state:            SortUnsorted,
			expectedIconName: data.IconSortSvg.Name()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSortingLabel("", func() {})
			s.SetState(tt.args.state)
			assert.Equal(t, tt.args.expectedIconName, s.Button.Icon.Name())
		})
	}
}

func TestSortingLabel_nextState(t *testing.T) {
	tests := []struct {
		name string
		s    *SortingLabel
		want SortState
	}{
		{name: "ascending->descending", s: &SortingLabel{State: SortAscending}, want: SortDescending},
		{name: "descending->ascending", s: &SortingLabel{State: SortAscending}, want: SortDescending},
		{name: "unsorted->ascending", s: &SortingLabel{State: SortUnsorted}, want: SortAscending},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.nextState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortingLabel.nextState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortingLabel_OnTapped(t *testing.T) {
	tests := []struct {
		name string
		s    *SortingLabel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.OnTapped()
		})
	}
}

func TestSortingLabel_CreateRenderer(t *testing.T) {
	tests := []struct {
		name string
		sl   *SortingLabel
		want fyne.WidgetRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.CreateRenderer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortingLabel.CreateRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortingLabelRenderer_MinSize(t *testing.T) {
	tests := []struct {
		name string
		r    *sortingLabelRenderer
		want fyne.Size
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MinSize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortingLabelRenderer.MinSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortingLabelRenderer_Layout(t *testing.T) {
	type args struct {
		size fyne.Size
	}
	tests := []struct {
		name string
		r    *sortingLabelRenderer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Layout(tt.args.size)
		})
	}
}

func Test_sortingLabelRenderer_Refresh(t *testing.T) {
	tests := []struct {
		name string
		r    *sortingLabelRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Refresh()
		})
	}
}

func Test_sortingLabelRenderer_Objects(t *testing.T) {
	tests := []struct {
		name string
		r    *sortingLabelRenderer
		want []fyne.CanvasObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Objects(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortingLabelRenderer.Objects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortingLabelRenderer_Destroy(t *testing.T) {
	tests := []struct {
		name string
		r    *sortingLabelRenderer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Destroy()
		})
	}
}
