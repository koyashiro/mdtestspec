package md2xlsx

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
