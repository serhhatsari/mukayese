package internal

import "os"

// isDirectory determines if a file represented
// by `path` is a directory or not
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func IsEmpty(datas ...map[string]string) bool {
	isEmpty := true
	for _, data := range datas {
		if 0 < len(data) {
			isEmpty = false
		}
	}
	return isEmpty
}
