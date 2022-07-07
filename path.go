package utils

import "path"

// ParsePath 解析路径
func ParsePath(filePath string) (dir, name, ext string) {
	ext = path.Ext(filePath)[1:]
	name = path.Base(filePath)[:len(path.Base(filePath))-len(ext)-1]
	dir = path.Dir(filePath)
	return
}
