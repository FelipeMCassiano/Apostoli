package pkg

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
)

func WalkThroughDir(path string, fn func(context.Context, string) error) error {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if err := fn(context.Background(), path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func RemoveLocalRepo(path string) {
	os.RemoveAll(path)
}
