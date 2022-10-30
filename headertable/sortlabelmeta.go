package headertable

import (
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var _ HeaderCellMeta = (*sortLabelHeaderCellMeta)(nil)

type sortLabelHeaderCellMeta struct {
	tableOpts  *TableOpts
	DataTable  *widget.Table
	sortLabels []*SortingLabel
}

func NewSortLabelHeaderCellMeta(tableOpts *TableOpts) HeaderCellMeta {
	return &sortLabelHeaderCellMeta{tableOpts: tableOpts, sortLabels: make([]*SortingLabel, len(tableOpts.ColAttrs))}
}

func stringSort(tableOpts *TableOpts, col int) SortFn {
	return func(ascending bool) {
		log.Printf("Request to sort column %d ascending: %t\n", col, ascending)
		bindings := tableOpts.Bindings
		sort.Slice(bindings, func(i int, j int) bool {
			b1 := bindings[i]
			b2 := bindings[j]
			d1, err := b1.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				log.Fatalln(err)
			}
			d2, err := b2.GetItem(tableOpts.ColAttrs[col].Name)
			if err != nil {
				log.Fatalln(err)
			}
			str1, err := d1.(binding.String).Get()
			if err != nil {
				log.Fatalln(err)
			}
			str2, err := d2.(binding.String).Get()
			if err != nil {
				log.Fatalln(err)
			}
			if ascending {
				return str1 < str2
			} else {
				return str1 > str2
			}
		})
	}
}

func (m sortLabelHeaderCellMeta) NewHeader(
	length func() (int, int),
	create func() fyne.CanvasObject,
	update func(widget.TableCellID, fyne.CanvasObject),
) *Header {
	h := &Header{Table: widget.Table{Length: length, CreateCell: create, UpdateCell: update}}
	h.ExtendBaseWidget(h)
	return h
}

func (m *sortLabelHeaderCellMeta) SetDataTable(t *widget.Table) {
	m.DataTable = t
}

func (m *sortLabelHeaderCellMeta) UpdateDataTable() {
	m.DataTable.Refresh()
}

func (m *sortLabelHeaderCellMeta) LengthFn() func() (int, int) {
	return func() (int, int) { return 1, len(m.tableOpts.ColAttrs) }
}

func (m *sortLabelHeaderCellMeta) CreateCellFn() func() fyne.CanvasObject {
	return func() fyne.CanvasObject { return NewSortingLabel("the content", func() {}) }
}

func (m *sortLabelHeaderCellMeta) UpdateCellFn() func(cellID widget.TableCellID, o fyne.CanvasObject) {
	return func(cellID widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*SortingLabel)
		m.sortLabels[cellID.Col] = l
		col := cellID.Col
		opts := m.tableOpts.ColAttrs[col]
		l.Sorter = stringSort(m.TableOpts(), col)
		l.OnAfterSort = func() {
			m.UpdateDataTable()
			// Set all but this column to unsorted
			for i, sl := range m.sortLabels {
				if i != cellID.Col {
					sl.SetState(SortUnsorted)
				}
			}
		}
		l.Col = col
		l.Label.SetText(opts.Header)
		l.Label.TextStyle = opts.TextStyle
		l.Label.Alignment = opts.Alignment
		l.Label.Wrapping = opts.Wrapping
		l.Refresh()
	}
}

func (m *sortLabelHeaderCellMeta) TableOpts() *TableOpts {
	return m.tableOpts
}

// package headertable

// import (
// 	"log"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/data/binding"
// 	"fyne.io/fyne/v2/widget"
// )

// type SortState int

// const (
// 	SortUnsorted SortState = iota
// 	SortAscending
// 	SortDesending
// )

// var _ fyne.Widget = (*SortingLabel)(nil)

// type SortingLabel struct {
// 	// widget.BaseWidget
// 	State     SortState
// 	widget.Label
// 	OnSort    func()
// 	widget.Button
// 	Container *fyne.Container
// 	IsSortCol binding.Bool
// }

// func NewSortingLabel(text string, sortFunc func()) *SortingLabel {
// 	sl := &SortingLabel{
// 		Label:     *widget.NewLabel(text),
// 		Button:    *widget.NewButton("", func() {}),
// 		State:     SortUnsorted,
// 		IsSortCol: binding.NewBool(),
// 	}
// 	sl.IsSortCol.Set(false)
// 	sl.SetState(SortUnsorted)
// 	// sl.Button.OnTapped = sl.onSortTapped
// 	// sl.updateIcon()
// 	sl.Container = container.NewHBox(sl.Label, sl.Button)

// 	sl.ExtendBaseWidget(sl)
// 	return sl
// }

// func (s *SortingLabel) SetState(state SortState) {
// 	s.State = state
// 	switch s.State {
// 	case SortUnsorted:
// 		// s.Button.SetIcon(theme.ListIcon())
// 		s.Button.SetIcon(data.IconSortSvg)
// 		s.IsSortCol.Set(false)
// 	case SortAscending:
// 		// s.Button.SetIcon(theme.MoveDownIcon())
// 		s.Button.SetIcon(data.IconSortDownSvg)
// 		s.IsSortCol.Set(true)
// 		// s.Button.OnTapped()
// 	case SortDesending:
// 		// s.Button.SetIcon(theme.MoveUpIcon())
// 		s.Button.SetIcon(data.IconSortUpSvg)
// 		s.IsSortCol.Set(true)
// 		// s.Button.OnTapped()
// 	default:
// 		log.Fatalf("Unknown sort label state: %d", s.State)
// 	}
// 	s.Button.Refresh()
// }

// var _ fyne.WidgetRenderer = (*sortingLabelRenderer)(nil)

// type sortingLabelRenderer struct {
// 	sortLabel *SortingLabel
// 	container *fyne.Container
// }

// func (sl *SortingLabel) CreateRenderer() fyne.WidgetRenderer {
// 	sl.Container= container.NewHBox(sl.Label, sl.Button)
// 	return &sortingLabelRenderer{
// 		sortLabel: sl,
// 		container: container.NewHBox(sl.Label, sl.Button),
// 	}
// }

// func (r *sortingLabelRenderer) MinSize() fyne.Size {
// 	return r.sortLabel.Container.MinSize()
// }

// func (r *sortingLabelRenderer) Layout(size fyne.Size) {
// 	r.sortLabel.Container.Resize(size)
// }

// func (r *sortingLabelRenderer) Refresh() {
// 	r.sortLabel.Container.Refresh()
// }

// func (r *sortingLabelRenderer) Objects() []fyne.CanvasObject {
// 	return []fyne.CanvasObject{r.sortLabel.Container}
// }

// func (r *sortingLabelRenderer) Destroy() {}
