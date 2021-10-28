package md2xlsx

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Spec struct {
	Name       string
	Categories []*Category
}

func (s *Spec) SaveAs(name string) error {
	f := excelize.NewFile()
	sb := strings.Builder{}
	sheet := "Sheet1"
	i := 0

	for _, c := range s.Categories {
		for _, sc := range c.SubCategories {
			for _, ssc := range sc.SubSubCategories {
				setCategory(f, sheet, i+1, c.Name)
				setSubCategory(f, sheet, i+1, sc.Name)
				setSubSubCategory(f, sheet, i+1, ssc.Name)
				setConfirmations(f, sheet, i+1, ssc.Confirmations, &sb)
				setProcedures(f, sheet, i+1, ssc.Procedures, &sb)
				i++
			}
		}
	}

	return f.SaveAs(name)
}

func setCategory(f *excelize.File, sheet string, row int, name string) {
	axis := fmt.Sprintf("%s%d", CategoryCol, row)
	f.SetCellValue(sheet, axis, name)
}

func setSubCategory(f *excelize.File, sheet string, row int, name string) {
	axis := fmt.Sprintf("%s%d", SubCategoryCol, row)
	f.SetCellValue(sheet, axis, name)
}

func setSubSubCategory(f *excelize.File, sheet string, row int, name string) {
	axis := fmt.Sprintf("%s%d", SubSubCategoryCol, row)
	f.SetCellValue(sheet, axis, name)
}

func setProcedures(f *excelize.File, sheet string, row int, procedures []string, sb *strings.Builder) {
	sb.Reset()
	axis := fmt.Sprintf("%s%d", ProceduresCol, row)
	for j, p := range procedures {
		if j != 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(fmt.Sprintf("%d. ", j+1))
		sb.WriteString(p)
	}
	f.SetCellValue(sheet, axis, sb.String())
}

func setConfirmations(f *excelize.File, sheet string, row int, confirmations []string, sb *strings.Builder) {
	sb.Reset()
	axis := fmt.Sprintf("%s%d", ConfirmationsCol, row)
	for j, p := range confirmations {
		if j != 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(fmt.Sprintf("%d. ", j+1))
		sb.WriteString(p)
	}
	f.SetCellValue(sheet, axis, sb.String())
}

type Category struct {
	Name          string
	SubCategories []*SubCategory
}

type SubCategory struct {
	Name             string
	SubSubCategories []*SubSubCategory
}

type SubSubCategory struct {
	Name          string
	Procedures    []string
	Confirmations []string
}

const CategoryCol = "A"

const SubCategoryCol = "B"

const SubSubCategoryCol = "C"

const ProceduresCol = "D"

const ConfirmationsCol = "E"
