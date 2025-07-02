package pike

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

type gitCloneError struct {
	repository  string
	destination string
	err         error
}

func (m *gitCloneError) Error() string {
	return fmt.Sprintf("failed to clone repository %s %s %v", m.repository, m.destination, m.err)
}

type gitHeadError struct {
	repository  string
	destination string
	err         error
}

func (m *gitHeadError) Error() string {
	return fmt.Sprintf("failed to get head %s %s %v", m.repository, m.destination, m.err)
}

type gitCommitObjectError struct {
	repository  string
	destination string
	err         error
}

func (m *gitCommitObjectError) Error() string {
	return fmt.Sprintf("failed to get commit object %s %s %v", m.repository, m.destination, m.err)
}

func Repository(repository, destination, directory, output string, init, write, enableResources bool) error {
	if _, err := os.Stat(destination); !os.IsNotExist(err) {
		log.Info().Msgf("%s was not empty, removing", destination)
		_ = os.RemoveAll(destination)
	}

	// Clone the given repository to the given directory
	log.Info().Msgf("git clone %s %s --recursive", repository, destination)

	r, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:               repository,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Depth:             1,
	})
	if err != nil {
		return &gitCloneError{repository, destination, err}
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return &gitHeadError{repository, destination, err}
	}
	// ... retrieving the commit object
	_, err = r.CommitObject(ref.Hash())
	if err != nil {
		return &gitCommitObjectError{repository, destination, err}
	}

	return Scan(filepath.Join(destination, directory), output, nil, init, write, enableResources, "", "", "")
}
