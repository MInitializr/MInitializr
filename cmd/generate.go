package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"example.com/minitializr/initializers"
	"example.com/minitializr/types"
	"example.com/minitializr/utils"
	cp "github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
)

// command flags
var (
	FromFile bool
	AsZip    bool
	Dest     string
)

func init() {
	generateCmd.Flags().BoolVarP(&FromFile, "file", "f", true, "generate from file")
	generateCmd.Flags().BoolVarP(&AsZip, "zip", "z", false, "generate as a zip")
	generateCmd.Flags().StringVarP(&Dest, "dest", "d", ".", "destination folder")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate",
	RunE:  generate,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		// Run the custom validation logic
		_, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}


var (
	userHomeDir string
	miDataDir   string
	projectPath string
)

func generate(cmd *cobra.Command, args []string) error {
	miConfig, err := getMIConfig(args[0])
	if err != nil {
		return err
	}

	initVariables(miConfig)

	for serviceName, service := range miConfig.Services {
		getInitializer(serviceName, service).Initialize(miConfig)
	}

	if FromFile {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		miDataDir = fmt.Sprintf("%s/.minitializer", userHomeDir)
		projectPath = fmt.Sprintf("%s/%s", miDataDir, miConfig.Metadata["name"])
		if AsZip {
			err = utils.ZipDirectory(Dest+"/"+miConfig.Metadata["name"]+".zip", projectPath)
			if err != nil {
				return err
			}
		} else {
			err := cp.Copy(projectPath, Dest+"/"+miConfig.Metadata["name"])
			if err != nil {
				return err
			}
		}
	}

	err = cleanUp(miConfig)
	if err != nil {
		return err
	}
	return nil
}

func cleanUp(miConfig *types.MIConfig) error {
	err := os.RemoveAll(projectPath)
	if err != nil {
		return err
	}
	return nil
}

func initVariables(miConfig *types.MIConfig) {
	miDataDir = fmt.Sprintf("%s/.minitializer", userHomeDir)
	projectPath = fmt.Sprintf("%s/%s", miDataDir, miConfig.Metadata["name"])
}

func getMIConfig(filePath string) (*types.MIConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var miConfig types.MIConfig
	err = yaml.Unmarshal(yamlFile, &miConfig)
	if err != nil {
		return nil, err
	}
	err = validateMiFile(&miConfig)
	if err != nil {
		return nil, err
	}
	return &miConfig, err
}

func validateMiFile(miConfig *types.MIConfig) error {
	if miConfig.Metadata["name"] == "" {
		return fmt.Errorf("property mi.metadata.name cannot be null or empty")
	}
	return nil
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
