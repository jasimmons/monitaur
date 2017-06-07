package monitaur

import (
	"bytes"
	"github.com/jasimmons/monitaur/log"
	"os"
	"os/exec"
	"syscall"
)

const (
	CodeRunError   = -1
	CodeRunSuccess = 0
)

var (
	DefaultShellRunner = ShellRunner{
		shell:     "/bin/sh",
		shellArgs: []string{"-c"},
	}
	DefaultRubyRunner = ShellRunner{
		shell: "/usr/bin/ruby",
	}
)

type Runner interface {
	Run(command string) (exitCode int, stdout, stderr string)
}

type ShellRunner struct {
	shell     string
	shellArgs []string
}

func (sh ShellRunner) Run(command string) (exitCode int, stdout, stderr string) {
	var outbuf, errbuf bytes.Buffer
	cmdSuffix := append(sh.shellArgs, command)
	cmd := exec.Command(sh.shell, cmdSuffix...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			ws := exitErr.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			log.Errorf("unable to get exit code for command: %v, %v\n", sh.shell, cmdSuffix)
			exitCode = CodeRunError
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}

func ParseRunner(runnerType string) Runner {
	switch runnerType {
	case "sh", "bash", "shell":
		if shEnv := os.Getenv("MONITAUR_UNIX_SHELL"); shEnv != "" {
			return ShellRunner{shell: shEnv}
		}
		return DefaultShellRunner
	case "ruby":
		if shEnv := os.Getenv("MONITAUR_RUBY_SHELL"); shEnv != "" {
			return ShellRunner{shell: shEnv}
		}
		return DefaultRubyRunner
	default:
		return DefaultShellRunner
	}
}
