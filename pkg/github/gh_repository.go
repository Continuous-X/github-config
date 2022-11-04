package github

import (
	"fmt"
	"log"

	"github.com/google/go-github/v47/github"
)

type GHRepository struct {
	Organisation   string
	RepositoryName string
	GhToken        string
}

func (ghRepo GHRepository) IsCompliant() bool {
	ghRepo.getRepositoryList()
	return false
}

func (ghRepo GHRepository) getRepositoryList() ([]*github.Repository, error) {

	client, ctx := GHBase{ghToken: ghRepo.GhToken}.getCient()
	repositoryList, _, listError := client.Repositories.List(ctx, "", nil)
	if listError != nil {
		return repositoryList, listError
	}
	for index, repository := range repositoryList {
		log.Print(fmt.Sprintf("%d: %s", index, repository.GetFullName()))
	}
	return repositoryList, nil
}
