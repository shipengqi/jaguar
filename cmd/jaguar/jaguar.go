package main

import "github.com/shipengqi/jcli"

func main() {
	app := jcli.New(
		"jaguar",
		jcli.WithDesc(""),
		jcli.DisableConfig(),
	)

	app.AddCommands(newCmd())

	app.Run()
}
