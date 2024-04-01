package datahandler

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/dimchansky/utfbom"
)

func ReadCSVData(filelocation string) {
	sf, err := os.Open(filelocation)
	check(err)

	csvFile, enc := utfbom.Skip(sf)
	fmt.Printf("Detected encoding: %s\n", enc)
	fmt.Println("Successfully Opened CSV file")
	defer sf.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	check(err)

	for _, line := range csvLines[1:] { //skip header
		ct := line[4] //config type = linux or windows
		fn := line[0]

		generateAllConfigs(ct, line, fn)

		fmt.Println("Generated YAML file successfully.")

		//fmt.Println(s1.ServerName + " " + s1.Url + " " + s1.Datah)
	}
}
