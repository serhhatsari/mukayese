package internal

import (
	"log"
	"os"
	"path/filepath"
)

func deleteFirstPart(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			return s[i+1:]
		}
	}
	return s
}

func ListFilesChecksums(summary map[string]string, path string) {

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	//infos := make([]fs.FileInfo, 0, len(entries))

	for _, entry := range entries {

		// If it exists in exclude list, bypass
		if Exclude(entry.Name()) {
			continue
		}

		newPath := filepath.Join(path, entry.Name())

		// TODO: Error for dot path
		// project/dll/hello.cs -> dll/hello.cs
		//x := strings.Split(newPath, string(os.PathSeparator))
		//var y string
		//for _, val := range x[1:] {
		//	y = filepath.Join(y, val)
		//}

		// check is dir
		if !entry.IsDir() {

			// remove the first part of path
			updatedPath := deleteFirstPart(newPath)

			summary[updatedPath] = Hasher(&newPath)
		} else {
			ListFilesChecksums(summary, newPath)
		}

	}
}
