package logsPage

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"lore-fetcher/cmd/ui/tui/tuiHelpers"
	"lore-fetcher/internal/logger"
)

func GetLogsPage(lg *logger.Logger, pages *tview.Pages, app *tview.Application, menu *tview.List) {
	textView := tview.NewTextView().
		SetScrollable(true).
		SetWrap(true)

	refresh := func() {
		entries := lg.Entries()
		textView.SetText(strings.Join(entries, "\n"))
		textView.ScrollToEnd()
	}

	lg.SetOnChange(func() {
		app.QueueUpdateDraw(refresh)
	})

	refresh()

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			lg.SetOnChange(nil)
			pages.SwitchToPage("list")
			app.SetFocus(menu)
		}
		return event
	})

	flex := tuiHelpers.GetPageFlex(textView, "Logs", "Navigate with arrows, go back with Escape")
	app.SetFocus(textView)
	pages.AddPage("logs", flex, true, true)
}
