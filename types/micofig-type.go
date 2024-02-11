package types

type MIConfig struct {
	ApiVersion string               `yaml:"apiVersion"`
	Metadata   map[string]string    `yaml:"metadata"`
	Services   map[string]MIService `yaml:"services"`
}
