package util

import (
	"gopkg.in/yaml.v3"
	"log"
	"log/slog"
	"os"
	"time"
)

func LoadConfigs(filePath string) {
	configFile, err := readConfigFile(filePath)
	if err != nil {
		log.Fatalf("Error during reading file: %s", err)
	}
	fileContent, err := parseConfigFile(configFile)
	if err != nil {
		log.Fatalf("Error during parsing values from the config file: %s", err)
	}
	err = setConfigVariables(fileContent)
	if err != nil {
		log.Fatalf("Error during setting values from the config file: %s", err)
	}
}

func readConfigFile(configFilePath string) ([]byte, error) {
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	return configFile, nil
}

func parseConfigFile(configFile []byte) (map[string]string, error) {
	fileContent := make(map[string]string)
	err := yaml.Unmarshal(configFile, &fileContent)
	if err != nil {
		return nil, err
	}
	return fileContent, err
}

func setConfigVariables(fileContent map[string]string) error {
	var err error
	for key, value := range fileContent {
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseTimeConfigVariable(configTimeVariable string) time.Duration {
	duration, err := time.ParseDuration(configTimeVariable)
	if err != nil {
		slog.Error("Error during parsing the time env variable, getting the default value 1m", "err", err)
		return 1 * time.Minute
	}
	return duration
}
