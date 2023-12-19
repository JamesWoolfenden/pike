package pike

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

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
		return err
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return err
	}
	// ... retrieving the commit object
	_, err = r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	return Scan(filepath.Join(destination, directory), output, nil, init, write, enableResources)
}
