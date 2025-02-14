package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/migrate"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var mirgrateCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "migrate <FOLDER>",
	Short:                 "Migrate all the repos in a directory and its subdirectories",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/#-fgh-migrate",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			folder   = migrate.EnsureFolderExists(args)
			oldRepos = repos.Repos(folder)
			config   = configuration.GetConfig(false)
			newPaths = migrate.NewPaths(oldRepos, config)
		)
		migrate.ConfirmMove(newPaths)
		utils.MoveRepos(newPaths)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(mirgrateCmd)
}
