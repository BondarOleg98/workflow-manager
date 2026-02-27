package unit

import (
	"os"
	"testing"
	"workflowmanager/app/util"
)

func TestLoadExistedConfigs(test *testing.T) {
	const existedConfigFilePath string = "../../app/resources/dev_env.yaml"
	util.LoadConfigs(existedConfigFilePath)
	if os.Getenv("POSTGRES_USER") != "postgres" ||
		os.Getenv("SERVICE_PORT") != "8080" {
		test.Errorf("the configs have diference values")
	}
}
