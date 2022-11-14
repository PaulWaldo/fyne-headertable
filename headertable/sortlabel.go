package headertable

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/fyne-headertable/headertable/data"

	"fyne.io/fyne/v2/widget"
)

type SortState int

const (
	SortUnsorted SortState = iota
	SortAscending
	SortDescending
)

var _ fyne.Widget = (*sortingLabel)(nil)

type SortFn func(ascending bool)

type sortingLabel struct {
	widget.BaseWidget
	State       SortState
	Label       *widget.Label
	Sorter      SortFn
	OnAfterSort func()
	Button      *widget.Button
	Col         int
}

func NewSortingLabel(text string) *sortingLabel {
	sl := &sortingLabel{
		Label:  widget.NewLabel(text),
		Button: widget.NewButton("", func() {}),
		State:  SortUnsorted,
	}
	sl.SetState(SortUnsorted)
	sl.Button.OnTapped = sl.OnTapped

	sl.ExtendBaseWidget(sl)
	return sl
}

func (s *sortingLabel) SetState(state SortState) {
	s.State = state
	switch s.State {
	case SortUnsorted:
		s.Button.SetIcon(data.IconSortSvg)
	case SortAscending:
		s.Button.SetIcon(data.IconSortDownSvg)
	case SortDescending:
		s.Button.SetIcon(data.IconSortUpSvg)
	default:
		log.Fatalf("Unknown sort label state: %d", s.State)
	}
	s.Button.Refresh()
}

func (s *sortingLabel) nextState() SortState {
	switch s.State {
	case SortUnsorted:
		return SortAscending
	case SortDescending:
		return SortAscending
	case SortAscending:
		return SortDescending
	default:
		log.Printf("error checking nextState: current state invalid: %v\n", s.State)
		return SortUnsorted
	}
}

func (s *sortingLabel) OnTapped() {
	s.SetState(s.nextState())
	if s.Sorter != nil {
		s.Sorter(s.State == SortAscending)
	}
	if s.OnAfterSort != nil {
		s.OnAfterSort()
	}
}

func (sl *sortingLabel) CreateRenderer() fyne.WidgetRenderer {
	return &sortingLabelRenderer{
		sortLabel: sl,
		container: container.NewHBox(sl.Label, sl.Button),
	}
}

var _ fyne.WidgetRenderer = (*sortingLabelRenderer)(nil)

type sortingLabelRenderer struct {
	sortLabel *sortingLabel
	container *fyne.Container
}

func (r *sortingLabelRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *sortingLabelRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *sortingLabelRenderer) Refresh() {
	r.container.Refresh()
}

func (r *sortingLabelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *sortingLabelRenderer) Destroy() {}
