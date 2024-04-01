package datastructs

import "strings"

const MIMIR_CPU_MEM_DISK_TENANT_URL = "https://mimir-push-cpu-mem-disk-metrics.prod.azure.example.net/api/v1/push"
const MIMIR_SERVICE_PROCESS_TENANT_URL = "https://mimir-push-service-process-metrics.prod.azure.example.net/api/v1/push"
const MIMIR_AGENT_TENANT_URL = "https://mimir-push-agent-metrics.prod.azure.example.net/api/v1/push"
const LOKI_PUSH_API_URL = "https://loki-write.prod.azure.example.net/loki/api/v1/push"

type WindowsConfig struct {
	Server       Server       `json:"server"`
	Metrics      Metrics      `json:"metrics"`
	Logs         Logs         `json:"logs"`
	Integrations Integrations `json:"integrations"`
	TemplateData TemplateData
}

type Integrations struct {
	Metrics IntegrationsMetrics `json:"metrics"`
	Windows Windows             `json:"windows"`
	Agent   Agent               `json:"agent"`
}

type Agent struct {
	Autoscrape AgentAutoscrape `json:"autoscrape"`
}

type AgentAutoscrape struct {
	MetricsInstance string          `json:"metrics_instance"`
	RelabelConfigs  []RelabelConfig `json:"relabel_configs"`
}

type RelabelConfig struct {
	SourceLabels []string `json:"source_labels"`
	TargetLabel  string   `json:"target_label"`
	Regex        string   `json:"regex"`
	Replacement  string   `json:"replacement"`
}

type IntegrationsMetrics struct {
	Autoscrape MetricsAutoscrape `json:"autoscrape"`
}

type MetricsAutoscrape struct {
	Enable bool `json:"enable"`
}

type Windows struct {
	Autoscrape        AgentAutoscrape `json:"autoscrape"`
	EnabledCollectors string          `json:"enabled_collectors"`
}

type Logs struct {
	PositionsDirectory string       `json:"positions_directory"`
	Configs            []LogsConfig `json:"configs"`
}

type LogsConfig struct {
	Name          string         `json:"name"`
	Clients       []Client       `json:"clients"`
	ScrapeConfigs []ScrapeConfig `json:"scrape_configs"`
}

type Client struct {
	URL            string         `json:"url"`
	TLSConfig      TLSConfig      `json:"tls_config"`
	ExternalLabels ExternalLabels `json:"external_labels"`
}

type ExternalLabels struct {
	Owner        string `json:"owner"`
	Team         string `json:"team"`
	Dept         string `json:"dept"`
	Env          string `json:"env"`
	App          string `json:"app"`
	Subscription string `json:"subscription"`
	Datacenter   string `json:"datacenter"`
	Silence      string `json:"silence"` //bool
	Los          string `json:"los"`
	SupportTier  string `json:"supporttier"`
	Virtual      string `json:"virtual"` //bool
	Hostname     string `json:"hostname,omitempty"`
}

type TLSConfig struct {
	InsecureSkipVerify bool `json:"insecure_skip_verify"`
}

type ScrapeConfig struct {
	JobName       string        `json:"job_name"`
	WindowsEvents WindowsEvents `json:"windows_events"`
}

type WindowsEvents struct {
	EventlogName         string `json:"eventlog_name"`
	UseIncomingTimestamp bool   `json:"use_incoming_timestamp"`
	BookmarkPath         string `json:"bookmark_path"`
	ExcludeEventData     bool   `json:"exclude_event_data"`
	Locale               int64  `json:"locale"`
	Labels               Labels `json:"labels"`
}

type Labels struct {
}

type Metrics struct {
	WalDirectory string          `json:"wal_directory"`
	Global       Global          `json:"global"`
	Configs      []MetricsConfig `json:"configs"`
}

type MetricsConfig struct {
	Name          string        `json:"name"`
	RemoteWrite   []RemoteWrite `json:"remote_write"`
	ScrapeConfigs []interface{} `json:"scrape_configs"`
}

type RemoteWrite struct {
	URL                 string               `json:"url"`
	WriteRelabelConfigs []WriteRelabelConfig `json:"write_relabel_configs"`
}

type WriteRelabelConfig struct {
	Action       string   `json:"action"`
	SourceLabels []string `json:"source_labels"`
	Separator    string   `json:"separator"`
	Regex        string   `json:"regex"`
}

type Global struct {
	ScrapeInterval string         `json:"scrape_interval"`
	ExternalLabels ExternalLabels `json:"external_labels"`
}

type Server struct {
	LogLevel string `json:"log_level"`
}
type TemplateData struct {
	CPUMemDiskHopURL     string
	ServiceProcessHopURL string
	AgentHopURL          string
	LogsHopURL           string
	CustomLog            bool
}

func GenerateConfig(col []string) *WindowsConfig {
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

	config := WindowsConfig{
		TemplateData: TemplateData{
			CPUMemDiskHopURL:     CPUMemDiskHopURL,
			ServiceProcessHopURL: ServiceProcessHopURL,
			AgentHopURL:          AgentHopURL,
			LogsHopURL:           LogsHopURL,
		},
		Metrics: Metrics{
			Global: Global{
				ExternalLabels: ExternalLabels{
					Owner:        col[8],
					Team:         col[10],
					Dept:         col[7],
					Env:          col[6],
					App:          col[5],
					Subscription: col[1],
					Datacenter:   col[3],
					Silence:      strings.ToLower(col[11]), // bool
					Los:          col[12],
					SupportTier:  col[9],
					Virtual:      strings.ToLower(col[13]), //bool
				},
			},
		},
		Logs: Logs{
			Configs: []LogsConfig{
				{Name: "default",
					Clients: []Client{
						{
							URL: LogsHopURL,
							ExternalLabels: ExternalLabels{
								Owner:        col[8],
								Team:         col[10],
								Dept:         col[7],
								Env:          col[6],
								App:          col[5],
								Subscription: col[1],
								Datacenter:   col[3],
								Silence:      strings.ToLower(col[11]), // bool
								Los:          col[12],
								SupportTier:  col[9],
								Virtual:      strings.ToLower(col[13]), //bool
							},
						},
					},
				},
			},
		},
	}
	return &config

}
