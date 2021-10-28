package main

import (
	"fmt"

	"github.com/koyashiro/md2xlsx/pkg/md2xlsx"
)

func main() {
	s := exampleSpec()

	b := md2xlsx.NewBook()
	b.WriteSpec(s)
	if err := b.SaveAs("book.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func exampleSpec() *md2xlsx.Spec {
	spec := md2xlsx.Spec{
		Name: "Spec1",
		Categories: []*md2xlsx.Category{
			{Name: "Category1", SubCategories: []*md2xlsx.SubCategory{
				{Name: "SubCategory1", SubSubCategories: []*md2xlsx.SubSubCategory{
					{Name: "SubSubCategory1", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory2", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory3", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory4", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
				}},
				{Name: "SubCategory2", SubSubCategories: []*md2xlsx.SubSubCategory{
					{Name: "SubSubCategory1", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory2", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory3", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
				}},
			}},
			{Name: "Category2", SubCategories: []*md2xlsx.SubCategory{
				{Name: "SubCategory1", SubSubCategories: []*md2xlsx.SubSubCategory{
					{Name: "SubSubCategory1", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory2", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory3", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory4", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory5", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
				}},
				{Name: "SubCategory2", SubSubCategories: []*md2xlsx.SubSubCategory{
					{Name: "SubSubCategory1", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
					{Name: "SubSubCategory2", Procedures: []string{"Procedure1", "Procedure2", "Procedure3"}, Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}},
				}},
			}},
		},
	}

	return &spec
}
