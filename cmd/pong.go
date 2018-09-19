package cmd

import (
	"flag"
	"pong/cmd/pong"

	"github.com/spf13/cobra"
)

var settingsPath string

func init() {
	flag.StringVar(&settingsPath, "s", "pong.json", "Path to settings file.")
	flag.Parse()
	for _, cmd := range commands(settingsPath) {
		rootCmd.AddCommand(cmd)
	}
}

func commands(settingsPath string) []*cobra.Command {
	var cmds []*cobra.Command
	commands, paths := pong.ParseSettings(settingsPath)
	for commandName, description := range commands {
		var cmd = &cobra.Command{
			Use:   commandName,
			Short: description,
		}
		for _, pathName := range paths {
			var pathCmd = &cobra.Command{
				Use: pathName,
				RunE: func(_ *cobra.Command, _ []string) error {
					return pong.Run(settingsPath, cmd.Use, pathName)
				},
			}
			cmd.AddCommand(pathCmd)
		}

		cmds = append(cmds, cmd)
	}
	return cmds
}
