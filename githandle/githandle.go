package githandle

import "github.com/go-git/go-git/v5"

func PlainClone(dirName, repoURL string) error {
	_, err := git.PlainClone(dirName, false, &git.CloneOptions{
		URL:               repoURL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	return err
}
