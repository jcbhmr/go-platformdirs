package main_test

import (
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
import assert from "node:assert";
import test from "node:test";
import { $ } from "execa";
import which from "which";

	test("matches python platformdirs", async (t) => {
		// https://docs.astral.sh/uv/getting-started/installation/
		if (!(await which("uv", { nothrow: true }))) {
			t.skip("uv not found");
			return;
		}

		// Strip first line which has version info like "-- platformdirs 1.2.3 --".
		const expected = (
			await $`uv run --with platformdirs python -m platformdirs`
		).stdout.replace(/^.*[]/m, "");
		const actual = (await $`node src/main.js`).stdout.replace(/^.*[]/m, "");
		assert.equal(actual, expected);
	});
*/
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
