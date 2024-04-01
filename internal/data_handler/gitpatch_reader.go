package datahandler

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

type MarksStruct struct {
	Sceince     string `yaml:"science"`
	Mathematics string `yaml:"mathematics"`
	English     string `yaml:"english"`
}

type Student struct {
	Name   string      `yaml:"student-name"`
	Age    int8        `yaml:"student-age"`
	Marks  MarksStruct `yaml:"subject-marks"`
	Sports []string
}

func ParsePatch(file string) ([]string, []string) {
	patch, err := os.Open(file)
	check(err)

	// files is a slice of *gitdiff.File describing the files changed in the patch
	// preamble is a string of the content of the patch before the first file
	files, _, err := gitdiff.Parse(patch)
	check(err)

	fmt.Println("===================NEXT=============")
	var sadd []string //myslice of additions
	var sdel []string //myslice of deletions

	for _, frag := range files[0].TextFragments { // Struct
		fmt.Printf("text fragment with +%d/-%d lines\n", frag.LinesAdded, frag.LinesDeleted)
		fmt.Println(frag.Lines)

		for _, mylines := range frag.Lines {
			mystring := mylines.String()

			if strings.HasPrefix(mystring, "+") {
				fmt.Println("adding")
				fmt.Println("this line: ", mylines)
				sadd = append(sadd, mystring)

			} else if strings.HasPrefix(mystring, "-") {
				fmt.Println("Deleting")
				fmt.Println("this line: ", mylines)
				sdel = append(sdel, mystring)
			}

		}
	}

	return sadd, sdel
}

func UpdateConfig(sadd []string, sdel []string) {

	fmt.Println("loading additions")

	for _, line := range sadd {
		r := csv.NewReader(strings.NewReader(line))
		record, err := r.Read()
		check(err)

		fmt.Println("adding additions")
		ct := record[4]
		fn := strings.TrimLeft(record[0][1:], "+")

		generateAllConfigs(ct, record, fn)

		fmt.Println("Generated YAML file successfully.")
	}
}
