package unit

import (
	"os"
	"testing"
	"workflowmanager/app/util"
)

func TestLoadExistedConfigs(test *testing.T) {
	const existedConfigFilePath string = "../../app/resources/dev_env.yaml"
	err := util.LoadConfigs(existedConfigFilePath)
	if err != nil {
		test.FailNow()
	}
	if os.Getenv("POSTGRES_USER") != "postgres" ||
		os.Getenv("SERVICE_PORT") != "8080" {
		test.Errorf("the configs have diference values")
	}
}
