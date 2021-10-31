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
var categoryHeader string
var subCategoryHeader string
var subSubCategoryHeader string
var procedureHeader string
var confirmationHeader string
var remarksHeader string
var resultHeader string

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
			if output == "-" {
				return errors.New("format or output required")
			}
			format = filepath.Ext(output)[1:]
		}

		config := excel.DefaultConfig
		if categoryHeader != "" {
			config.Header.Category = categoryHeader
		}
		if subCategoryHeader != "" {
			config.Header.SubCategory = subCategoryHeader
		}
		if subSubCategoryHeader != "" {
			config.Header.SubSubCategory = subSubCategoryHeader
		}
		if procedureHeader != "" {
			config.Header.Procedure = procedureHeader
		}
		if confirmationHeader != "" {
			config.Header.Confirmation = confirmationHeader
		}
		if remarksHeader != "" {
			config.Header.Remarks = remarksHeader
		}
		if resultHeader != "" {
			config.Header.Result = resultHeader
		}

		var result []byte
		switch format {
		case "xlsx":
			book, err := excel.CreateBook(spec, &config)
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
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "-", "output filepath")
	rootCmd.Flags().StringVarP(&format, "format", "f", "auto", "output format")
	rootCmd.Flags().StringVarP(&categoryHeader, "header-category", "", "Category", "Category header")
	rootCmd.Flags().StringVarP(&subCategoryHeader, "header-subcategory", "", "Sub-category", "Sub-category header")
	rootCmd.Flags().StringVarP(&subSubCategoryHeader, "header-subsubcategory", "", "Sub-sub-category", "Sub-sub-category header")
	rootCmd.Flags().StringVarP(&procedureHeader, "header-procedure", "", "Procedure", "Procedure header")
	rootCmd.Flags().StringVarP(&confirmationHeader, "header-confirmation", "", "Confirmation", "Confirmation header")
	rootCmd.Flags().StringVarP(&remarksHeader, "header-remarks", "", "Remarks", "Remarks header")
	rootCmd.Flags().StringVarP(&resultHeader, "header-result", "", "Result", "Result header")
}
