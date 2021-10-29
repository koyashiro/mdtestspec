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
	s := &Spec{}
	for _, n := range m.node.GetChildren() {
		if heading, ok := n.(*ast.Heading); ok {
			switch heading.Level {
			case 1:
				s.Name = parseHeading(heading)
				break
			case 2:
				c := &Category{}
				c.Name = parseHeading(heading)
				s.Categories = append(s.Categories, c)
				break
			case 3:
				s.Name = parseHeading(heading)
				break
			case 4:
				s.Name = parseHeading(heading)
				break
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

func parseHeading(heading *ast.Heading) string {
	for _, n := range heading.Children {
		if t, ok := n.(*ast.Text); ok {
			if l := string(t.Literal); l != "" {
				return l
			}
		}
	}

	return ""
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
