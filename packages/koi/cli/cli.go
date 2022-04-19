package cli

import (
	"github.com/urfave/cli/v2"
	"koi/config"

	log "github.com/sirupsen/logrus"
)

const (
	KoiErrSpin = "KOI_ERR_SPIN"
)

var (
	// Log
	l = log.WithField("package", "cli")
)

func Run(args []string) error {
	l.Debug("Constructing cli app")

	cli.VersionPrinter = func(c *cli.Context) {
		l.Info(config.Version)
	}

	app := &cli.App{
		Name:    "Koi",
		Usage:   "The Koishi Launcher.",
		Version: config.Version,
		Authors: []*cli.Author{
			{
				Name:  "Il Harper",
				Email: "hi@ilharper.com",
			},
		},

		UseShortOptionHandling: true,
		EnableBashCompletion:   true,

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Use configuration from `FILE`",
				EnvVars: []string{"KOI_CONFIG"},
			},

			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug mode",
			},

			cli.HelpFlag,
			cli.VersionFlag,
			cli.BashCompletionFlag,
		},

		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run Koishi",
				Action: RunAction,
			},
		},

		Before: PreAction,
		CommandNotFound: func(context *cli.Context, s string) {
			l.Fatalf("Command not found: %s", s)
		},

		//Action: RunAction,
	}

	l.Debug("Running cli app")
	err := app.Run(args)
	if err != nil && err.Error() != KoiErrSpin {
		l.Fatal(err)
	}
	return err
}
