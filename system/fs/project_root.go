package fs

import (
	"github.com/ironstar-io/tokaido/constants"

	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ProjectRoot - Find the root directory path for the project
func ProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if IsProjectRoot(wd) == true {
		return wd
	}

	return ScanUp(wd)
}

// IsProjectRoot ...
func IsProjectRoot(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == constants.GitDirectory || f.Name() == constants.DockerComposeTokFile {
			return true
		}
	}

	return false
}

// ScanUp ...
func ScanUp(path string) string {
	if path == "/" {
		log.Fatal("Tokaido needs to be run from inside a Git repository, but could not find one. Exiting...")
	}

	sd := filepath.Dir(path)

	if IsProjectRoot(sd) == true {
		return sd
	}

	return ScanUp(sd)
}
