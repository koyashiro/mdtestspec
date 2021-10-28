package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Spec struct {
	Name       string
	Categories []*Category
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

func main() {
	spec := exampleSpec()

	f := excelize.NewFile()
	s := "Sheet1"

	for i, c := range spec.Categories {
		axis := fmt.Sprintf("%s%d", CategoryCol, i+1)
		f.SetCellValue(s, axis, c.Name)

		for _, sc := range c.SubCategories {
			axis := fmt.Sprintf("%s%d", SubCategoryCol, i+1)
			f.SetCellValue(s, axis, sc.Name)

			for _, ssc := range sc.SubSubCategories {
				axis := fmt.Sprintf("%s%d", SubSubCategoryCol, i+1)
				f.SetCellValue(s, axis, ssc.Name)

				sb := strings.Builder{}

				axis = fmt.Sprintf("%s%d", ProceduresCol, i+1)
				for j, p := range ssc.Procedures {
					if j != 0 {
						sb.WriteRune('\n')
					}
					sb.WriteString(fmt.Sprintf("%d. ", j+1))
					sb.WriteString(p)
				}
				f.SetCellValue(s, axis, sb.String())

				sb.Reset()

				axis = fmt.Sprintf("%s%d", ConfirmationsCol, i+1)
				for j, p := range ssc.Confirmations {
					if j != 0 {
						sb.WriteRune('\n')
					}
					sb.WriteRune('・')
					sb.WriteString(p)
				}
				f.SetCellValue(s, axis, sb.String())
			}
		}
	}

	if err := f.SaveAs("book.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func exampleSpec() *Spec {
	ssc := SubSubCategory{
		Name:          "SubSubCategory1",
		Procedures:    []string{"Procedure1", "Procedure2", "Procedure3"},
		Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}}
	sc := SubCategory{Name: "SubCategory1", SubSubCategories: []*SubSubCategory{&ssc}}
	c := Category{Name: "Category1", SubCategories: []*SubCategory{&sc}}
	spec := Spec{Name: "Spec1", Categories: []*Category{&c}}
	return &spec
}
