package tui

import (
	"github.com/rivo/tview"
	"lore-fetcher/cmd/ui/tui/jobsPage"
	"lore-fetcher/cmd/ui/tui/logsPage"
	"lore-fetcher/cmd/ui/tui/patchesPage"
	"lore-fetcher/cmd/ui/tui/tuiHelpers"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/internal/logger"
)

func RenderTuiMenu(dbsvc database.DatabaseService, lg *logger.Logger) {
	var pages *tview.Pages
	var list *tview.List

	app := tview.NewApplication()
	list = tview.NewList().
		AddItem("List patch history", "see the latest patches from the selected mailing list", '0', func() {
			patchesPage.GetPatchesPage(dbsvc, pages, app, list)
		}).
		AddItem("List job history", "see the latest CI jobs that were triggered", '1', func() {
			jobsPage.GetJobsPage(dbsvc, pages, app, list)
		}).
		AddItem("Logs", "watch live application logs", '2', func() {
			logsPage.GetLogsPage(lg, pages, app, list)
		}).
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
