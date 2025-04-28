package tui

import (
	"github.com/rivo/tview"
	"lore-fetcher/cmd/ui/tui/patchesPage"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/cmd/ui/tui/tuiHelpers"
)

func RenderTuiMenu(dbsvc database.DatabaseService) {
	var pages *tview.Pages
	var list *tview.List

	app := tview.NewApplication()
	list = tview.NewList().
		AddItem("List patch history", "see the latest patches from the selected mailing list", '0', func() {
			patchesPage.GetPatchesPage(dbsvc, pages, app, list)
		}).
		AddItem("List job history", "see the latest CI jobs that were triggered", '1', nil).
		AddItem("Quit", "Quit the app", 'q', func() {
			app.Stop()
		})

	flex := tuiHelpers.GetPageFlex(list, "Menu", "Navigate with arrows and select with Enter")

	pages = tview.NewPages().
					AddPage("list", flex, true, true)

	if err := app.SetRoot(pages, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
