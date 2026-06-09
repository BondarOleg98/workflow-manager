package unit

import (
	"os"
	"testing"
	"workflowmanager/app/util"
)

func TestLoadCorrectConfigs(test *testing.T) {
	const correctConfigFilePath string = "../resources/test_env.yaml"
	util.LoadConfigs(correctConfigFilePath)
	if os.Getenv("CORRECT_VALUE_STRING") != "postgres" ||
		os.Getenv("CORRECT_VALUE_NUM") != "8080" {
		test.Errorf("the configs have diference values")
	}
}
