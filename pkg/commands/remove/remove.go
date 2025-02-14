package remove

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Ask to remove each repo and then remove it
func RemoveRepos(clonedRepos []repos.LocalRepo) {
	for _, repo := range clonedRepos {
		committed, pushed := repos.WorkingState(repo.Path)
		if !committed {
			statuser.Warning(
				fmt.Sprintf("Repository located at %v has uncommitted changes", repo.Path),
			)
		}
		if !pushed {
			statuser.Warning(
				fmt.Sprintf("Repository located at %v has changes not pushed to a remote", repo.Path),
			)
		}
		remove := utils.Confirm(fmt.Sprintf(
			"Are you sure you want to permanently remove %v from your computer?", repo.Path,
		))
		if remove {
			err := os.RemoveAll(repo.Path)
			if err != nil {
				statuser.Error("Failed to remove "+repo.Path, err, 1)
			}
			statuser.Success("Removed " + repo.Path)
		}
	}
}
