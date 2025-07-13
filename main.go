package main

import "github.com/broothie/cli"

func main() {
	cli.Run("toss", "Send HTTP requests",
		cli.AddHelpFlag(cli.AddFlagShort('h')),

		cli.AddSubCmd("run", "Run Toss file",
			cli.AddArg(argFile, "File containing toss requests to be run"),
			cli.SetHandler(tossHandler),
		),

		cli.AddSubCmd("list", "List Toss files", cli.SetHandler(listHandler)),
	)
}
