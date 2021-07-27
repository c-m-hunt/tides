package tides

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ts Tides) Display() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)

	r := 0
	c := 0

	highColor := tcell.ColorRed
	lowColor := tcell.ColorDarkCyan
	for i, t := range ts {
		if r == 0 {
			table.SetCell(r, c,
				tview.NewTableCell(t.DateTime.Format("Mon Jan 2")).
					SetAlign(tview.AlignCenter))
			r += 1
		}
		color := highColor
		if t.State == "Low" {
			color = lowColor
		}
		cell := fmt.Sprintf("%v - %vm", t.DateTime.Format("15:04"), t.HeightMetres)
		table.SetCell(r, c,
			tview.NewTableCell(cell).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))
		r += 1
		if i == len(ts)-1 {
			break
		}
		if t.DateTime.Day() != ts[i+1].DateTime.Day() {
			c += 1
			r = 0
		}
	}

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
