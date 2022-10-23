package github

import (
	"github.com/google/go-github/v47/github"
)

type GHRepositoryContent struct {
	Organisation   string
	RepositoryName string
	GhToken        string
}

func (ghRepoContent GHRepositoryContent) CreateFile(path, branch, content, commitMessage, userName, userMail string) (*github.RepositoryContentResponse, error) {

	client, ctx := GHBase{ghToken: ghRepoContent.GhToken}.getCient()
	fileContent := []byte(content)
	opts := &github.RepositoryContentFileOptions{
		Message: github.String(commitMessage),
		Content: fileContent,
		Branch: github.String(branch),
		Committer: &github.CommitAuthor{
			Name: github.String(userName),
			Email: github.String(userMail),
		},
	}
	repoContentResponse, _, repoContentResponseError := client.Repositories.CreateFile(ctx, ghRepoContent.Organisation, ghRepoContent.RepositoryName, path, opts)
	if repoContentResponseError != nil {
		return nil, repoContentResponseError
	}
	return repoContentResponse, nil
}
