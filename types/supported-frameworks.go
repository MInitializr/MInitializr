package types

type SupportedFrameworks map[string]SupportedFramework

type SupportedFramework struct {
	Versions map[string]string `yaml:"versions" binding:"required"`
}