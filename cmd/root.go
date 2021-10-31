package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/koyashiro/md2xlsx/pkg/excel"
	"github.com/koyashiro/md2xlsx/pkg/parser"
)

var output string

var rootCmd = cobra.Command{
	Use:     "md2xlsx INPUT",
	Example: "md2xlsx spec.md",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]

		var data []byte
		if input == "-" {
			data = make([]byte, 0)
			r := bufio.NewReader(os.Stdin)
			for {
				l, err := r.ReadByte()
				if err == io.EOF {
					break
				} else if err != nil {
					return err
				}
				data = append(data, l)
			}
		} else {
			var err error
			data, err = ioutil.ReadFile(input)
			if err != nil {
				return err
			}
		}

		spec, err := parser.ParseSpec(data)
		if err != nil {
			return err
		}

		book, err := excel.CreateBook(spec)
		if err != nil {
			return err
		}

		if err := book.SaveAs(output); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(err)
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

func init() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "output.xlsx", "output filepath")
}
