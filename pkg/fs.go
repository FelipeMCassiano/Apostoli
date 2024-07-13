package pkg

import "os"

// TODO: waldir and remove local repo

func RemoveLocalRepo(path string) {
	os.RemoveAll(path)
}
