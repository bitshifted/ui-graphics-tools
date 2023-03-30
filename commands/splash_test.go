package commands

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...) //nolint:gosec
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func Test_SplashScreenGenerateSuccess(t *testing.T) {
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
