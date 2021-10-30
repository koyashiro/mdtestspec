package parser

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"

	"github.com/koyashiro/md2xlsx/pkg/spec"
)

func ParseSpec(input []byte) (*spec.Spec, error) {
	p := parser.New()
	n := p.Parse(input)
	doc, ok := n.(*ast.Document)
	if !ok {
		panic("invalid node")
	}

	s := &spec.Spec{}
	for _, n := range doc.Children {
		if heading, ok := n.(*ast.Heading); ok {
			switch heading.Level {
			case 1:
				if s.Name != "" {
					return nil, errors.New("unexpected h1 element")
				}
				s.Name = parseHeading(heading)
			case 2:
				if s.Name == "" {
					return nil, errors.New("unexpected h2 element")
				}
				c := &spec.Category{Name: parseHeading(heading)}
				s.Categories = append(s.Categories, c)
			case 3:
				if len(s.Categories) == 0 {
					return nil, errors.New("unexpected h3 element")
				}
				c := s.Categories[len(s.Categories)-1]
				sc := &spec.SubCategory{Name: parseHeading(heading)}
				c.SubCategories = append(c.SubCategories, sc)
			case 4:
				if len(s.Categories) == 0 || len(s.Categories[len(s.Categories)-1].SubCategories) == 0 {
					return nil, errors.New("unexpected h4 element")
				}
				c := s.Categories[len(s.Categories)-1]
				sc := c.SubCategories[len(c.SubCategories)-1]
				ssc := &spec.SubSubCategory{Name: parseHeading(heading)}
				sc.SubSubCategories = append(sc.SubSubCategories, ssc)
			}
		}
	}

	return s, nil
}

func OpenSpec(filename string) (*spec.Spec, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseSpec(data)
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
