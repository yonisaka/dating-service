package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/server"
	"github.com/yonisaka/dating-service/pkg/logger"
	"github.com/yonisaka/dating-service/pkg/migrator"
)

// httpStart is a method to start http server
func (cmd *Command) httpStart() *cli.Command {
	return &cli.Command{
		Name:  "http:start",
		Usage: "A command to start http server",
		Action: func(c *cli.Context) error {
			httpServer := server.NewHTTPServer()

			logger.Info(
				logger.MessageFormat("starting dating-service services... %d", cmd.App.Port),
				logger.EventName(consts.LogEventNameServiceStarting),
			)
			if err := httpServer.Run(); err != nil {
				return err
			}

			return nil
		},
	}
}

// migrateDatabase is a method to migrate database
func (cmd *Command) migrateDatabase() *cli.Command {
	cfg := config.New()

	return &cli.Command{
		Name:  "db:migrate",
		Usage: "A command to migrate database",
		Action: func(c *cli.Context) error {
			migrator.DatabaseMigration(&cfg.MasterDB)

			return nil
		},
	}
}
