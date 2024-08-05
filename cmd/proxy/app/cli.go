package app

import "github.com/desertbit/grumble"

// App is used to register the grumble
var App = grumble.New(&grumble.Config{
	Name:                  "ljpos-li",
	Description:           "Ljpos-li - An advanced, yet simple tunneling tool",
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})
