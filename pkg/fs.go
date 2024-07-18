package pkg

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func WalkThroughDir(path string, fn func(string) error) error {
	// Better than using normal go funcs 'cause treats errors
	errg, _ := errgroup.WithContext(context.Background())

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			errg.Go(func() error {
				if strings.Contains(path, ".git") {
					return nil
				}
				err = fn(path)
				return err
			})
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := errg.Wait(); err != nil {
		return err
	}
	return nil
}

func RemoveLocalRepo(path string) {
	os.RemoveAll(path)
}
