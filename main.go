package main

import (
	"bufio"
	"droidsh/constants"
	"droidsh/handlers"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"log"
	"net"
	"net/textproto"
	"strings"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func main() {

	conn, err := net.Dial("tcp", "localhost:5554")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	defer conn.Close()

	go func() {
		for {
			response, _ := tp.ReadLine()
			if response == "OK" {
				fmt.Printf(constants.PROMPT)
			} else {
				fmt.Println(response)
			}
		}
	}()

	l, err := readline.NewEx(&readline.Config{
		Prompt:          constants.PROMPT,
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	setPasswordCfg := l.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		l.SetPrompt(fmt.Sprintf("Enter password(%v): ", len(line)))
		l.Refresh()
		return nil, 0, false
	})

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line, args, fn, ok := parse(line)
		if ok {
			// Handle builtin
			fn(conn, l, args)
		} else {
			// pass it on to android emulator
			if len(line) > 0 {
				fmt.Fprintln(conn, line)
			}
		}
	}
}

func parse(line string) (string, []string, handlers.CmdHandler, bool) {
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, " ")
	cmd := tokens[0]
	args := tokens[1:]
	fn, ok := handlers.LookUp(cmd)
	return line, args, fn, ok
}
