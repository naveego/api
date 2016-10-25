package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Shell struct {
	gobin string
}

func NewShell(gobin string) *Shell {
	return &Shell{
		gobin: gobin,
	}
}

func (s *Shell) GoRun(directory, packageName string, arguments []string) (output string, err error) {
	argsStr := strings.Join(arguments, " ")
	cmd := NewCommand(directory, s.gobin, "run", "-i", argsStr).Execute()
	return cmd.Output, cmd.Error
}

func (s *Shell) GoBuild(directory, packageName, outputName string, arguments []string) (outputPath string, err error) {
	argsStr := strings.Join(arguments, " ")
	cmd := NewCommand(directory, s.gobin, "build", "-i", argsStr).Execute()
	outputPath = filepath.Join(directory, outputName)
	return outputPath, cmd.Error
}

type Command struct {
	directory  string
	executable string
	arguments  []string

	Output string
	Error  error
}

func NewCommand(directory, executable string, arguments ...string) Command {
	return Command{
		directory:  directory,
		executable: executable,
		arguments:  arguments,
	}
}

func (c Command) Execute() Command {
	if len(c.executable) == 0 {
		return c
	}

	if len(c.Output) > 0 || c.Error != nil {
		return c
	}

	command := exec.Command(c.executable, c.arguments...)
	command.Dir = c.directory
	command.Env = []string{"GOOS=linux", "GOARCH=amd64", fmt.Sprintf("GOPATH=%s", os.Getenv("GOPATH"))}
	var rawOutput []byte
	rawOutput, c.Error = command.CombinedOutput()
	c.Output = string(rawOutput)
	return c
}
