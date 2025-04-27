package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfigs() error {
	const nameResourcesDirectory string = "resources"
	const configFileName string = "env.yaml"
	configFile, err := readConfigFile(
		fmt.Sprintf("%s/%s", nameResourcesDirectory, configFileName))
	if err != nil {
		return err
	}
	fileContent, err := parseConfigFile(configFile)
	if err != nil {
		return err
	}
	err = setConfigVariables(fileContent)
	if err != nil {
		return err
	}
	return nil
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
