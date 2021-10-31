package excel

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"

	"github.com/koyashiro/mdtestspec/pkg/spec"
)

const CategoryCol = "A"
const SubCategoryCol = "B"
const SubSubCategoryCol = "C"
const ProceduresCol = "D"
const ConfirmationsCol = "E"
const ConfirmationPrefix = '・'
const RemarksCol = "F"
const RemarksPrefix = '・'
const ResultCol = "G"

const CategoryWidth = 20.
const SubCategoryWidth = 25.
const SubSubCategoryWidth = 25.
const ProcedureWidth = 55.
const ConfirmationWidth = 55.
const RemarkWidth = 55.
const ResultWidth = 10.

type Book struct {
	file *excelize.File
}

func (b *Book) SaveAs(name string) error {
	return b.file.SaveAs(name)
}

func (b *Book) WriteTo(w io.Writer) (int64, error) {
	return b.file.WriteTo(w)
}

func (b *Book) WriteToBuffer() (*bytes.Buffer, error) {
	return b.file.WriteToBuffer()
}

type Header struct {
	Category       string
	SubCategory    string
	SubSubCategory string
	Procedure      string
	Confirmation   string
	Remarks        string
	Result         string
}

type Config struct {
	Header Header
}

var DefaultConfig = Config{
	Header: Header{
		Category:       "Category",
		SubCategory:    "Sub-category",
		SubSubCategory: "Sub-sub-category",
		Procedure:      "Procedure",
		Confirmation:   "Confirmation",
		Remarks:        "Remarks",
		Result:         "Result",
	},
}

func CreateBook(spec *spec.Spec, config *Config) (*Book, error) {
	file := excelize.NewFile()

	var sheet string
	if spec.Name == "" {
		sheet = "no title"
	} else {
		sheet = spec.Name
	}

	if sheet != "sheet1" {
		file.NewSheet(sheet)
		file.DeleteSheet("sheet1")
	}

	if err := setCelsWidth(file, sheet); err != nil {
		return nil, err
	}

	if config == nil {
		config = &DefaultConfig
	}
	if err := setHeaders(file, sheet, &config.Header); err != nil {
		return nil, err
	}

	sb := strings.Builder{}
	row := 2

	for i, c := range spec.Categories {
		if i != 0 {
			row++
		}

		if err := setCategory(file, sheet, row, c.Name); err != nil {
			return nil, err
		}

		categoryFrom := row
		for j, sc := range c.SubCategories {
			if j != 0 {
				row++
			}

			if err := setSubCategory(file, sheet, row, sc.Name); err != nil {
				return nil, err
			}

			subCategoryFrom := row
			for k, ssc := range sc.SubSubCategories {
				if k != 0 {
					row++
				}

				if err := setSubSubCategory(file, sheet, row, ssc.Name); err != nil {
					return nil, err
				}
				if err := setConfirmations(file, sheet, row, ssc.Confirmations, &sb); err != nil {
					return nil, err
				}
				if err := setProcedures(file, sheet, row, ssc.Procedures, &sb); err != nil {
					return nil, err
				}
				if err := setRemarks(file, sheet, row, ssc.Remarks, &sb); err != nil {
					return nil, err
				}
			}
			subCategoryTo := row

			hcell := fmt.Sprintf("%s%d", SubCategoryCol, subCategoryFrom)
			vcell := fmt.Sprintf("%s%d", SubCategoryCol, subCategoryTo)
			if err := file.MergeCell(sheet, hcell, vcell); err != nil {
				return nil, err
			}
		}
		categoryTo := row

		hcell := fmt.Sprintf("%s%d", CategoryCol, categoryFrom)
		vcell := fmt.Sprintf("%s%d", CategoryCol, categoryTo)
		if err := file.MergeCell(sheet, hcell, vcell); err != nil {
			return nil, err
		}
	}

	if err := setCellStyle(file, sheet, row); err != nil {
		return nil, err
	}

	return &Book{file: file}, nil
}

func setCelsWidth(f *excelize.File, sheet string) error {
	if err := f.SetColWidth(sheet, CategoryCol, CategoryCol, CategoryWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, SubCategoryCol, SubCategoryCol, SubCategoryWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, SubSubCategoryCol, SubSubCategoryCol, SubSubCategoryWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, ProceduresCol, ProceduresCol, ProcedureWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, ConfirmationsCol, ConfirmationsCol, ConfirmationWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, RemarksCol, RemarksCol, RemarkWidth); err != nil {
		return err
	}
	if err := f.SetColWidth(sheet, ResultCol, ResultCol, ResultWidth); err != nil {
		return err
	}
	return nil
}

func setHeaders(f *excelize.File, sheet string, header *Header) error {
	const headerRow = 1

	categoryColAxis := fmt.Sprintf("%s%d", CategoryCol, headerRow)
	if err := f.SetCellStr(sheet, categoryColAxis, header.Category); err != nil {
		return err
	}

	subCategoryColAxis := fmt.Sprintf("%s%d", SubCategoryCol, headerRow)
	if err := f.SetCellStr(sheet, subCategoryColAxis, header.SubCategory); err != nil {
		return err
	}

	subSubCategoryColAxis := fmt.Sprintf("%s%d", SubSubCategoryCol, headerRow)
	if err := f.SetCellStr(sheet, subSubCategoryColAxis, header.SubSubCategory); err != nil {
		return err
	}

	proceduresColAxis := fmt.Sprintf("%s%d", ProceduresCol, headerRow)
	if err := f.SetCellStr(sheet, proceduresColAxis, header.Procedure); err != nil {
		return err
	}

	confirmationsColAxis := fmt.Sprintf("%s%d", ConfirmationsCol, headerRow)
	if err := f.SetCellStr(sheet, confirmationsColAxis, header.Confirmation); err != nil {
		return err
	}

	remarksColAxis := fmt.Sprintf("%s%d", RemarksCol, headerRow)
	if err := f.SetCellStr(sheet, remarksColAxis, header.Remarks); err != nil {
		return err
	}

	resultColAxis := fmt.Sprintf("%s%d", ResultCol, headerRow)
	if err := f.SetCellStr(sheet, resultColAxis, header.Result); err != nil {
		return err
	}

	styleID, err := f.NewStyle(`{
		"border": [
			{
				"type": "top",
				"color": "#5B9BD5",
				"style": 2
			},
			{
				"type": "right",
				"color": "#5B9BD5",
				"style": 2
			},
			{
				"type": "bottom",
				"color": "#5B9BD5",
				"style": 2
			},
		    {
				"type": "left",
				"color": "#5B9BD5",
				"style": 2
			}
		],
		"fill": { "type": "pattern", "pattern": 1, "color": ["#5B9BD5", "#5B9BD5"] },
		"font": { "color": "#FFFFFF" },
		"alignment": { "horizontal": "center", "vertical": "center" }
	}`)
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, categoryColAxis, resultColAxis, styleID)
}

func setCellStyle(f *excelize.File, sheet string, rowTo int) error {
	const rowFrom = 2
	styleID, err := f.NewStyle(`{
		"border": [
			{
				"type": "top",
				"color": "#5B9BD5",
				"style": 2
			},
			{
				"type": "right",
				"color": "#5B9BD5",
				"style": 2
			},
			{
				"type": "bottom",
				"color": "#5B9BD5",
				"style": 2
			},
		    {
				"type": "left",
				"color": "#5B9BD5",
				"style": 2
			}
		],
		"alignment": { "horizontal": "left", "vertical": "top", "wrap_text": true }
	}`)
	if err != nil {
		return err
	}
	hcell := fmt.Sprintf("%s%d", CategoryCol, rowFrom)
	vcell := fmt.Sprintf("%s%d", ResultCol, rowTo)
	return f.SetCellStyle(sheet, hcell, vcell, styleID)
}

func setCategory(f *excelize.File, sheet string, row int, name string) error {
	axis := fmt.Sprintf("%s%d", CategoryCol, row)
	return f.SetCellStr(sheet, axis, name)
}

func setSubCategory(f *excelize.File, sheet string, row int, name string) error {
	axis := fmt.Sprintf("%s%d", SubCategoryCol, row)
	return f.SetCellStr(sheet, axis, name)
}

func setSubSubCategory(f *excelize.File, sheet string, row int, name string) error {
	axis := fmt.Sprintf("%s%d", SubSubCategoryCol, row)
	return f.SetCellStr(sheet, axis, name)
}

func setProcedures(f *excelize.File, sheet string, row int, procedures []string, sb *strings.Builder) error {
	sb.Reset()
	axis := fmt.Sprintf("%s%d", ProceduresCol, row)
	for j, p := range procedures {
		if j != 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(fmt.Sprintf("%d. ", j+1))
		sb.WriteString(p)
	}
	return f.SetCellStr(sheet, axis, sb.String())
}

func setConfirmations(f *excelize.File, sheet string, row int, confirmations []string, sb *strings.Builder) error {
	sb.Reset()
	axis := fmt.Sprintf("%s%d", ConfirmationsCol, row)
	for j, p := range confirmations {
		if j != 0 {
			sb.WriteRune('\n')
		}
		sb.WriteRune(ConfirmationPrefix)
		sb.WriteString(p)
	}
	return f.SetCellStr(sheet, axis, sb.String())
}

func setRemarks(f *excelize.File, sheet string, row int, remarks []string, sb *strings.Builder) error {
	sb.Reset()
	axis := fmt.Sprintf("%s%d", RemarksCol, row)
	for j, p := range remarks {
		if j != 0 {
			sb.WriteRune('\n')
		}
		sb.WriteRune(RemarksPrefix)
		sb.WriteString(p)
	}
	return f.SetCellStr(sheet, axis, sb.String())
}
