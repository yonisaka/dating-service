package main

import (
	"log"
	"os"

	"github.com/yonisaka/dating-service/cmd"
	"github.com/yonisaka/dating-service/internal/di"
)

func main() {
	cfg := di.GetConfig()
	command := cmd.NewCommand(
		cmd.WithConfig(cfg),
	)

	app := cmd.NewCLI()
	app.Commands = command.Build()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
