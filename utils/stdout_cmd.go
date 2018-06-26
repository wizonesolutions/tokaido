package utils

import (
	"bitbucket.org/ironstar/tokaido-cli/system/fs"

	"fmt"
	"log"
	"os/exec"
	"strings"
)

// StdoutCmd - Execute a command on the users' OS
func StdoutCmd(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	cmd.Dir = fs.WorkDir()
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Tokaido encountered a fatal error and had to stop at command '%s %s'\n%s", name, strings.Join(args, " "), stdoutStderr)
		log.Fatal(err)
	}

	DebugOutput(stdoutStderr)

	return strings.TrimSpace(string(stdoutStderr))
}
