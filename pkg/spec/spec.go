package spec

type Spec struct {
	Name       string      `json:"name" yaml:"name"`
	Categories []*Category `json:"categories" yaml:"categories"`
}

type Category struct {
	Name          string         `json:"name" yaml:"name"`
	SubCategories []*SubCategory `json:"subCategories" yaml:"sub_categories"`
}

type SubCategory struct {
	Name             string            `json:"name" yaml:"name"`
	SubSubCategories []*SubSubCategory `json:"subSubCategories" yaml:"sub_sub_categories"`
}

type SubSubCategory struct {
	Name          string   `json:"name" yaml:"name"`
	Procedures    []string `json:"procedures" yaml:"procedures"`
	Confirmations []string `json:"confirmations" yaml:"confirmations"`
	Remarks       []string `json:"remarks" yaml:"remarks"`
}
