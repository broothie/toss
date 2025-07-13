package main

import "github.com/broothie/cli"

func main() {
	cli.Run("toss", "Send HTTP requests",
		cli.AddHelpFlag(
			cli.AddFlagShort('h'),
			cli.SetFlagIsInherited(true),
		),

		cli.AddSubCmd("run", "Run Toss file",
			cli.AddArg(argFile, "File containing toss requests to be run"),
			cli.SetHandler(tossHandler),
		),

		cli.AddSubCmd("list", "List Toss files", cli.SetHandler(listHandler)),

		cli.AddSubCmd("init", "Create a new toss file",
			cli.AddArg(argInitName, "Optional name prefix for the toss file", cli.SetArgDefault("")),

			cli.AddFlag(flagInitDirectory, "Directory to create the file in",
				cli.AddFlagShort('d'),
				cli.SetFlagDefault("."),
			),

			cli.AddFlag(flagInitFileType, "File type (yml, json, toml)",
				cli.AddFlagShort('f'),
				cli.SetFlagDefault("yml"),
			),

			cli.SetHandler(initHandler),
		),
	)
}
