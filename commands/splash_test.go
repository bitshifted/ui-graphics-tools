package commands

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplashScreenGenerateSuccess(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	outputDir, err := os.MkdirTemp("", "splash-test")
	assert.NoError(t, err)
	defer os.Remove(outputDir)
	params := SplashParams{
		ConfigFile: "test-resources/test-config.yaml",
		OutputDir:  outputDir,
	}
	result := GenerateSplashScreen(&params)
	assert.Nil(t, result)
}
