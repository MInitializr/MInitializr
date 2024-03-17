package types

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type MIConfig struct {
	ApiVersion string               `json:"apiVersion" binding:"required"`
	Metadata   map[string]string    `json:"metadata" binding:"required"`
	Services   map[string]MIService `json:"services" binding:"required"`
}

func (mi *MIConfig) Validate(supportedFrameworks SupportedFrameworks) error {
	for _, service := range mi.Services {
		err := mi.validateVersion(&service, supportedFrameworks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*MIConfig) validateVersion(service *MIService, supportedFrameworks SupportedFrameworks) error {
	validVersion := false
	versions := supportedFrameworks[service.Technology].Versions
	serviceVersion := fmt.Sprintf("%v", service.Version)
	for alias, version := range versions {
		if version == serviceVersion || alias == serviceVersion {
			service.Version = alias
			validVersion = true
		}
	}
	if !validVersion {
		return fmt.Errorf("validation error: not a valid %s version, actual version %s, supported versions %v", service.Technology, serviceVersion, maps.Values(versions))
	}
	return nil
}
