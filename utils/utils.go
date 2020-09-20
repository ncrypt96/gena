package utils

import (
	"os"
)

// CreateDirIfNotExists creates directory if the directory does not exist
func CreateDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 755)
		if err != nil {
			return err
		}

	}
	return nil
}
