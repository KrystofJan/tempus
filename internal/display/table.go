package display

import (
	"time"

	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Table struct {
	headerFmt *color.Color
	columnFmt *color.Color
	Table     *table.Table
}

func NewTable(tbl table.Table) *Table {
	headerFmt := color.New(color.FgGreen, color.Underline)
	columnFmt := color.New(color.FgYellow)
	tbl.
		WithHeaderFormatter(headerFmt.SprintfFunc()).
		WithFirstColumnFormatter(columnFmt.SprintfFunc())
	return &Table{
		Table:     &tbl,
		headerFmt: headerFmt,
		columnFmt: columnFmt,
	}
}

func PrintTasks(tasks []repository.Task) {
	tbl := NewTable(table.New(
		"ID",
		"Name",
		"StartTime",
		"Finished",
		"RecordedTime",
	))
	table := *tbl.Table
	for _, task := range tasks {
		startTime := time.Unix(task.StartTimestamp, 0)
		d := time.Duration(task.RecordedTime.Int64) * time.Second
		table.AddRow(
			task.ID,
			task.Name,
			startTime.Format(time.Kitchen),
			task.Finished,
			d.String(),
		)
	}
	table.Print()
}

func PrintEntries(entries []repository.Entry) {
	tbl := NewTable(table.New(
		"ID",
		"taskId",
		"StartTime",
		"Finished",
		"RecordedTime",
	))
	table := *tbl.Table
	for _, entry := range entries {
		startTime := time.Unix(entry.StartTimestamp, 0)
		d := time.Duration(entry.RecordedTime.Int64) * time.Second
		table.AddRow(
			entry.ID,
			entry.TaskID,
			startTime.Format(time.Kitchen),
			entry.Finished,
			d.String(),
		)
	}
	table.Print()
}
