package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/koyashiro/md2xlsx/pkg/excel"
	"github.com/koyashiro/md2xlsx/pkg/parser"
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

	input := os.Args[1]
	output := os.Args[2]

	data, err := readData(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := parser.ParseSpec(data)

	b := excel.NewBook()
	b.WriteSpec(s)
	if err := b.SaveAs(output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readData(input string) ([]byte, error) {
	var data []byte
	if input == "-" {
		data = make([]byte, 0)
		r := bufio.NewReader(os.Stdin)
		for {
			l, err := r.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			data = append(data, l)
		}
	} else {
		var err error
		data, err = ioutil.ReadFile(input)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
