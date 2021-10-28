package md2xlsx_test

import (
	"testing"

	"github.com/koyashiro/md2xlsx/pkg/md2xlsx"
)

var p = md2xlsx.NewMarkdownParser()

func TestParseHeading(t *testing.T) {
	input := "# Heading"
	md := p.Parse([]byte(input))
	s := md.AsSpec()
	if s.Name != "Heading" {
		t.Errorf("s.Name = %v, want %v", s.Name, "Heading")
	}
}
