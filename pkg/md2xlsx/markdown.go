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
	for _, n := range m.node.GetChildren() {
		if heading, ok := n.(*ast.Heading); ok {
			if heading.Level == 1 {
				return parseHeading1(heading)
			}
		}
	}
	return &Spec{}
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

func parseOrderdList(list *ast.List) []string {
	procedures := make([]string, 0)

	if list.ListFlags != 17 {
		panic(fmt.Sprintf("List type is not ordered, list type: %d", list.ListFlags))
	}

	for _, n := range list.Children {
		if li, ok := n.(*ast.ListItem); ok {
			for _, n := range li.Children {
				if p, ok := n.(*ast.Paragraph); ok {
					for _, n := range p.Children {
						if t, ok := n.(*ast.Text); ok {
							if l := string(t.Literal); l != "" {
								procedures = append(procedures, l)
								continue
							}
						}
						continue
					}
				}
				continue
			}
			continue
		}
	}

	return procedures
}

func parseUnorderdList(list *ast.List) []string {
	procedures := make([]string, 0)

	if list.ListFlags != 16 {
		panic(fmt.Sprintf("List type is not unordered, list type: %d", list.ListFlags))
	}

	for _, n := range list.Children {
		if li, ok := n.(*ast.ListItem); ok {
			for _, n := range li.Children {
				if p, ok := n.(*ast.Paragraph); ok {
					for _, n := range p.Children {
						if t, ok := n.(*ast.Text); ok {
							if l := string(t.Literal); l != "" {
								procedures = append(procedures, l)
								continue
							}
						}
						continue
					}
				}
				continue
			}
			continue
		}
	}

	return procedures
}
