package clone

import (
	"fmt"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get the meta data about the repo
func GetRepository(secrets configure.SecretsOutline, args []string) api.Repo {
	owner, name := OwnerAndName(configuration.GetSecrets().Username, args)
	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Getting metadata for %v/%v", owner, name)
	spin.Start()

	client := api.GenerateClient(configuration.GetSecrets().PAT)
	repo, err := api.RepoData(client, owner, name)
	if err != nil {
		fmt.Println()
		statuser.Error("Failed to get repo information", err, 1)
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("Got metadata for %v/%v\n", owner, name))
	return repo
}

// Get the name of the repo and the of the owner
func OwnerAndName(username string, args []string) (owner string, name string) {
	if strings.Contains(args[0], "/") {
		parts := strings.Split(args[0], "/")
		owner = parts[0]
		name = parts[1]
	} else {
		owner = username
		name = args[0]
	}

	if owner == "" {
		statuser.ErrorMsg("No owner provided", 1)
	}
	if name == "" {
		statuser.ErrorMsg("No repository name provided", 1)
	}

	return owner, name
}
