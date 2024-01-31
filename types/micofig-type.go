package types

type MIConfig struct {
	Mi       map[string]string    `yaml:"mi"`
	Services map[string]MIService `yaml:"services"`
}
