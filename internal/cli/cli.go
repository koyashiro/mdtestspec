package cli

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/koyashiro/mdtestspec/pkg/excel"
	"github.com/koyashiro/mdtestspec/pkg/parser"
)

var output string
var format string

var rootCmd = cobra.Command{
	Use:     "mdtestspec INPUT",
	Example: "mdtestspec spec.md",
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
			data, err = os.ReadFile(input)
			if err != nil {
				return err
			}
		}

		spec, err := parser.ParseSpec(data)
		if err != nil {
			return err
		}

		if format == "auto" {
			format = filepath.Ext(output)[1:]
		}

		var result []byte
		switch format {
		case "xlsx":
			book, err := excel.CreateBook(spec)
			if err != nil {
				return err
			}

			buf, err := book.WriteToBuffer()
			if err != nil {
				return err
			}

			result = buf.Bytes()
		case "json":
			var err error
			result, err = json.Marshal(spec)
			if err != nil {
				return err
			}
			result = append(result, '\n')
		case "yaml", "yml":
			var err error
			result, err = yaml.Marshal(spec)
			if err != nil {
				return err
			}
		default:
			return errors.New("invalid format")
		}

		if output == "-" {
			if _, err := os.Stdout.Write(result); err != nil {
				return err
			}
		} else {
			if err := os.WriteFile(output, result, 0644); err != nil {
				return err
			}
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

func init() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "-", "output filepath")
	rootCmd.Flags().StringVarP(&format, "format", "f", "auto", "output format")
}
