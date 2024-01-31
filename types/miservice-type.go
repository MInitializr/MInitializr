package types

type MIService struct {
	Technology string         `yaml:"technology"`
	Config     map[string]any `yaml:"config"`
}
