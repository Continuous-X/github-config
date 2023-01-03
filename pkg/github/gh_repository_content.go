package github

import (
	"github.com/google/go-github/v48/github"
)

type GHRepositoryContent struct {
	Organisation       string
	RepositoryName     string
	GhToken            string
	GhEnterpriseDomain string
}

func (ghRepoContent GHRepositoryContent) WriteContent(path, branch, content, commitMessage, userName, userMail string) (*github.RepositoryContentResponse, error) {

	fileContent, fileContentErr := ghRepoContent.GetFileContent(path, branch)
	if fileContentErr != nil {
		return ghRepoContent.CreateFile(path, branch, content, commitMessage, userName, userMail)
	}
	return ghRepoContent.UpdateFile(path, branch, content, *fileContent.SHA, commitMessage, userName, userMail)
}


func (ghRepoContent GHRepositoryContent) CreateFile(path, branch, content, commitMessage, userName, userMail string) (*github.RepositoryContentResponse, error) {

	client, ctx := GHBase{
		ghToken:   ghRepoContent.GhToken,
		gheDomain: ghRepoContent.GhEnterpriseDomain,
	}.getCient()
	fileContent := []byte(content)
	opts := &github.RepositoryContentFileOptions{
		Message: github.String(commitMessage),
		Content: fileContent,
		Branch:  github.String(branch),
		Committer: &github.CommitAuthor{
			Name:  github.String(userName),
			Email: github.String(userMail),
		},
	}
	repoContentResponse, _, repoContentResponseError := client.Repositories.CreateFile(ctx, ghRepoContent.Organisation, ghRepoContent.RepositoryName, path, opts)
	if repoContentResponseError != nil {
		return nil, repoContentResponseError
	}
	return repoContentResponse, nil
}

func (ghRepoContent GHRepositoryContent) UpdateFile(path, branch, content, sha, commitMessage, userName, userMail string) (*github.RepositoryContentResponse, error) {

	client, ctx := GHBase{
		ghToken:   ghRepoContent.GhToken,
		gheDomain: ghRepoContent.GhEnterpriseDomain,
	}.getCient()
	fileContent := []byte(content)
	opts := &github.RepositoryContentFileOptions{
		Message: github.String(commitMessage),
		Content: fileContent,
		Branch:  github.String(branch),
		SHA: &sha,
		Committer: &github.CommitAuthor{
			Name:  github.String(userName),
			Email: github.String(userMail),
		},
	}
	fileContentResponse, _, fileContentResponseError := client.Repositories.UpdateFile(ctx, ghRepoContent.Organisation, ghRepoContent.RepositoryName, path, opts)
	if fileContentResponseError != nil {
		return nil, fileContentResponseError
	}
	return fileContentResponse, nil
}

func (ghRepoContent GHRepositoryContent) GetFileContent(path, branch string) (*github.RepositoryContent, error) {

	client, ctx := GHBase{
		ghToken:   ghRepoContent.GhToken,
		gheDomain: ghRepoContent.GhEnterpriseDomain,
	}.getCient()

	opts := &github.RepositoryContentGetOptions{
		Ref: branch,
	}

	fileContent, _, _, repoContentResponseError := client.Repositories.GetContents(ctx, ghRepoContent.Organisation, ghRepoContent.RepositoryName, path, opts)
	if repoContentResponseError != nil {
		return nil, repoContentResponseError
	}

	return fileContent, nil
}

func (ghRepoContent GHRepositoryContent) GetFileContentDecoded(path, branch string) (string, error) {

	fileContent, fileContentErr := ghRepoContent.GetFileContent(path, branch)
	if fileContentErr != nil {
		return "", fileContentErr
	}

	decodedFileContent, decodedErr := fileContent.GetContent()
	if decodedErr != nil {
		return "", decodedErr
	}
	return decodedFileContent, nil
}
