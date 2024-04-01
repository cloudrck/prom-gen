package datastructs

type LinuxConfig struct {
	Server       Server       `yaml:"server"`
	Metrics      Metrics      `yaml:"metrics"`
	Logs         Logs         `yaml:"logs"`
	Integrations Integrations `yaml:"integrations"`
	TemplateData TemplateData
}

type Integrations struct {
	Metrics      IntegrationsMetrics `yaml:"metrics"`
	NodeExporter NodeExporter        `yaml:"node_exporter"`
	Agent        Agent               `yaml:"agent"`
}

type Agent struct {
	Autoscrape AgentAutoscrape `yaml:"autoscrape"`
}

type AgentAutoscrape struct {
	MetricsInstance string          `yaml:"metrics_instance"`
	RelabelConfigs  []RelabelConfig `yaml:"relabel_configs"`
}

type RelabelConfig struct {
	SourceLabels []string `yaml:"source_labels"`
	TargetLabel  string   `yaml:"target_label"`
	Regex        string   `yaml:"regex"`
	Replacement  string   `yaml:"replacement"`
}

type IntegrationsMetrics struct {
	Autoscrape MetricsAutoscrape `yaml:"autoscrape"`
}

type MetricsAutoscrape struct {
	Enable bool `yaml:"enable"`
}

type NodeExporter struct {
	Autoscrape AgentAutoscrape `yaml:"autoscrape"`
	ProcfsPath string          `yaml:"procfs_path"`
	SysfsPath  string          `yaml:"sysfs_path"`
}

type Logs struct {
	PositionsDirectory string       `yaml:"positions_directory"`
	Configs            []LogsConfig `yaml:"configs"`
}

type LogsConfig struct {
	Name          string         `yaml:"name"`
	Clients       []Client       `yaml:"clients"`
	Positions     Positions      `yaml:"positions"`
	ScrapeConfigs []ScrapeConfig `yaml:"scrape_configs"`
}

type Client struct {
	URL            string         `yaml:"url"`
	TLSConfig      TLSConfig      `yaml:"tls_config"`
	ExternalLabels ExternalLabels `yaml:"external_labels"`
}

type ExternalLabels struct {
	Owner      string  `yaml:"owner"`
	Team       string  `yaml:"team"`
	Dept       string  `yaml:"dept"`
	Env        string  `yaml:"env"`
	App        string  `yaml:"app"`
	Datacenter string  `yaml:"datacenter"`
	Hostname   *string `yaml:"hostname,omitempty"`
}

type TLSConfig struct {
	InsecureSkipVerify bool `yaml:"insecure_skip_verify"`
}

type Positions struct {
	Filename string `yaml:"filename"`
}

type ScrapeConfig struct {
	JobName       string         `yaml:"job_name"`
	StaticConfigs []StaticConfig `yaml:"static_configs"`
}

type StaticConfig struct {
	Targets []string `yaml:"targets"`
	Labels  Labels   `yaml:"labels"`
}

type Labels struct {
	Job  string `yaml:"job"`
	Path string `yaml:"__path__"`
}

type Metrics struct {
	WalDirectory string          `yaml:"wal_directory"`
	Global       Global          `yaml:"global"`
	Configs      []MetricsConfig `yaml:"configs"`
}

type MetricsConfig struct {
	Name          string        `yaml:"name"`
	RemoteWrite   []RemoteWrite `yaml:"remote_write"`
	ScrapeConfigs []interface{} `yaml:"scrape_configs"`
}

type RemoteWrite struct {
	URL                 string               `yaml:"url"`
	WriteRelabelConfigs []WriteRelabelConfig `yaml:"write_relabel_configs"`
}

type WriteRelabelConfig struct {
	Action       string   `yaml:"action"`
	SourceLabels []string `yaml:"source_labels"`
	Separator    string   `yaml:"separator"`
	Regex        string   `yaml:"regex"`
}

type Global struct {
	ScrapeInterval string         `yaml:"scrape_interval"`
	RemoteWrite    []RemoteWrite  `yaml:"remote_write"`
	ExternalLabels ExternalLabels `yaml:"external_labels"`
}

type Server struct {
	LogLevel string `yaml:"log_level"`
}

type TemplateData struct {
	CPUMemDiskHopURL     string
	ServiceProcessHopURL string
	AgentHopURL          string
	CustomLog            bool
}

func GenerateConfig(col []string, sconfig LinuxSysConfig) *LinuxConfig {
	var is_customlog bool
	if len(col[15]) > 0 {
		is_customlog = true
	} else {
		is_customlog = false
	}

	config := LinuxConfig{
		Metrics: Metrics{
			Global: Global{
				RemoteWrite: []RemoteWrite{
					{URL: col[1]},
				},
			},
		},
		Logs: Logs{
			Configs: []LogsConfig{
				{Name: "default",
					Clients: []Client{
						{
							URL: col[2],
						},
					},
					ScrapeConfigs: []ScrapeConfig{
						{JobName: "jobname goes here",
							StaticConfigs: []StaticConfig{
								{Labels: Labels{
									Path: "mypath",
								}},
							}},
					},
				},
			},
		},
		TemplateData: TemplateData{
			CPUMemDiskHopURL:     sconfig.CPUMemDiskHopURL,
			ServiceProcessHopURL: sconfig.ServiceProcessHopURL,
			AgentHopURL:          sconfig.AgentHopURL,
			CustomLog:            is_customlog,
		},
	}

	return &config

}
