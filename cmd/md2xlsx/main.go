package main

import (
	"fmt"
	"os"

	"github.com/koyashiro/md2xlsx/pkg/md2xlsx"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("input path required")
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		fmt.Println("output path required")
		os.Exit(1)
	}

	mdFilename := os.Args[1]
	xlsxFilename := os.Args[2]

	md, err := md2xlsx.OpenMarkdown(mdFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := md.AsSpec()

	b := md2xlsx.NewBook()
	b.WriteSpec(s)
	if err := b.SaveAs(xlsxFilename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
