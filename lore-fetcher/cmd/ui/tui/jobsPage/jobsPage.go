package jobsPage

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"lore-fetcher/internal/core/services/database"
	"lore-fetcher/cmd/ui/tui/tuiHelpers"
)

func GetJobsPage(dbsvc database.DatabaseService, pages *tview.Pages, app *tview.Application, menu *tview.List) {
	table := tview.NewTable().
		SetFixed(1, 0).
		SetSeparator(tview.BoxDrawingsLightHorizontal).
		SetBordersColor(tcell.ColorYellow).
		SetSelectable(true, false)
	frame := tview.NewFrame(table).
		SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true).
		SetTitle(fmt.Sprintf(`Contents of table "%s"`, "jobs"))

	loadRows := func(offset int) {
		jobs, err := dbsvc.ReadJobs()
		if err != nil {
			panic(err)
		}

		for _, job := range jobs {
			patch, _ := dbsvc.ReadPatch(job.PatchID)
			row := table.GetRowCount()
			table.SetCellSimple(row, 0, patch.Title)
			table.SetCellSimple(row, 1, job.Status)
		}
	}

	loadRows(0)

	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			pages.SwitchToPage("list")
			app.SetFocus(menu)
		}
	})
	flex := tuiHelpers.GetPageFlex(table, "Jobs", "Navigate with arrows, select with Enter and go back with Escape")
	app.SetFocus(table)
	pages.AddPage("jobs", flex, true, true)
}

