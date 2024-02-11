package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"example.com/minitializr/initializers"
	"example.com/minitializr/types"
	"example.com/minitializr/utils"
	cp "github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
)

var File string
var AsZip bool
var Dest string

func init() {
	generateCmd.Flags().StringVarP(&File, "file", "f", "", "file path")
	generateCmd.MarkFlagRequired("file")
	generateCmd.Flags().BoolVarP(&AsZip, "zip", "z", false, "generate as a zip")
	generateCmd.Flags().StringVarP(&Dest, "dest", "d", ".", "destination folder")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate",
	RunE:  generate,
}

func generate(cmd *cobra.Command, args []string) error {
	miConfig := getMIConfig(File)
	for serviceName, service := range miConfig.Services {
		getInitializer(serviceName, service).Initialize(&miConfig)
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	projectPath := fmt.Sprintf("%s/.minitializer/%s", userHomeDir, miConfig.Metadata["name"])
	if AsZip {
		err = utils.ZipDirectory(Dest + "/" + miConfig.Metadata["name"]+".zip", projectPath)
		if err != nil {
			return err
		}
	} else {
		err := cp.Copy(projectPath, Dest + "/" + miConfig.Metadata["name"])
		if err != nil {
			return err
		}
	}
	return nil
}

func getMIConfig(filePath string) types.MIConfig {
	yamlFile, err := os.ReadFile(filePath)
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
	baseInitializer := initializers.BaseInitializer{
		ServiceName: serviceName,
		Service:     service,
	}
	switch service.Technology {
	case "spring-boot":
		return initializers.SpringBootInitializer(baseInitializer)
	case "micronaut":
		return initializers.MicronautInitializer(baseInitializer)
	case "quarkus":
		return initializers.QuarkusInitializer(baseInitializer)
	case "grails":
		return initializers.GrailsInitializer(baseInitializer)
	case "vertx":
		return initializers.VertxInitializer(baseInitializer)
	}
	return baseInitializer
}
