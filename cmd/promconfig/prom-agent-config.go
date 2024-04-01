package main

//datahandler "prom-agent-config/pkg/data_handler"
import (
	"flag"
	datahandler "prom-agent-config/internal/data_handler"
)

func main() {
	var outDir string
	var csvIn string
	initDb := flag.Bool("init", false, "a bool")
	flag.StringVar(&outDir, "outdir", "output_dir", "Directory")
	flag.StringVar(&csvIn, "csv-in", "configs/csv/serverlist-detailed.csv", "Location to CSV File to parse")
	flag.Parse()

	datahandler.OutDir = outDir

	if *initDb {
		csvreader := datahandler.ReadCSVData

		//csvreader("configs/csv/AzMachines.csv")
		csvreader(csvIn)

	} else {
		ppAdd, ppDel := datahandler.ParsePatch("git-patches/1commit.patch")

		datahandler.UpdateConfig(ppAdd, ppDel)

	}

}
