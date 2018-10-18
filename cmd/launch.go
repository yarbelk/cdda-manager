package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yarbelk/cdda-manager/lib/game"
	"github.com/yarbelk/cdda-manager/lib/launcher"
)

var (
	gamePath   string
	workingDir string
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch a game",
	Long: `Launch a downloaded version of the game:

Eventually; you should be able just say launch world X,
and it will launch it with your configuration and mods`,
	RunE: func(cmd *cobra.Command, args []string) error {
		gameVersion := game.Game{
			Path:     gamePath,
			GameData: workingDir,
		}
		return launcher.LaunchGame(gameVersion)
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

	launchCmd.Flags().StringVar(&gamePath, "gameExecutable", "", "Executable for game")
	launchCmd.Flags().StringVar(&workingDir, "workingDir", "", "working dir")
}
