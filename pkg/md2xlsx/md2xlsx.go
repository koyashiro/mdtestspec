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
	sheet := "Sheet1"
	i := 0

	for _, c := range s.Categories {
		for _, sc := range c.SubCategories {
			for _, ssc := range sc.SubSubCategories {
				axis := fmt.Sprintf("%s%d", CategoryCol, i+1)
				f.SetCellValue(sheet, axis, c.Name)

				axis = fmt.Sprintf("%s%d", SubCategoryCol, i+1)
				f.SetCellValue(sheet, axis, sc.Name)

				axis = fmt.Sprintf("%s%d", SubSubCategoryCol, i+1)
				f.SetCellValue(sheet, axis, ssc.Name)

				sb := strings.Builder{}

				axis = fmt.Sprintf("%s%d", ProceduresCol, i+1)
				for j, p := range ssc.Procedures {
					if j != 0 {
						sb.WriteRune('\n')
					}
					sb.WriteString(fmt.Sprintf("%d. ", j+1))
					sb.WriteString(p)
				}
				f.SetCellValue(sheet, axis, sb.String())

				sb.Reset()

				axis = fmt.Sprintf("%s%d", ConfirmationsCol, i+1)
				for j, p := range ssc.Confirmations {
					if j != 0 {
						sb.WriteRune('\n')
					}
					sb.WriteRune('・')
					sb.WriteString(p)
				}
				f.SetCellValue(sheet, axis, sb.String())

				i++
			}
		}
	}

	return f.SaveAs(name)
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
