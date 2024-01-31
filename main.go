package main

import (
	"log"
	"os"
	"example.com/minitializr/initializers"
	"example.com/minitializr/types"
	"gopkg.in/yaml.v3"
)

func main() {
	miConfig := getMIConfig()
	for serviceName, service := range miConfig.Services {
		getInitializer(serviceName, service).Initialize()
	} 
}

func getMIConfig() types.MIConfig {
	yamlFile, err := os.ReadFile("init-example.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var miConfig types.MIConfig
	err = yaml.Unmarshal(yamlFile, &miConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return miConfig
}


func getInitializer(serviceName string, service types.MIService) types.Initializer {
	var initializer types.Initializer;
	baseInitializer := initializers.BaseIntializer {
		ServiceName : serviceName,
		Service : service,
	}
	switch service.Technology {
	case "spring-boot":
		initializer = initializers.SpringBootIntializer(baseInitializer)
	case "micronaut":
		initializer = initializers.MicronautInitializer(baseInitializer)
	case "quarkus":
		initializer = initializers.QuarkusInitializer(baseInitializer)
	}
	return initializer
}