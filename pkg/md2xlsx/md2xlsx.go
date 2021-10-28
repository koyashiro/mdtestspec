package md2xlsx

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

const CategoryCol = "A"
const SubCategoryCol = "B"
const SubSubCategoryCol = "C"
const ProceduresCol = "D"
const ConfirmationsCol = "E"
const ConfirmationPrefix = '・'

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

type Book struct {
	file *excelize.File
}

func NewBook() *Book {
	f := excelize.NewFile()
	return &Book{file: f}
}

func OpenBook(filename string) (*Book, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return &Book{file: f}, nil
}

func (b *Book) SaveAs(name string) error {
	return b.file.SaveAs(name)
}

func (b *Book) WriteSpec(spec *Spec) {
	sb := strings.Builder{}
	sheet := "Sheet1"
	i := 0

	for _, c := range spec.Categories {
		for _, sc := range c.SubCategories {
			for _, ssc := range sc.SubSubCategories {
				i++
				setCategory(b.file, sheet, i, c.Name)
				setSubCategory(b.file, sheet, i, sc.Name)
				setSubSubCategory(b.file, sheet, i, ssc.Name)
				setConfirmations(b.file, sheet, i, ssc.Confirmations, &sb)
				setProcedures(b.file, sheet, i, ssc.Procedures, &sb)
			}
		}
	}
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
		sb.WriteRune(ConfirmationPrefix)
		sb.WriteString(p)
	}
	f.SetCellValue(sheet, axis, sb.String())
}
