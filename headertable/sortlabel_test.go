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
		text string
	}
	tests := []struct {
		name string
		args args
		want *sortingLabel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSortingLabel(tt.args.text); !reflect.DeepEqual(got, tt.want) {
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
			s := NewSortingLabel("")
			s.SetState(tt.args.state)
			assert.Equal(t, tt.args.expectedIconName, s.Button.Icon.Name())
		})
	}
}

func TestSortingLabel_nextState(t *testing.T) {
	tests := []struct {
		name string
		s    *sortingLabel
		want SortState
	}{
		{name: "ascending->descending", s: &sortingLabel{State: SortAscending}, want: SortDescending},
		{name: "descending->ascending", s: &sortingLabel{State: SortAscending}, want: SortDescending},
		{name: "unsorted->ascending", s: &sortingLabel{State: SortUnsorted}, want: SortAscending},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.nextState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortingLabel.nextState() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestSortingLabel_CreateRenderer(t *testing.T) {
	tests := []struct {
		name string
		sl   *sortingLabel
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
