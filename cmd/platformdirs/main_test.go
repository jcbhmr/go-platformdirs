package main_test

import (
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesPythonPlatformdirs(t *testing.T) {
	// if _, err := exec.LookPath("uv"); err != nil {
	// 	t.Skip("uv not found")
	// }

	// Strip first line which has version info like "-- platformdirs 1.2.3 --".
	cmd := exec.Command("uv", "run", "--with", "platformdirs", "python", "-m", "platformdirs")
	stdout, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	expected := regexp.MustCompile(`^.*\n`).ReplaceAllString(string(stdout), "")

	cmd = exec.Command("go", "run", ".")
	stdout, err = cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	actual := regexp.MustCompile(`^.*\n`).ReplaceAllString(string(stdout), "")

	assert.Equal(t, expected, actual)
}
