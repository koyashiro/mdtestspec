package parser

import (
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func TestParse(t *testing.T) {
	input := []byte(`
# Spec

## Category 1

### Sub Category 1-1

#### Sub Sub Category 1-1-1

1. Procedure 1-1-1-1

#### Sub Sub Category 1-1-2

1. Procedure 1-1-2-1
2. Procedure 1-1-2-2

#### Sub Sub Category 1-1-3

1. Procedure 1-1-3-1
2. Procedure 1-1-3-2
3. Procedure 1-1-3-3

### Sub Category 1-2

#### Sub Sub Category 1-2-1

1. Procedure 1-2-1-1

#### Sub Sub Category 1-2-2

1. Procedure 1-2-2-1
2. Procedure 1-2-2-2

## Category 2

### Sub Category 2-1

#### Sub Sub Category 2-1-1

1. Procedure 2-1-1-1

#### Sub Sub Category 2-1-2

1. Procedure 2-1-2-1
2. Procedure 2-1-2-2

#### Sub Sub Category 2-1-3

1. Procedure 2-1-3-1
2. Procedure 2-1-3-2
3. Procedure 2-1-3-3

### Sub Category 2-2

#### Sub Sub Category 2-2-1

1. Procedure 2-2-1-1

### Sub Category 2-3

#### Sub Sub Category 2-3-1

1. Procedure 2-3-1-1

#### Sub Sub Category 2-3-2

1. Procedure 2-3-2-1
2. Procedure 2-3-2-2

#### Sub Sub Category 2-3-3

1. Procedure 2-3-3-1
2. Procedure 2-3-3-2
3. Procedure 2-3-3-3

#### Sub Sub Category 2-3-4

1. Procedure 2-3-4-1
2. Procedure 2-3-4-2
3. Procedure 2-3-4-3
4. Procedure 2-3-4-4
`)

	s, err := ParseSpec(input)
	if err != nil {
		t.Errorf("err = %v, want %v", err, "nil")
	}

	if s.Name != "Spec" {
		t.Errorf("s.Name = %v, want %v", s.Name, "Spec")
	}

	if len(s.Categories) != 2 {
		t.Errorf("len(s.Categories) = %v, want %v", len(s.Categories), 2)
	}

	// Category 1
	c := s.Categories[0]
	if c.Name != "Category 1" {
		t.Errorf("c.Name = %v, want %v", c.Name, "Category 1")
	}
	if len(c.SubCategories) != 2 {
		t.Errorf("len(c.SubCategories) = %v, want %v", len(c.SubCategories), 2)
	}

	// Sub Category 1-1
	sc := c.SubCategories[0]
	if sc.Name != "Sub Category 1-1" {
		t.Errorf("sc.Name = %v, want %v", sc.Name, "Sub Category 1-1")
	}
	if len(sc.SubSubCategories) != 3 {
		t.Errorf("len(sc.SubSubCategories) = %v, want %v", len(sc.SubSubCategories), 3)
	}

	// Sub Sub Category 1-1-1
	ssc := sc.SubSubCategories[0]
	if ssc.Name != "Sub Sub Category 1-1-1" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 1-1-1")
	}
	if len(ssc.Procedures) != 1 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 1)
	}

	// Procedure 1-1-1-1
	p := ssc.Procedures[0]
	if p != "Procedure 1-1-1-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-1-1")
	}

	// Sub Sub Category 1-1-2
	ssc = sc.SubSubCategories[1]
	if ssc.Name != "Sub Sub Category 1-1-2" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 1-1-2")
	}
	if len(ssc.Procedures) != 2 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 2)
	}

	// Procedure 1-1-2-1
	p = ssc.Procedures[0]
	if p != "Procedure 1-1-2-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-2-1")
	}

	// Procedure 1-1-2-2
	p = ssc.Procedures[1]
	if p != "Procedure 1-1-2-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-2-2")
	}

	// Sub Sub Category 1-1-3
	ssc = sc.SubSubCategories[2]
	if ssc.Name != "Sub Sub Category 1-1-3" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 1-1-3")
	}
	if len(ssc.Procedures) != 3 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 3)
	}

	// Procedure 1-1-3-1
	p = ssc.Procedures[0]
	if p != "Procedure 1-1-3-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-3-1")
	}

	// Procedure 1-1-3-2
	p = ssc.Procedures[1]
	if p != "Procedure 1-1-3-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-3-2")
	}

	// Procedure 1-1-3-3
	p = ssc.Procedures[2]
	if p != "Procedure 1-1-3-3" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-1-3-3")
	}

	// Sub Category 1-2
	sc = c.SubCategories[1]
	if sc.Name != "Sub Category 1-2" {
		t.Errorf("sc.Name = %v, want %v", sc.Name, "Sub Category 1-2")
	}
	if len(sc.SubSubCategories) != 2 {
		t.Errorf("len(sc.SubSubCategories) = %v, want %v", len(sc.SubSubCategories), 2)
	}

	// Sub Sub Category 1-2-1
	ssc = sc.SubSubCategories[0]
	if ssc.Name != "Sub Sub Category 1-2-1" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 1-2-2")
	}
	if len(ssc.Procedures) != 1 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 1)
	}

	// Procedure 1-2-1-1
	p = ssc.Procedures[0]
	if p != "Procedure 1-2-1-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-2-1-1")
	}

	// Sub Sub Category 1-2-2
	ssc = sc.SubSubCategories[1]
	if ssc.Name != "Sub Sub Category 1-2-2" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 1-2-2")
	}
	if len(ssc.Procedures) != 2 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 2)
	}

	// Procedure 1-2-2-1
	p = ssc.Procedures[0]
	if p != "Procedure 1-2-2-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 1-2-2-1")
	}

	// Category 2
	c = s.Categories[1]
	if c.Name != "Category 2" {
		t.Errorf("c.Name = %v, want %v", c.Name, "Category 2")
	}
	if len(c.SubCategories) != 3 {
		t.Errorf("len(c.SubCategories) = %v, want %v", len(c.SubCategories), 3)
	}

	// Sub Category 2-1
	sc = c.SubCategories[0]
	if sc.Name != "Sub Category 2-1" {
		t.Errorf("sc.Name = %v, want %v", sc.Name, "Sub Category 2-1")
	}
	if len(sc.SubSubCategories) != 3 {
		t.Errorf("len(sc.SubSubCategories) = %v, want %v", len(sc.SubSubCategories), 3)
	}

	// Sub Sub Category 2-1-1
	ssc = sc.SubSubCategories[0]
	if ssc.Name != "Sub Sub Category 2-1-1" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-1-1")
	}
	if len(ssc.Procedures) != 1 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 1)
	}

	// Procedure 2-1-1-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-1-1-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-1-1")
	}

	// Sub Sub Category 2-1-2
	ssc = sc.SubSubCategories[1]
	if ssc.Name != "Sub Sub Category 2-1-2" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-1-2")
	}
	if len(ssc.Procedures) != 2 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 2)
	}

	// Procedure 2-1-2-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-1-2-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-2-1")
	}

	// Procedure 2-1-2-2
	p = ssc.Procedures[1]
	if p != "Procedure 2-1-2-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-2-2")
	}

	// Sub Sub Category 2-1-3
	ssc = sc.SubSubCategories[2]
	if ssc.Name != "Sub Sub Category 2-1-3" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-1-3")
	}
	if len(ssc.Procedures) != 3 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 3)
	}

	// Procedure 2-1-3-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-1-3-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-3-1")
	}

	// Procedure 2-1-3-2
	p = ssc.Procedures[1]
	if p != "Procedure 2-1-3-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-3-2")
	}

	// Procedure 2-1-3-3
	p = ssc.Procedures[2]
	if p != "Procedure 2-1-3-3" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-1-3-3")
	}

	// Sub Category 2-2
	sc = c.SubCategories[1]
	if sc.Name != "Sub Category 2-2" {
		t.Errorf("sc.Name = %v, want %v", sc.Name, "Sub Category 2-2")
	}
	if len(sc.SubSubCategories) != 1 {
		t.Errorf("len(sc.SubSubCategories) = %v, want %v", len(sc.SubSubCategories), 1)
	}

	// Sub Sub Category 2-2-1
	ssc = sc.SubSubCategories[0]
	if ssc.Name != "Sub Sub Category 2-2-1" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-2-1")
	}
	if len(ssc.Procedures) != 1 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 1)
	}

	// Procedure 2-2-1-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-2-1-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-2-1-1")
	}

	// Sub Category 2-3
	sc = c.SubCategories[2]
	if sc.Name != "Sub Category 2-3" {
		t.Errorf("sc.Name = %v, want %v", sc.Name, "Sub Category 2-3")
	}
	if len(sc.SubSubCategories) != 4 {
		t.Errorf("len(sc.SubSubCategories) = %v, want %v", len(sc.SubSubCategories), 4)
	}

	// Sub Sub Category 2-3-1
	ssc = sc.SubSubCategories[0]
	if ssc.Name != "Sub Sub Category 2-3-1" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-3-1")
	}
	if len(ssc.Procedures) != 1 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 1)
	}

	// Procedure 2-3-1-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-3-1-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-1-1")
	}

	// Sub Sub Category 2-3-2
	ssc = sc.SubSubCategories[1]
	if ssc.Name != "Sub Sub Category 2-3-2" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-3-2")
	}
	if len(ssc.Procedures) != 2 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 2)
	}

	// Procedure 2-3-2-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-3-2-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-2-1")
	}

	// Procedure 2-3-2-2
	p = ssc.Procedures[1]
	if p != "Procedure 2-3-2-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-2-2")
	}

	// Sub Sub Category 2-3-3
	ssc = sc.SubSubCategories[2]
	if ssc.Name != "Sub Sub Category 2-3-3" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-3-3")
	}
	if len(ssc.Procedures) != 3 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 3)
	}

	// Procedure 2-3-3-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-3-3-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-3-1")
	}

	// Procedure 2-3-3-2
	p = ssc.Procedures[1]
	if p != "Procedure 2-3-3-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-3-2")
	}

	// Procedure 2-3-3-3
	p = ssc.Procedures[2]
	if p != "Procedure 2-3-3-3" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-3-3")
	}

	// Sub Sub Category 2-3-4
	ssc = sc.SubSubCategories[3]
	if ssc.Name != "Sub Sub Category 2-3-4" {
		t.Errorf("ssc.Name = %v, want %v", ssc.Name, "Sub Sub Category 2-3-4")
	}
	if len(ssc.Procedures) != 4 {
		t.Errorf("len(ssc.Procedures) = %v, want %v", len(ssc.Procedures), 4)
	}

	// Procedure 2-3-4-1
	p = ssc.Procedures[0]
	if p != "Procedure 2-3-4-1" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-4-1")
	}

	// Procedure 2-3-4-2
	p = ssc.Procedures[1]
	if p != "Procedure 2-3-4-2" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-4-2")
	}

	// Procedure 2-3-4-3
	p = ssc.Procedures[2]
	if p != "Procedure 2-3-4-3" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-4-3")
	}

	// Procedure 2-3-4-4
	p = ssc.Procedures[3]
	if p != "Procedure 2-3-4-4" {
		t.Errorf("p = %v, want %v", p, "Procedure 2-3-4-4")
	}
}

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

func TestParseOrderedList(t *testing.T) {
	input := `
1. Procedure 1
2. Procedure 2
3. Procedure 3
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if l, ok := n.(*ast.List); ok {
		p := parseOrderedList(l)
		if len(p) != 3 {
			t.Errorf("len(p) = %v, want %v", len(p), 3)
		}

		if p[0] != "Procedure 1" {
			t.Errorf("p[0] = %v, want %v", p[0], "Procedure 1")
		}

		if p[1] != "Procedure 2" {
			t.Errorf("p[1] = %v, want %v", p[1], "Procedure 2")
		}

		if p[2] != "Procedure 3" {
			t.Errorf("p[2] = %v, want %v", p[1], "Procedure 3")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}

func TestParseCheckList(t *testing.T) {
	input := `
- [ ] Confirmation 1
- [ ] Confirmation 2
- [ ] Confirmation 3
`
	p := parser.New()
	n := p.Parse([]byte(input)).GetChildren()[0]
	if l, ok := n.(*ast.List); ok {
		c := parseCheckList(l)
		if len(c) != 3 {
			t.Errorf("len(c) = %v, want %v", len(c), 3)
		}

		if c[0] != "Confirmation 1" {
			t.Errorf("c[0] = %v, want %v", c[0], "Confirmation 1")
		}

		if c[1] != "Confirmation 2" {
			t.Errorf("c[1] = %v, want %v", c[1], "Confirmation 2")
		}

		if c[2] != "Confirmation 3" {
			t.Errorf("c[2] = %v, want %v", c[1], "Confirmation 3")
		}
	} else {
		t.Errorf("ok = %v, want %v", ok, true)
	}
}
