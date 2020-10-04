package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncrypt96/gena/constants"
)

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DirIsEmpty(directory string) (bool, error) {
	c := 0
	err := filepath.Walk(directory, func(path string, i os.FileInfo, err error) error {
		if c > 1 {
			return fmt.Errorf(constants.DirNotEmptyErr, directory)
		}
		c++
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, err
}

func CreateDirIfNotExists(dirName string) error {
	err := os.MkdirAll(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func GetPaths(dirName string, addGit bool) (paths []string, err error) {
	err = filepath.Walk(dirName, func(path string, i os.FileInfo, err error) error {
		if !addGit && strings.Contains(path, ".git") && i.IsDir() {
			return filepath.SkipDir
		}
		if !i.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	return
}
