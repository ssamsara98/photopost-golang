package bootstrap

import (
	"github.com/spf13/cobra"
	"github.com/ssamsara98/photopost-golang/src/commands"
)

var rootCmd = &cobra.Command{
	Use:   "clean-architecture",
	Short: "Commander for clean architecture",
	Long: `
		█▀▀ █░░ █▀▀ ▄▀█ █▄░█ ▄▄ █▀▀ █ █▄░█
		█▄▄ █▄▄ ██▄ █▀█ █░▀█ ░░ █▄█ █ █░▀█

	This is a command runner or cli for api architecture in golang.
	Using this we can use underlying dependency injection container for running scripts.
	Main advantage is that, we can use same services, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

// App root of the application
type App struct {
	*cobra.Command
}

// NewApp creates new root command
func NewApp() *App {
	cmd := &App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
