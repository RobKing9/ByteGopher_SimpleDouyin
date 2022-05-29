package utils

import "os"

// 根据文件路径删除一个文件
func DeleteFile(localfile string) error {
	return os.Remove(localfile)
}
