package python

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	rt "runtime"
	"strings"
)

var pythonVersionCommand = []string{"python3", "--version"}


func pythonVersionInfo() (version, runtime, vm string, err error) {
	var out bytes.Buffer

	cmd := exec.Command(pythonVersionCommand[0], pythonVersionCommand[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to check python version: %v", err)
		return
	}

	version, runtime, vm = parsePythonVersionOutput(out.String())
	return
}

func parsePythonVersionOutput(infoString string) (version, runtime, vm string) {
	infoString = strings.TrimSpace(infoString)

	lines := strings.Split(infoString, "\n")
	

	if len(lines) != 1 {
		// unexpected output format, don't attempt to parse output for version
		return "", "", ""
	}

	versionString := strings.TrimSpace(lines[0])

	re := regexp.MustCompile(`Python (\d\.\d.\d)`)
	if match := re.FindStringSubmatch(lines[0]); len(match) == 2 {
		versionString = match[1]
	}

	return versionString, "", ""
}
