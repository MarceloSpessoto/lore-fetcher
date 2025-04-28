package tuiHelpers

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func GetPageFlex(content tview.Primitive, boxText string, helpText string) tview.Primitive {
	frame := tview.NewFrame(content).
						AddText(boxText, true, tview.AlignCenter, tcell.ColorWhite).
						AddText(helpText, false, tview.AlignCenter, tcell.ColorWhite).
						SetBorders(2, 2, 2, 2, 4, 4)

	flex := tview.NewFlex().
						AddItem(tview.NewBox(), 0, 1, false).
						AddItem(frame, 0, 1, false).
						AddItem(tview.NewBox(), 0, 1, false)

	return flex
}

type PageManager struct {
	pages tview.Pages
	app tview.Application
	focusList map[string]*tview.Primitive
}

func NewPageManager(app tview.Application) *PageManager {
	return &PageManager{
		pages:     *tview.NewPages(),
		app:       app,
		focusList: make(map[string]*tview.Primitive),
	}
}

func (pm *PageManager) AddPage(name string, page tview.Primitive) {
	pm.pages.AddPage(name, page, true, true)
	pm.focusList[name] = &page
}

func (pm *PageManager) ChangePage(name string) {
	focus := pm.focusList[name]
	pm.pages.SwitchToPage(name)
	pm.app.SetFocus(*focus)
}

func (pm *PageManager) GetPages() *tview.Pages {
	return &pm.pages
}
