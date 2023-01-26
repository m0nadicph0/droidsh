package handlers

import (
	"droidsh/cmd"
	"github.com/chzyer/readline"
	"net"
)

type CmdHandler func(conn net.Conn, rl *readline.Instance, args []string)

var handlerMap = map[string]CmdHandler{
	"prompt": cmd.SetPrompt,
	"clear":  cmd.Clear,
	"loc":    cmd.Location,
}

func LookUp(cmd string) (CmdHandler, bool) {
	h, ok := handlerMap[cmd]
	return h, ok
}
