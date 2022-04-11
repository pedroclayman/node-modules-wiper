package modulesearch

import (
	"os"
	"path"
	"path/filepath"
)

func GetNodeModuleDirectories(root string, foundDirs *[]string) error {
	possibleDirs, err := os.ReadDir(root)

	if err != nil {
		return err
	}

	for _, dir := range possibleDirs {
		if dir.IsDir() {
			if dir.Name() == "node_modules" {
				*foundDirs = append(*foundDirs, root+"/"+dir.Name())
			} else {
				GetNodeModuleDirectories(path.Join(root, dir.Name()), foundDirs)
			}
		}
	}

	return nil
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
