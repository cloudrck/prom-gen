package datastructs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const MIMIR_CPU_MEM_DISK_TENANT_URL = "https://mimir-push-cpu-mem-disk-metrics.prod.azure.example.net/api/v1/push"
const MIMIR_SERVICE_PROCESS_TENANT_URL = "https://mimir-push-service-process-metrics.prod.azure.example.net/api/v1/push"
const MIMIR_AGENT_TENANT_URL = "https://mimir-push-agent-metrics.prod.azure.example.net/api/v1/push"
const LOKI_PUSH_API_URL = "https://loki-write.prod.azure.example.net/loki/api/v1/push"

type LinuxSysConfig struct {
	CustomArgs           string
	RestartOnUpgrade     bool
	AgentLogLevel        string
	CPUMemDiskHopURL     string
	ServiceProcessHopURL string
	AgentHopURL          string
	LogsHopURL           string
	OwnerLabel           string
	TeamLabel            string
	DeptLabel            string
	EnvLabel             string
	AppLabel             string
	DatacenterLabel      string
	SubLabel             string
	SilenceLabel         string
	LosLabel             string
	SupportTierLabel     string
	VirtualLabel         string
}

func readTemplateFromFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func replacePlaceholders(templateText string, data LinuxSysConfig) string {
	tmpl, err := template.New("textTemplate").Parse(templateText)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return ""
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return ""
	}

	return buf.String()
}

func writeToFile(filepath string, content string) error {
	createF(filepath)
	return os.WriteFile(filepath, []byte(content), 0644)
}

func GenerateSysConfig(iniTmpl string, iniF string, col []string) *LinuxSysConfig {
	// Read the content from the file
	templateFilepath := iniTmpl
	templateText, err := readTemplateFromFile(templateFilepath)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		//return
	}
	// START-SAME AS WINDOWS
	var CPUMemDiskHopURL string
	var ServiceProcessHopURL string
	var AgentHopURL string
	var LogsHopURL string

	if col[2] == "" { //relay server is null
		CPUMemDiskHopURL = MIMIR_CPU_MEM_DISK_TENANT_URL
		ServiceProcessHopURL = MIMIR_SERVICE_PROCESS_TENANT_URL
		AgentHopURL = MIMIR_AGENT_TENANT_URL
		LogsHopURL = LOKI_PUSH_API_URL

	} else {
		CPUMemDiskHopURL = "http://" + col[2] + ":12345/agent/api/v1/metrics/instance/cpu-mem-disk/write"
		ServiceProcessHopURL = "http://" + col[2] + ":12345/agent/api/v1/metrics/instance/service-process/write"
		AgentHopURL = "http://" + col[2] + ":12345/agent/api/v1/metrics/instance/agent/write"
		LogsHopURL = "http://" + col[2] + ":3500/loki/api/v1/push"
	}
	// END-SAME AS WINDOWS

	// Prepare the data you want to set in the template using the struct
	data := LinuxSysConfig{
		//CustomArgs:           "-server.http.address=0.0.0.0:12345 -server.grpc.address=0.0.0.0:12346 -enable-features=integrations-next -config.expand-env=true -disable-reporting",
		//RestartOnUpgrade:     true,
		//AgentLogLevel:        "info",
		CPUMemDiskHopURL:     CPUMemDiskHopURL,
		ServiceProcessHopURL: ServiceProcessHopURL,
		AgentHopURL:          AgentHopURL,
		LogsHopURL:           LogsHopURL,
		OwnerLabel:           col[8],
		TeamLabel:            col[10],
		DeptLabel:            col[7],
		EnvLabel:             col[6],
		AppLabel:             col[5],
		SubLabel:             col[1],
		DatacenterLabel:      col[3],
		SilenceLabel:         strings.ToLower(col[11]), // bool
		LosLabel:             col[12],
		SupportTierLabel:     col[9],
		VirtualLabel:         strings.ToLower(col[13]), // bool
	}

	// Replace placeholders with the struct values
	updatedContent := replacePlaceholders(templateText, data)

	outputFilepath := iniF // Replace with the desired output file path
	err = writeToFile(outputFilepath, updatedContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		//return
	}

	fmt.Println("Linux sysconfig generated")
	return &data
}
func createF(p string) (*os.File, error) { //create folder for output if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
