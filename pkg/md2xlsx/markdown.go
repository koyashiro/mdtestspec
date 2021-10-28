package md2xlsx

import (
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
		switch n.(type) {
		case *ast.Heading:
			s.Name = getHeadingContent(n)
		}
	}
	return &s
}

func ParseMarkdown(input []byte) *Markdown {
	p := parser.New()
	n := p.Parse(input)
	return &Markdown{node: n}
}

func OpenMarkdown(filename string) (*Markdown, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseMarkdown(b), nil
}

func getHeadingContent(heading ast.Node) string {
	for _, n := range heading.GetChildren() {
		if l := n.AsLeaf(); l != nil {
			return string(l.Literal)
		}
	}
	return ""
}
