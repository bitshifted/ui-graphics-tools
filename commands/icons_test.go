package commands

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIconsGenerateSuccess(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	outputDir, err := os.MkdirTemp("", "splash-test")
	assert.NoError(t, err)
	defer os.Remove(outputDir)
	params := IconsParams{
		InputFile:       "test-resources/test-icon.svg",
		OutputDirectory: outputDir,
		Size16x16:       true,
		Size32x32:       true,
		Size64x64:       true,
		Size128x128:     true,
		Size256x256:     true,
		Size512x512:     true,
		Verbose:         true,
		GenerateIco:     true,
		GenerateIcns:    true,
	}
	result := GenerateIcons(&params)
	assert.NoError(t, result)
}
