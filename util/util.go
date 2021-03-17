package util

import (
	"os"
)

func IsExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
