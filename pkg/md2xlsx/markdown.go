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
