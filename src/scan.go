package src

import (
	"os"
	"path/filepath"
)

func scanMigration(mPath string) []string {
	var files []string
	err := filepath.Walk(mPath, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
