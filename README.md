# mdtestspec

## Installation

```sh
go install github.com/koyashiro/mdtestspec
```

## Usage

```sh
mdtestspec [OPTIONS] INPUT
```

## Arguments

### `INPUT`

input filepath.

## Options

### `--format`, `-f`

output format.

- `auto` (file extention)
- `xlsx`
- `json`
- `yaml`,`yml`

default: `auto`

### `--output`, `-o`

output filepath.

default: `-` (stdout)

## Getting Started

example spec file is [here](https://github.com/koyashiro/mdtestspec/blob/develop/example/spec.md).

### Convert to xlsx

```sh
mdtestspec -o output.xlsx spec.md
```

![image](https://user-images.githubusercontent.com/6698252/139558046-dc1d36e9-0fdc-40c3-9edd-3b3f2f36e382.png)

### Convert to json

```sh
mdtestspec -o output.json spec.md
```

```json
{"name":"Spec","categories":[{"name":"Category 1","subCategories":[{"name":"Sub-category 1-1","subSubCategories":[{"name":"Sub-sub-category 1-1-1","procedures":["Procedure 1-1-1-1","Procedure 1-1-1-2","Procedure 1-1-1-3"],"confirmations":["Confirmation 1-1-1-1","Confirmation 1-1-1-2","Confirmation 1-1-1-3"],"remarks":["Remarks 1-1-1-1","Remarks 1-1-1-2"]},{"name":"Sub-sub-category 1-1-2","procedures":["Procedure 1-1-2-1","Procedure 1-1-2-2","Procedure 1-1-2-3"],"confirmations":["Confirmation 1-1-2-1","Confirmation 1-1-2-2","Confirmation 1-1-2-3"],"remarks":[]}]},{"name":"Sub-category 1-2","subSubCategories":[{"name":"Sub-sub-category 1-2-1","procedures":["Procedure 1-2-1-1","Procedure 1-2-1-2","Procedure 1-2-1-3"],"confirmations":["Confirmation 1-2-1-1","Confirmation 1-2-1-2","Confirmation 1-2-1-3"],"remarks":[]},{"name":"Sub-sub-category 1-2-2","procedures":["Procedure 1-2-2-1","Procedure 1-2-2-2","Procedure 1-2-2-3"],"confirmations":["Confirmation 1-2-2-1","Confirmation 1-2-2-2","Confirmation 1-2-2-3"],"remarks":[]}]}]},{"name":"Category 2","subCategories":[{"name":"Sub-category 2-1","subSubCategories":[{"name":"Sub-sub-category 2-1-1","procedures":["Procedure 2-1-1-1","Procedure 2-1-1-2","Procedure 2-1-1-3"],"confirmations":["Confirmation 2-1-1-1","Confirmation 2-1-1-2","Confirmation 2-1-1-3"],"remarks":[]},{"name":"Sub-sub-category 2-1-2","procedures":["Procedure 2-1-2-1","Procedure 2-1-2-2","Procedure 2-1-2-3"],"confirmations":["Confirmation 2-1-2-1","Confirmation 2-1-2-2","Confirmation 2-1-2-3"],"remarks":["Remarks 2-1-2-1"]}]},{"name":"Sub-category 2-2","subSubCategories":[{"name":"Sub-sub-category 2-2-1","procedures":["Procedure 2-2-1-1","Procedure 2-2-1-2","Procedure 2-2-1-3"],"confirmations":["Confirmation 2-2-1-1","Confirmation 2-2-1-2","Confirmation 2-2-1-3"],"remarks":[]},{"name":"Sub-sub-category 2-2-2","procedures":["Procedure 2-2-2-1","Procedure 2-2-2-2","Procedure 2-2-2-3"],"confirmations":["Confirmation 2-2-2-1","Confirmation 2-2-2-2","Confirmation 2-2-2-3"],"remarks":[]}]}]}]}
```

### Convert to yaml

```sh
mdtestspec -o output.yml spec.md
```

```yml
name: Spec
categories:
  - name: Category 1
    sub_categories:
      - name: Sub-category 1-1
        subsubcategories:
          - name: Sub-sub-category 1-1-1
            procedures:
              - Procedure 1-1-1-1
              - Procedure 1-1-1-2
              - Procedure 1-1-1-3
            confirmations:
              - Confirmation 1-1-1-1
              - Confirmation 1-1-1-2
              - Confirmation 1-1-1-3
            remarks:
              - Remarks 1-1-1-1
              - Remarks 1-1-1-2
          - name: Sub-sub-category 1-1-2
            procedures:
              - Procedure 1-1-2-1
              - Procedure 1-1-2-2
              - Procedure 1-1-2-3
            confirmations:
              - Confirmation 1-1-2-1
              - Confirmation 1-1-2-2
              - Confirmation 1-1-2-3
            remarks: []
      - name: Sub-category 1-2
        subsubcategories:
          - name: Sub-sub-category 1-2-1
            procedures:
              - Procedure 1-2-1-1
              - Procedure 1-2-1-2
              - Procedure 1-2-1-3
            confirmations:
              - Confirmation 1-2-1-1
              - Confirmation 1-2-1-2
              - Confirmation 1-2-1-3
            remarks: []
          - name: Sub-sub-category 1-2-2
            procedures:
              - Procedure 1-2-2-1
              - Procedure 1-2-2-2
              - Procedure 1-2-2-3
            confirmations:
              - Confirmation 1-2-2-1
              - Confirmation 1-2-2-2
              - Confirmation 1-2-2-3
            remarks: []
  - name: Category 2
    sub_categories:
      - name: Sub-category 2-1
        subsubcategories:
          - name: Sub-sub-category 2-1-1
            procedures:
              - Procedure 2-1-1-1
              - Procedure 2-1-1-2
              - Procedure 2-1-1-3
            confirmations:
              - Confirmation 2-1-1-1
              - Confirmation 2-1-1-2
              - Confirmation 2-1-1-3
            remarks: []
          - name: Sub-sub-category 2-1-2
            procedures:
              - Procedure 2-1-2-1
              - Procedure 2-1-2-2
              - Procedure 2-1-2-3
            confirmations:
              - Confirmation 2-1-2-1
              - Confirmation 2-1-2-2
              - Confirmation 2-1-2-3
            remarks:
              - Remarks 2-1-2-1
      - name: Sub-category 2-2
        subsubcategories:
          - name: Sub-sub-category 2-2-1
            procedures:
              - Procedure 2-2-1-1
              - Procedure 2-2-1-2
              - Procedure 2-2-1-3
            confirmations:
              - Confirmation 2-2-1-1
              - Confirmation 2-2-1-2
              - Confirmation 2-2-1-3
            remarks: []
          - name: Sub-sub-category 2-2-2
            procedures:
              - Procedure 2-2-2-1
              - Procedure 2-2-2-2
              - Procedure 2-2-2-3
            confirmations:
              - Confirmation 2-2-2-1
              - Confirmation 2-2-2-2
              - Confirmation 2-2-2-3
            remarks: []
```
