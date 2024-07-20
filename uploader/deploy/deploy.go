package deploy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/FelipeMCassiano/Apostoli/cg"
	"github.com/FelipeMCassiano/Apostoli/pkg"
	"github.com/google/uuid"
)

type deployRequest struct {
	Url string `json:"url"`
}
type deployResponse struct {
	DeployId string `json:"deployid"`
}

func Deploy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if err := pkg.WalkThroughDir(path, pkg.UploadFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reponse := deployResponse{
			DeployId: dId.String(),
		}

		rctx := r.Context()

		cg.RedisClient.LPush(rctx, "builds", dId.String(), "uploading")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(reponse); err != nil {
			http.Error(w, "Error when encoding the response", http.StatusUnprocessableEntity)
			return
		}
	}
}
