package report

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

//Generate func create report file
func Generate(title string, query string, headers string) string {
	fmt.Println("generando reporte")

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("detalle")
	if err != nil {
		log.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "hola!"
	err = file.Save("name-file.xlsx")
	if err != nil {
		log.Printf(err.Error())
	}

	return "path_file"
}
