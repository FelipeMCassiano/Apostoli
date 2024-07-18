package deploy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeMCassiano/Apostoli/pkg"
	"github.com/google/uuid"
)

type deployRequest struct {
	Url string `json:"url"`
}
type deployResponse struct {
	DeployId string `json:"deployid"`
}

// TODO: upload into a s3, push to redis
const (
	a = iota
	b
	c
	d
	e
)

func Deploy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dR := new(deployRequest)

		err := json.NewDecoder(r.Body).Decode(dR)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(a)

		tDir := os.TempDir()

		dId := uuid.New()
		path := fmt.Sprintf("%s/%s/", tDir, dId)

		_, err = pkg.CloneRepo(dR.Url, path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(b)
		defer pkg.RemoveLocalRepo(path)

		log.Println(c)

		if err := pkg.WalkThroughDir(path, pkg.UploadFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(d)
		reponse := deployResponse{
			DeployId: dId.String(),
		}
		log.Println(e)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(reponse); err != nil {
			http.Error(w, "Error when encoding the response", http.StatusUnprocessableEntity)
			return
		}
	}
}
