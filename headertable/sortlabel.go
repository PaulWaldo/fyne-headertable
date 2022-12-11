package headertable

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/PaulWaldo/fyne-headertable/headertable/data"

	"fyne.io/fyne/v2/widget"
)

type SortState int

const (
	SortUnsorted SortState = iota
	SortAscending
	SortDescending
)

var _ fyne.Widget = (*SortingLabel)(nil)

type SortFn func(ascending bool)

type SortingLabel struct {
	widget.BaseWidget
	State       SortState
	Label       *widget.Label
	Sorter      SortFn
	OnAfterSort func()
	Button      *widget.Button
	Col         int
}

func NewSortingLabel(text string) *SortingLabel {
	sl := &SortingLabel{
		Label:  widget.NewLabel(text),
		Button: widget.NewButton("", func() {}),
		State:  SortUnsorted,
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
		s.Button.SetIcon(data.IconSortSvg)
		s.Button.Importance = widget.MediumImportance
	case SortAscending:
		s.Button.SetIcon(data.IconSortDownSvg)
		s.Button.Importance = widget.HighImportance
	case SortDescending:
		s.Button.SetIcon(data.IconSortUpSvg)
		s.Button.Importance = widget.HighImportance
	default:
		log.Printf("Unknown sort label state: %d\n", s.State)
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
	if s.Sorter != nil {
		s.Sorter(s.State == SortAscending)
	}
	if s.OnAfterSort != nil {
		s.OnAfterSort()
	}
}

func (sl *SortingLabel) CreateRenderer() fyne.WidgetRenderer {
	spacer := &layout.Spacer{FixHorizontal: true}
	spacedButton := container.NewHBox(sl.Button, spacer)
	return &sortingLabelRenderer{
		sortLabel: sl,
		container: container.NewBorder(nil, nil, nil, spacedButton, sl.Label),
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

func (*sortingLabelRenderer) Destroy() {}
