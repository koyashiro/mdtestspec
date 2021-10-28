package main

import (
	"fmt"

	"github.com/koyashiro/md2xlsx/pkg/md2xlsx"
)

func main() {
	s := exampleSpec()

	if err := s.SaveAs("book.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func exampleSpec() *md2xlsx.Spec {
	ssc := md2xlsx.SubSubCategory{
		Name:          "SubSubCategory1",
		Procedures:    []string{"Procedure1", "Procedure2", "Procedure3"},
		Confirmations: []string{"Confirmation1", "Confirmation2", "Confirmation3"}}
	sc := md2xlsx.SubCategory{Name: "SubCategory1", SubSubCategories: []*md2xlsx.SubSubCategory{&ssc}}
	c := md2xlsx.Category{Name: "Category1", SubCategories: []*md2xlsx.SubCategory{&sc}}
	spec := md2xlsx.Spec{Name: "Spec1", Categories: []*md2xlsx.Category{&c}}
	return &spec
}
