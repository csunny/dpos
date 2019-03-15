package main

import (
	"io"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
)

// PluginHandler is capable of parsing command line arguments
// and performing
type PluginHandler interface {
	// Lookup receives a potential filename and returns
	// a full or receive path to an excutable, if one
	// exists at the given filename, or an error
	Lookup(filename string) (string, error)

	// Execute receives an executable's filepath, a slice of argument
	// and a slice of environment variables to relay to the executable
	Execute(executablePath string, cmdArgs, environment []string) error
}

type defaultPluginHandler struct {
}

// Lookup implements PluginHandler
func (h *defaultPluginHandler) Lookup(filename string) (string, error) {
	// if on Windows, append the "exe" extension
	// to the filename that we are looking up.

	if runtime.GOOS == "windows" {
		filename = filename + ".exe"
	}

	return exec.LookPath(filename)
}

func (h *defaultPluginHandler) Execute(executablePath string, cmdArgs, environment []string) error {
	return syscall.Exec(executablePath, cmdArgs, environment)
}

// NewDefaultVoteCommand new command
func NewDefaultVoteCommand() *cobra.Command {
	return nil
}

// NewDefaultVoteCommandWithArgs  new command with args
func NewDefaultVoteCommandWithArgs(pluginHandler PluginHandler, args []string, in io.Reader, out, errout io.Writer) *cobra.Command {
	return nil
}

// NewVoteCommand new vote command
func NewVoteCommand(in io.reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added
	cmds := &cobra.Command{
		Use: "dpos",
		Short: i18n.T()
	}

}
