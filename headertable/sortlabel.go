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

// var _ fyne.Widget = (*SortingLabel)(nil)

type SortFn func(ascending bool)

type SortingLabel struct {
	widget.BaseWidget
	State  SortState
	Label  *widget.Label
	Sorter SortFn
	OnAfterSort    func()
	Button *widget.Button
	Col    int
	// Container *fyne.Container
	// IsSortCol binding.Bool
}

func NewSortingLabel(text string, sortFunc func()) *SortingLabel {
	sl := &SortingLabel{
		Label:  widget.NewLabel(text),
		Button: widget.NewButton("", func() {}),
		State:  SortUnsorted,
		// IsSortCol: binding.NewBool(),
	}
	sl.SetState(SortUnsorted)
	sl.Button.OnTapped = sl.OnTapped

	sl.ExtendBaseWidget(sl)
	return sl
}

func (s *SortingLabel) SetState(state SortState) {
	s.State = state
	switch s.State {
	case SortUnsorted:
		// s.Button.SetIcon(theme.ListIcon())
		s.Button.SetIcon(data.IconSortSvg)
		// s.IsSortCol.Set(false)
	case SortAscending:
		// s.Button.SetIcon(theme.MoveDownIcon())
		s.Button.SetIcon(data.IconSortDownSvg)
		// s.IsSortCol.Set(true)
		// s.Button.OnTapped()
	case SortDescending:
		// s.Button.SetIcon(theme.MoveUpIcon())
		s.Button.SetIcon(data.IconSortUpSvg)
		// s.IsSortCol.Set(true)
		// s.Button.OnTapped()
	default:
		log.Fatalf("Unknown sort label state: %d", s.State)
	}
	s.Button.Refresh()
}

func (s *SortingLabel) nextState() SortState {
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

func (s *SortingLabel) OnTapped() {
	s.SetState(s.nextState())
	s.Sorter(s.State == SortAscending)
	s.OnAfterSort()
}

func (sl *SortingLabel) CreateRenderer() fyne.WidgetRenderer {
	return &sortingLabelRenderer{
		sortLabel: sl,
		container: container.NewHBox(sl.Label, sl.Button),
	}
}

var _ fyne.WidgetRenderer = (*sortingLabelRenderer)(nil)

type sortingLabelRenderer struct {
	sortLabel *SortingLabel
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
