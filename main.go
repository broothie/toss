package main

import "github.com/broothie/cli"

func main() {
	cli.Run("toss", "Send HTTP requests",
		cli.AddHelpFlag(cli.AddFlagShort('h')),
		cli.AddArg(argFile, "File containing toss requests to be run"),
		cli.SetHandler(tossHandler),
	)
}
