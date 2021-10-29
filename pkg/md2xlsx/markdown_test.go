package md2xlsx

import (
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func TestParseHeading1(t *testing.T) {
	input := `
# Spec
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if h, ok := n.(*ast.Heading); ok {
		h1 := parseHeading(h)
		if h1 != "Spec" {
			t.Errorf("h1 = %v, want %v", h1, "Spec")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseHeading2(t *testing.T) {
	input := `
## Category
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if h, ok := n.(*ast.Heading); ok {
		h2 := parseHeading(h)
		if h2 != "Category" {
			t.Errorf("h2 = %v, want %v", h2, "Category")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseHeading3(t *testing.T) {
	input := `
### SubCategory
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if h, ok := n.(*ast.Heading); ok {
		h3 := parseHeading(h)
		if h3 != "SubCategory" {
			t.Errorf("h3 = %v, want %v", h3, "SubCategory")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseHeading4(t *testing.T) {
	input := `
#### SubSubCategory
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if h, ok := n.(*ast.Heading); ok {
		h4 := parseHeading(h)
		if h4 != "SubSubCategory" {
			t.Errorf("h4 = %v, want %v", h4, "SubSubCategory")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseUnorderedList(t *testing.T) {
	input := `
* Procedure 1
* Procedure 2
* Procedure 3
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if l, ok := n.(*ast.List); ok {
		p := parseUnorderdList(l)
		if len(p) != 3 {
			t.Errorf("len(p) = %v, want %v", len(p), 3)
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseOrderedList(t *testing.T) {
	input := `
1. Procedure 1
2. Procedure 2
3. Procedure 3
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if l, ok := n.(*ast.List); ok {
		p := parseOrderdList(l)
		if len(p) != 3 {
			t.Errorf("len(p) = %v, want %v", len(p), 3)
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}
