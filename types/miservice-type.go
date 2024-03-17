package types

type MIService struct {
	Technology string		`json:"technology" binding:"required"`
	Version any			`json:"version" binding:"required"`
	Config	map[string]any	`json:"config" binding:"required"`
}
