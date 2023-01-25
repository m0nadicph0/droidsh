package cmd

import (
	"droidsh/constants"
	"github.com/chzyer/readline"
	"net"
	"os"
	"strings"
)

func SetPrompt(conn net.Conn, rl *readline.Instance, args []string) {
	rl.Refresh()
	switch len(args) {
	case 0:
		rl.SetPrompt(constants.PROMPT)
	default:
		rl.SetPrompt(strings.Join(args, " ") + " ")
	}
}

func Clear(conn net.Conn, rl *readline.Instance, args []string) {
	readline.ClearScreen(os.Stdout)
}
