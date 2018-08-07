package command

import (
	"fmt"
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Help)
	commands[cmd.Name()] = cmd
}

//Help is the Help command
type Help struct{}

//Name implements Name of Command interface
func (c *Help) Name() string {
	return "HELP"
}

//Help implements Help of Command interface
func (c *Help) Help() string {
	return `Usage: HELP command
Show the usage of the given command`
}

//Execute implements Execute of Command interface
func (c *Help) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Err: ErrWrongArgsNumber}
	}

	cmdName := args[0]
	if cmd, ok := commands[strings.ToUpper(cmdName)]; ok {
		return HelpResult{Cmd: cmd}
	}
	return ErrResult{Err: fmt.Errorf("command %q not found", cmdName)}
}
