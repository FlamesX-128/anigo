package utils

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {

			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ReadDir(path string) []fs.FileInfo {
	fi, err := ioutil.ReadDir(path)

	if err != nil {
		log.Panicln(err)

	}

	return fi
}
