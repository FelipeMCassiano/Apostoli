package deploy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/FelipeMCassiano/Apostoli/pkg"
	"github.com/google/uuid"
)

type deployRequest struct {
	Url string `json:"url"`
}

// TODO: upload into a s3, push to redis

func Deploy(w http.ResponseWriter, r *http.Request) {
	dR := new(deployRequest)

	err := json.NewDecoder(r.Body).Decode(dR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tDir := os.TempDir()

	dId := uuid.New()
	path := fmt.Sprintf("%s/%s/", tDir, dId)

	_, err = pkg.CloneRepo(dR.Url, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer pkg.RemoveLocalRepo(path)
}
