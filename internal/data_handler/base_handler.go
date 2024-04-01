package datahandler

import (
	dsl "prom-agent-config/pkg/data_structs/linux"
	dsw "prom-agent-config/pkg/data_structs/windows"
	"os"
	"path/filepath"
	"strings"
)

var OutDir string

func generateAllConfigs(ct string, record []string, fn string) {
	switch strings.ToLower(ct) {
	case "linux":
		path := filepath.Join("configs", "linux", "01-config.yml")
		tmpl, err := parseYamltmpl(path)
		check(err)

		path1 := filepath.Join("configs", "linux", "grafana-agent-sysconfig")
		path2 := filepath.Join(OutDir, "linux", fn, fn+"-sysconfig")
		sconfig := dsl.GenerateSysConfig(path1, path2, record)

		config := dsl.GenerateConfig(record, *sconfig)

		// Execute the template and write the generated YAML to the file
		path3 := filepath.Join(OutDir, "linux", fn, fn+".yml")
		GeneateYaml(tmpl, path3, *config)

	case "windows":
		path := filepath.Join("configs", "windows", "01-config.yml")
		tmpl, err := parseYamltmpl(path)
		check(err)

		config := dsw.GenerateConfig(record)

		// Execute the template and write the generated YAML to the file
		path2 := filepath.Join(OutDir, "windows", fn+".yml")
		GeneateYaml(tmpl, path2, *config)
	}
}

func createF(p string) (*os.File, error) { //create folder for output if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
