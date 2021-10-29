package md2xlsx

import (
	"fmt"
	"io/ioutil"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type Markdown struct {
	node ast.Node
}

func (m *Markdown) AsSpec() *Spec {
	s := Spec{}
	for _, n := range m.node.GetChildren() {
		if heading, ok := n.(*ast.Heading); ok {
			switch heading.Level {
			case 1:
				s.Name = getHeadingContent(heading)
			case 2:
				panic("not implemented")
			case 3:
				panic("not implemented")
			case 4:
				panic("not implemented")
			case 5:
				panic("not implemented")
			case 6:
				panic("not implemented")
			}
		}
	}
	return &s
}

type MarkdownParser struct {
	parser *parser.Parser
}

func NewMarkdownParser() *MarkdownParser {
	p := parser.New()
	return &MarkdownParser{parser: p}
}

func (p *MarkdownParser) Parse(input []byte) *Markdown {
	n := p.parser.Parse(input)
	return &Markdown{node: n}
}

func OpenMarkdown(filename string) (*Markdown, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	p := NewMarkdownParser()
	md := p.Parse(b)
	return md, nil
}

func getHeadingContent(heading ast.Node) string {
	for _, n := range heading.GetChildren() {
		if l := n.AsLeaf(); l != nil {
			return string(l.Literal)
		}
	}
	return ""
}

func parseHeading1(heading *ast.Heading) *Spec {
	s := &Spec{}

	if heading.Level != 1 {
		panic(fmt.Sprintf("Heading level is not 1, %d", heading.Level))
	}

	for _, n := range heading.Children {
		if t, ok := n.(*ast.Text); ok {
			if l := string(t.Literal); l != "" {
				s.Name = l
			}
			continue
		}

		if h, ok := n.(*ast.Heading); ok && h.Level == 2 {
			c := parseHeading2(h)
			s.Categories = append(s.Categories, c)
			continue
		}
	}

	return s
}

func parseHeading2(heading *ast.Heading) *Category {
	c := &Category{}

	if heading.Level != 2 {
		panic(fmt.Sprintf("Heading level is not 2, heading level: %d", heading.Level))
	}

	for _, n := range heading.Children {
		if t, ok := n.(*ast.Text); ok {
			if l := string(t.Literal); l != "" {
				c.Name = l
			}
			continue
		}

		if h, ok := n.(*ast.Heading); ok && h.Level == 3 {
			sc := parseHeading3(h)
			c.SubCategories = append(c.SubCategories, sc)
			continue
		}
	}

	return c
}

func parseHeading3(heading *ast.Heading) *SubCategory {
	sc := &SubCategory{}

	if heading.Level != 3 {
		panic(fmt.Sprintf("Heading level is not 3, heading level: %d", heading.Level))
	}

	for _, n := range heading.Children {
		if t, ok := n.(*ast.Text); ok {
			if l := string(t.Literal); l != "" {
				sc.Name = l
			}
			continue
		}

		if h, ok := n.(*ast.Heading); ok && h.Level == 3 {
			ssc := parseHeading4(h)
			sc.SubSubCategories = append(sc.SubSubCategories, ssc)
			continue
		}
	}

	return sc
}

func parseHeading4(heading *ast.Heading) *SubSubCategory {
	ssc := &SubSubCategory{}

	if heading.Level != 4 {
		panic(fmt.Sprintf("Heading level is not 4, heading level: %d", heading.Level))
	}

	for _, n := range heading.Children {
		if t, ok := n.(*ast.Text); ok {
			if l := string(t.Literal); l != "" {
				ssc.Name = l
			}
			continue
		}
	}

	return ssc
}
