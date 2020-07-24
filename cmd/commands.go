package cmd

import (
	"github.com/spf13/cobra"
	"pqredis/config"
	"pqredis/monitoring/logger"
	"pqredis/postgres"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use: "pqredis",
	}
	cli.AddCommand(newStartCmd())

	return cli
}

func newStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "start",
		Short:   "Start pqredis",
		Aliases: []string{"start"},
		Run: func(_ *cobra.Command, _ []string) {
			l := logger.New(config.New())
			l.Infof("hello pqredis")

			cfg := config.New()
			_, err := postgres.New(cfg, l)
			if err != nil {
				l.Errorf(err.Error())
			} else {
				l.Infof("DB is set up")
			}
		},
	}
}
