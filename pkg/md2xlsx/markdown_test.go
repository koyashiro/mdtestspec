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
		s := parseHeading1(h)
		if s.Name != "Spec" {
			t.Errorf("s.Name = %v, want %v", s.Name, "Spec")
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
		c := parseHeading2(h)
		if c.Name != "Category" {
			t.Errorf("c.Name = %v, want %v", c.Name, "Category")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseHeading3(t *testing.T) {
	input := "### SubCategory"
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if h, ok := n.(*ast.Heading); ok {
		sc := parseHeading3(h)
		if sc.Name != "SubCategory" {
			t.Errorf("sc.Name = %v, want %v", sc.Name, "SubCategory")
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
		ssc := parseHeading4(h)
		if ssc.Name != "SubSubCategory" {
			t.Errorf("ssc.Name = %v, want %v", ssc.Name, "SubSubCategory")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}
