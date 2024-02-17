package types

type MIConfig struct {
	ApiVersion string               `json:"apiVersion" binding:"required"`
	Metadata   map[string]string    `json:"metadata" binding:"required"`
	Services   map[string]MIService `json:"services" binding:"required"`
}
