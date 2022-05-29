package utils

import "os"

func DeleteFile(localfile string) error {
	return os.Remove(localfile)
}
