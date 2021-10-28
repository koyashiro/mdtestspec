package md2xlsx_test

import (
	"testing"

	"github.com/koyashiro/md2xlsx/pkg/md2xlsx"
)

func TestParseHeading(t *testing.T) {
	input := "# Heading"
	md := md2xlsx.ParseMarkdown([]byte(input))
	s := md.AsSpec()
	if s.Name != "Heading" {
		t.Errorf("s.Name = %v, want %v", s.Name, "Heading")
	}
}
