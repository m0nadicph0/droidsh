package main

import (
	"droidsh/utils"
	"github.com/chzyer/readline"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("prompt", readline.PcItemDynamic(utils.PromptGenerator)),
	readline.PcItem("clear"),
	readline.PcItem("loc", readline.PcItemDynamic(utils.CityGenerator)),
)
