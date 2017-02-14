package main

import (
	"github.com/WindomZ/cli-mate"
	"github.com/WindomZ/gitclone/gitclone"
)

func main() {
	cmd := cli_mate.NewApp(
		"gitclone",
		"A cli tool, git clone repository in the `go get` style.",
		"0.3",
	)

	cmd.Action = gitclone.RootAction

	cmd.AddFlags([]cli_mate.Flag{
		gitclone.UrlFlag, // url
	})

	cmd.AddCommands([]cli_mate.Command{
		gitclone.CloneCommand, // clone
		gitclone.ListCommand,  // list
		gitclone.MkdirCommand, // mkdir
	})

	cmd.RunOSArgs()
}
